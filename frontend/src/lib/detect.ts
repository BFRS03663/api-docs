export type DetectedFormat = "openapi3" | "openapi2" | "postman" | "unknown";

export const formatLabel: Record<DetectedFormat, string> = {
  openapi3: "OpenAPI 3",
  openapi2: "Swagger 2.0",
  postman: "Postman collection",
  unknown: "Unrecognised",
};

/**
 * Cheap client-side sniff of an uploaded file so the import form can show
 * what it thinks the file is before the server parses it. The backend does
 * the authoritative detection.
 */
export function detectFormat(text: string): DetectedFormat {
  const head = text.slice(0, 64 * 1024);
  if (/["']?openapi["']?\s*:\s*["']?3\./.test(head)) return "openapi3";
  if (/["']?swagger["']?\s*:\s*["']?2\.0/.test(head)) return "openapi2";
  if (/_postman_id|getpostman\.com\/json\/collection/.test(head)) return "postman";
  if (/"item"\s*:\s*\[/.test(head) && /"info"\s*:/.test(head)) return "postman";
  return "unknown";
}

/** Derives a URL-safe slug from a name or filename. */
export function slugify(input: string): string {
  return input
    .replace(/\.(json|ya?ml)$/i, "")
    .replace(/\.postman_collection$/i, "")
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, "-")
    .replace(/^-+|-+$/g, "")
    .slice(0, 64);
}

export const slugPattern = /^[a-z0-9]+(?:-[a-z0-9]+)*$/;
