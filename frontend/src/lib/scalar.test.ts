import { operationAnchor, scalarConfig, specUrl } from "./scalar";

test("specUrl points at the backend spec route and escapes the slug", () => {
  expect(specUrl("petstore")).toBe("/docs/petstore/openapi.json");
  expect(specUrl("a b")).toBe("/docs/a%20b/openapi.json");
});

test("scalarConfig wires url and the backend proxy by default", () => {
  const cfg = scalarConfig("petstore");
  expect(cfg.url).toBe("/docs/petstore/openapi.json");
  expect(cfg.proxyUrl).toBe("/api/v1/proxy");
  expect(scalarConfig("petstore", "").proxyUrl).toBeUndefined();
});

test("operationAnchor follows Scalar's tag/METHOD/path format", () => {
  expect(operationAnchor({ method: "get", path: "/pet/{petId}", tags: ["pet"] })).toBe("#tag/pet/GET/pet/{petId}");
  expect(operationAnchor({ method: "POST", path: "/x", tags: [] })).toBe("");
});
