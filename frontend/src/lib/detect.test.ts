import { detectFormat, slugify, slugPattern } from "./detect";

test("detectFormat recognises the three supported inputs", () => {
  expect(detectFormat("openapi: 3.0.4\ninfo:\n  title: x")).toBe("openapi3");
  expect(detectFormat('{"openapi":"3.1.0"}')).toBe("openapi3");
  expect(detectFormat('{"swagger":"2.0","info":{}}')).toBe("openapi2");
  expect(detectFormat('{"info":{"_postman_id":"abc","name":"x"},"item":[]}')).toBe("postman");
  expect(detectFormat('{"info":{"schema":"https://schema.getpostman.com/json/collection/v2.1.0/collection.json"}}')).toBe("postman");
  expect(detectFormat("hello world")).toBe("unknown");
});

test("slugify produces valid slugs from names and filenames", () => {
  for (const [input, want] of [
    ["Shiprocket API", "shiprocket-api"],
    ["petstore-v3.yaml", "petstore-v3"],
    ["Acme.postman_collection.json", "acme"],
    ["  Weird__Name!! ", "weird-name"],
  ]) {
    expect(slugify(input)).toBe(want);
    expect(slugPattern.test(slugify(input))).toBe(true);
  }
});
