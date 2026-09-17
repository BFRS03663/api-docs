import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import type { ApiReferenceConfiguration } from "@scalar/api-reference-react";
import App from "@/App";
import { clearSession, setSession } from "@/lib/auth";

// Scalar runs a Vue app jsdom cannot execute; the stub prints the inline
// document it was given so tests can see what the preview would render.
vi.mock("@scalar/api-reference-react", () => ({
  ApiReferenceReact: ({ configuration }: { configuration: Partial<ApiReferenceConfiguration> }) => (
    <pre data-testid="scalar">{String(configuration.content ?? "")}</pre>
  ),
}));

// CodeMirror needs layout APIs jsdom lacks; a textarea stands in for it.
vi.mock("@uiw/react-codemirror", () => ({
  default: ({ value, onChange }: { value: string; onChange: (next: string) => void }) => (
    <textarea aria-label="YAML" value={value} onChange={(e) => onChange(e.target.value)} />
  ),
  EditorView: { lineWrapping: [] },
}));

const MINI = "openapi: 3.0.3\ninfo:\n  title: Mini\n  version: 1.0.0\npaths:\n  /ping:\n    get:\n      responses:\n        '200':\n          description: ok\n";

const petstore = {
  id: "c1",
  slug: "petstore",
  name: "Petstore",
  description: "",
  version: "1.0.0",
  source: { type: "openapi3", filename: "petstore.yaml", importedAt: "2026-09-17T00:00:00Z" },
  servers: [],
  operationCount: 1,
  createdAt: "",
  updatedAt: "",
};

type Hit = { status: number; body: unknown };
type Call = { method: string; path: string; headers: Headers; body: BodyInit | null | undefined };

function mockFetch(routes: Record<string, Hit>) {
  const calls: Call[] = [];
  vi.stubGlobal(
    "fetch",
    vi.fn(async (input: RequestInfo | URL, init?: RequestInit) => {
      const url = new URL(typeof input === "string" ? input : input.toString(), "http://test");
      const method = init?.method ?? "GET";
      calls.push({ method, path: url.pathname, headers: new Headers(init?.headers), body: init?.body });
      const hit = routes[`${method} ${url.pathname}`] ?? routes[url.pathname] ?? { status: 404, body: { error: "not found" } };
      return new Response(JSON.stringify(hit.body), { status: hit.status, headers: { "Content-Type": "application/json" } });
    }),
  );
  return calls;
}

function readFile(file: File): Promise<string> {
  return new Promise((resolve) => {
    const reader = new FileReader();
    reader.onload = () => resolve(String(reader.result));
    reader.readAsText(file);
  });
}

function renderAt(path: string) {
  const qc = new QueryClient({ defaultOptions: { queries: { retry: false }, mutations: { retry: false } } });
  return render(
    <QueryClientProvider client={qc}>
      <MemoryRouter initialEntries={[path]}>
        <App />
      </MemoryRouter>
    </QueryClientProvider>,
  );
}

beforeEach(() => {
  clearSession();
  setSession("t0k", new Date(Date.now() + 3600_000).toISOString());
});

test("a new document starts as a skeleton with the top-level sections filled in", async () => {
  mockFetch({});
  renderAt("/admin/editor");

  const editor = (await screen.findByLabelText("YAML")) as HTMLTextAreaElement;
  for (const key of ["openapi: 3.0.3", "title: My API", "version: 1.0.0", "servers:", "tags:", "security:", "paths:", "securitySchemes:"]) {
    expect(editor.value).toContain(key);
  }
  expect(await screen.findByTestId("scalar")).toHaveTextContent("title: My API");
  expect(screen.getByLabelText("Slug")).toHaveValue("my-api");
  expect(editor.value).toContain("/health:");
  expect(screen.getByText(/OpenAPI 3 · 1 path/)).toBeInTheDocument();

  fireEvent.click(screen.getByRole("button", { name: "Load a full example" }));
  expect(editor.value).toContain("/orders:");
  expect(screen.queryByRole("button", { name: "Load a full example" })).not.toBeInTheDocument();
});

