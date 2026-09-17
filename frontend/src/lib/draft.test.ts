import { parseDraft, specToYaml, starterTemplate } from "./draft";

test("accepts an OpenAPI 3 document and reads its title", () => {
  const d = parseDraft("openapi: 3.0.3\ninfo:\n  title: Mini\n  version: '2'\npaths:\n  /a: {}\n  /b: {}\n");
  expect(d).toEqual({ ok: true, format: "openapi3", title: "Mini", version: "2", pathCount: 2 });
});

test("accepts Swagger 2.0", () => {
  expect(parseDraft('swagger: "2.0"\ninfo: {title: Old}\npaths: {}\n')).toMatchObject({ ok: true, format: "openapi2", title: "Old" });
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
  expect(parseDraft('{"info":{"_postman_id":"x"},"item":[]}')).toMatchObject({ ok: false, message: /Postman/ });
});

test("starter template is valid and converts JSON to YAML", () => {
  expect(parseDraft(starterTemplate)).toMatchObject({ ok: true, title: "My API", pathCount: 1 });
  const yaml = specToYaml('{"openapi":"3.0.0","info":{"title":"From JSON","version":"1"},"paths":{}}');
  expect(yaml).toContain("openapi: 3.0.0");
  expect(parseDraft(yaml)).toMatchObject({ ok: true, title: "From JSON" });
});
