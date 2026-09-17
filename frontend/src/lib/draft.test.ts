import { exampleTemplate, parseDraft, skeletonTemplate, specToYaml } from "./draft";

test("accepts an OpenAPI 3 document and reads its title", () => {
  const d = parseDraft("openapi: 3.0.3\ninfo:\n  title: Mini\n  version: '2'\npaths:\n  /a: {}\n  /b: {}\n");
  expect(d).toMatchObject({ ok: true });
  expect(d).toEqual({ ok: true, format: "openapi3", title: "Mini", version: "2", pathCount: 2 });
});

test("accepts Swagger 2.0", () => {
  expect(parseDraft('swagger: "2.0"\ninfo: {title: Old}\npaths: {/x: {}}\n')).toMatchObject({ ok: true, format: "openapi2", title: "Old" });
});

test("reports YAML syntax errors with their line", () => {
  const d = parseDraft("openapi: 3.0.3\ninfo: title: x\n");
  expect(d.ok).toBe(false);
  if (d.ok) return;
  expect(d.line).toBe(2);
  expect(d.message).toMatch(/line 2/);
});

test("rejects empty, non-object and unrecognised documents", () => {
  expect(parseDraft("   ")).toMatchObject({ ok: false, message: /empty/ });
  expect(parseDraft("- a\n- b\n")).toMatchObject({ ok: false, message: /top level/ });
  expect(parseDraft("hello: world\n")).toMatchObject({ ok: false, message: /openapi: 3\.0\.3/ });
  expect(parseDraft("openapi: 3.0.3\ninfo: {title: Empty, version: '1'}\npaths: {}\n")).toMatchObject({ ok: false, message: /at least one endpoint/ });
  expect(parseDraft('{"info":{"_postman_id":"x"},"item":[]}')).toMatchObject({ ok: false, message: /Postman/ });
});

test("skeleton and example templates are valid; JSON converts to YAML", () => {
  expect(parseDraft(skeletonTemplate)).toMatchObject({ ok: true, format: "openapi3", title: "My API", version: "1.0.0", pathCount: 1 });
  expect(parseDraft(exampleTemplate)).toMatchObject({ ok: true, title: "My API", pathCount: 1 });
  const yaml = specToYaml('{"openapi":"3.0.0","info":{"title":"From JSON","version":"1"},"paths":{"/x":{}}}');
  expect(yaml).toContain("openapi: 3.0.0");
  expect(parseDraft(yaml)).toMatchObject({ ok: true, title: "From JSON" });
});