test("renders a live preview from typed YAML and publishes it as a new collection", async () => {
  const calls = mockFetch({
    "POST /api/v1/admin/collections": { status: 201, body: { created: true, collection: { ...petstore, slug: "mini", name: "Mini" } } },
  });
  renderAt("/admin/editor");

  fireEvent.change(await screen.findByLabelText("YAML"), { target: { value: MINI } });
  await waitFor(() => expect(screen.getByTestId("scalar")).toHaveTextContent("title: Mini"));
  expect(screen.getByLabelText("Slug")).toHaveValue("mini");
  expect(screen.getByText(/OpenAPI 3 · 1 path/)).toBeInTheDocument();

  fireEvent.click(screen.getByRole("button", { name: "Publish" }));
  expect(await screen.findByRole("status")).toHaveTextContent("Published Mini with 1 endpoints");
  expect(screen.getByRole("link", { name: "Open documentation" })).toHaveAttribute("href", "/docs/mini");

  const upload = calls.find((c) => c.method === "POST" && c.path === "/api/v1/admin/collections");
  expect(upload?.headers.get("authorization")).toBe("Bearer t0k");
  const form = upload?.body as FormData;
  expect(form.get("slug")).toBe("mini");
  expect(await readFile(form.get("file") as File)).toBe(MINI);
  // The page now edits the published collection without refetching it.
  expect(screen.getByLabelText("Slug")).toBeDisabled();
  expect(calls.some((c) => c.path.endsWith("/source"))).toBe(false);
});

test("keeps the last good preview and blocks publishing while the YAML is invalid", async () => {
  mockFetch({});
  renderAt("/admin/editor");
  const editor = await screen.findByLabelText("YAML");

  fireEvent.change(editor, { target: { value: MINI } });
  await waitFor(() => expect(screen.getByTestId("scalar")).toHaveTextContent("title: Mini"));

  fireEvent.change(editor, { target: { value: "openapi: 3.0.3\ninfo: title: x\n" } });
  expect(await screen.findByRole("alert")).toHaveTextContent(/line 2/i);
  expect(screen.getByTestId("scalar")).toHaveTextContent("title: Mini");
  expect(screen.getByRole("button", { name: "Publish" })).toBeDisabled();
});

test("loads an existing collection and publishes changes with PUT", async () => {
  const calls = mockFetch({
    "/api/v1/admin/collections/petstore/source": { status: 200, body: { slug: "petstore", filename: "petstore.yaml", type: "openapi3", canonical: false, content: MINI } },
    "/api/v1/collections/petstore": { status: 200, body: petstore },
    "PUT /api/v1/admin/collections/petstore": { status: 200, body: { created: false, collection: petstore } },
  });
  renderAt("/admin/editor/petstore");

  expect(await screen.findByLabelText("YAML")).toHaveValue(MINI);
  expect(screen.getByLabelText("Slug")).toHaveValue("petstore");
  expect(screen.getByLabelText("Slug")).toBeDisabled();
  await waitFor(() => expect(screen.getByLabelText("Display name")).toHaveValue("Petstore"));

  const publish = screen.getByRole("button", { name: "Publish changes" });
  await waitFor(() => expect(publish).toBeEnabled());
  fireEvent.click(publish);
  expect(await screen.findByRole("status")).toHaveTextContent("Updated Petstore");
  expect(calls.some((c) => c.method === "PUT" && c.path === "/api/v1/admin/collections/petstore")).toBe(true);
});

test("opens a Postman-sourced collection as its converted OpenAPI YAML", async () => {
  mockFetch({
    "/api/v1/admin/collections/legacy/source": {
      status: 200,
      body: { slug: "legacy", filename: "legacy.postman_collection.json", type: "postman", canonical: true, content: '{"openapi":"3.0.0","info":{"title":"Legacy","version":"1"},"paths":{"/x":{}}}' },
    },
  });
  renderAt("/admin/editor/legacy");

  const editor = (await screen.findByLabelText("YAML")) as HTMLTextAreaElement;
  await waitFor(() => expect(editor.value).toContain("openapi: 3.0.0"));
  expect(screen.getByText(/imported from Postman/)).toBeInTheDocument();
  expect(await screen.findByTestId("scalar")).toHaveTextContent("title: Legacy");
});
