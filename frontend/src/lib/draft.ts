import { parseDocument, stringify } from "yaml";
import type { DetectedFormat } from "./detect";

export interface DraftOk {
  ok: true;
  format: "openapi3" | "openapi2";
  title: string;
  version: string;
  pathCount: number;
}

export interface DraftError {
  ok: false;
  message: string;
  line?: number;
  col?: number;
}

export type Draft = DraftOk | DraftError;

/**
 * Parses the editor text far enough to know whether Scalar can render it and
 * to fill the form (title → slug). The backend does the authoritative
 * validation on publish.
 */
export function parseDraft(text: string): Draft {
  if (!text.trim()) return { ok: false, message: "The document is empty" };
  const doc = parseDocument(text, { prettyErrors: true });
  const first = doc.errors[0];
  if (first) {
    const pos = first.linePos?.[0];
    return { ok: false, message: first.message.split("\n")[0], line: pos?.line, col: pos?.col };
  }
  let root: unknown;
  try {
    root = doc.toJS();
  } catch (err) {
    return { ok: false, message: err instanceof Error ? err.message : "Could not read the document" };
  }
  if (!root || typeof root !== "object" || Array.isArray(root)) {
    return { ok: false, message: "The top level must be a mapping with keys such as openapi, info and paths" };
  }
  const spec = root as Record<string, unknown>;
  const info = (typeof spec.info === "object" && spec.info ? spec.info : {}) as Record<string, unknown>;
  switch (sniff(spec, info)) {
    case "openapi3":
    case "openapi2":
      break;
    case "postman":
      return { ok: false, message: "Postman collections cannot be edited here. Import the file instead; it is converted to OpenAPI." };
    default:
      return { ok: false, message: 'Add "openapi: 3.0.3" (or swagger: "2.0") at the top level so the document is recognised as an API description' };
  }
  const pathCount = typeof spec.paths === "object" && spec.paths ? Object.keys(spec.paths).length : 0;
  // Mirrors the importer, which refuses documents without endpoints.
  if (pathCount === 0) return { ok: false, message: "Add at least one endpoint under paths; a document without endpoints cannot be published" };
  return {
    ok: true,
    format: typeof spec.openapi === "string" ? "openapi3" : "openapi2",
    title: asString(info.title),
    version: asString(info.version),
    pathCount,
  };
}

function sniff(spec: Record<string, unknown>, info: Record<string, unknown>): DetectedFormat {
  if (typeof spec.openapi === "string" && spec.openapi.startsWith("3.")) return "openapi3";
  if (spec.swagger === "2.0" || spec.swagger === 2) return "openapi2";
  if (info._postman_id || Array.isArray(spec.item)) return "postman";
  return "unknown";
}

function asString(v: unknown): string {
  return typeof v === "string" ? v : v == null ? "" : String(v);
}

/** Converts a JSON document (e.g. the canonical spec of a Postman import) to YAML for editing. */
export function specToYaml(json: string): string {
  return stringify(JSON.parse(json) as unknown, { lineWidth: 0, aliasDuplicateObjects: false });
}

/**
 * What a new document starts as: the top-level sections filled in with
 * placeholders and comments, and no endpoints yet.
 */
export const skeletonTemplate = `openapi: 3.0.3
info:
  title: My API
  version: 1.0.0
  description: |
    What this API does, who it is for and how to get access.
    Markdown is supported.
  contact:
    name: API Support
    email: support@example.com
servers:
  - url: https://api.example.com/v1
    description: Production
  - url: https://sandbox.example.com/v1
    description: Sandbox
tags:
  - name: Example
    description: Endpoints are grouped in the sidebar by tag.
security:
  - bearerAuth: []
paths:
  # One entry per URL path. Replace this placeholder with your endpoints.
  /health:
    get:
      tags: [Example]
      summary: Health check
      operationId: getHealth
      responses:
        "200":
          description: The service is up
components:
  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
  schemas: {}
`;

/** A small but complete OpenAPI 3 document with two documented endpoints. */
export const exampleTemplate = `openapi: 3.0.3
info:
  title: My API
  version: 1.0.0
  description: |
    Describe what this API does. Markdown is supported.
servers:
  - url: https://api.example.com/v1
tags:
  - name: Orders
    description: Create and track orders
paths:
  /orders:
    get:
      tags: [Orders]
      summary: List orders
      operationId: listOrders
      parameters:
        - name: status
          in: query
          schema:
            type: string
            enum: [new, shipped, delivered]
      responses:
        "200":
          description: A page of orders
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/Order"
    post:
      tags: [Orders]
      summary: Create an order
      operationId: createOrder
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/Order"
      responses:
        "201":
          description: Created
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Order"
components:
  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
  schemas:
    Order:
      type: object
      required: [id, status]
      properties:
        id:
          type: string
          example: ord_123
        status:
          type: string
          example: new
security:
  - bearerAuth: []
`;
