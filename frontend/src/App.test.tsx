import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import App from "./App";
import { clearSession, setSession } from "@/lib/auth";

// Scalar pulls in a Vue runtime and CSS that jsdom cannot execute; the page
// that embeds it is exercised by the pure config test and manual runs.
vi.mock("@scalar/api-reference-react", () => ({
  ApiReferenceReact: () => <div data-testid="scalar" />,
}));

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

const petstore = {
  id: "c1",
  slug: "petstore",
  name: "Petstore",
  description: "Pets as a service",
  version: "1.0.0",
  source: { type: "openapi3", filename: "petstore.yaml", importedAt: "2026-09-17T00:00:00Z" },
  servers: [],
  operationCount: 19,
  createdAt: "",
  updatedAt: "",
};

type Route = { status: number; body: unknown };
type Handler = Route | ((req: Request) => Route);

function mockFetch(routes: Record<string, Handler>) {
  const calls: Request[] = [];
  vi.stubGlobal(
    "fetch",
    vi.fn(async (input: RequestInfo | URL, init?: RequestInit) => {
      const url = new URL(typeof input === "string" ? input : input.toString(), "http://test");
      const req = new Request(url.toString(), init);
      calls.push(req);
      const key = `${req.method} ${url.pathname}`;
      const handler = routes[key] ?? routes[url.pathname];
      const hit = typeof handler === "function" ? handler(req) : (handler ?? { status: 404, body: { error: "not found" } });
      return new Response(hit.status === 204 ? null : JSON.stringify(hit.body), {
        status: hit.status,
        headers: { "Content-Type": "application/json" },
      });
    }),
  );
  return calls;
}

beforeEach(() => clearSession());

test("home lists collections as cards linking to docs", async () => {
  mockFetch({ "/api/v1/collections": { status: 200, body: { collections: [petstore] } } });
  renderAt("/");
  expect(await screen.findByRole("heading", { name: "Petstore" })).toBeInTheDocument();
  expect(screen.getByRole("link", { name: /Petstore/ })).toHaveAttribute("href", "/docs/petstore");
  expect(screen.getByText("OpenAPI 3")).toBeInTheDocument();
});

test("home shows empty state", async () => {
  mockFetch({ "/api/v1/collections": { status: 200, body: { collections: [] } } });
  renderAt("/");
  expect(await screen.findByText(/No collections yet/)).toBeInTheDocument();
});

test("docs page renders Scalar for a known slug", async () => {
  mockFetch({ "/api/v1/collections/petstore": { status: 200, body: petstore } });
  renderAt("/docs/petstore");
  expect(await screen.findByTestId("scalar")).toBeInTheDocument();
});

test("docs page shows not found for an unknown slug", async () => {
  mockFetch({});
  renderAt("/docs/nope");
  expect(await screen.findByRole("heading", { name: "Page not found" })).toBeInTheDocument();
});

test("admin routes redirect to login when signed out", () => {
  mockFetch({});
  renderAt("/admin");
  expect(screen.getByRole("heading", { name: "Admin sign in" })).toBeInTheDocument();
});

test("login stores the session and shows the collections table", async () => {
  mockFetch({
    "POST /api/v1/auth/login": (req) => {
      const auth = req.headers.get("content-type");
      return auth?.includes("application/json")
        ? { status: 200, body: { token: "t0k", expiresAt: new Date(Date.now() + 3600_000).toISOString() } }
        : { status: 400, body: { error: "bad" } };
    },
    "/api/v1/collections": { status: 200, body: { collections: [petstore] } },
  });
  renderAt("/admin/login");
  fireEvent.change(screen.getByLabelText("Username"), { target: { value: "admin" } });
  fireEvent.change(screen.getByLabelText("Password"), { target: { value: "secret" } });
  fireEvent.submit(screen.getByRole("form", { name: "sign in" }));
  expect(await screen.findByRole("heading", { name: "Admin" })).toBeInTheDocument();
  expect(await screen.findByText("petstore")).toBeInTheDocument();
  expect(screen.getByRole("link", { name: "Re-import" })).toHaveAttribute("href", "/admin/import/petstore");
});

test("login failure shows the server message", async () => {
  mockFetch({ "POST /api/v1/auth/login": { status: 401, body: { error: "invalid credentials" } } });
  renderAt("/admin/login");
  fireEvent.change(screen.getByLabelText("Username"), { target: { value: "admin" } });
  fireEvent.change(screen.getByLabelText("Password"), { target: { value: "nope" } });
  fireEvent.submit(screen.getByRole("form", { name: "sign in" }));
  expect(await screen.findByRole("alert")).toHaveTextContent("invalid credentials");
});

test("import page uploads a file with the bearer token and shows the result", async () => {
  setSession("t0k", new Date(Date.now() + 3600_000).toISOString());
  const calls = mockFetch({
    "POST /api/v1/admin/collections": { status: 201, body: { created: true, collection: { ...petstore, slug: "acme", name: "Acme" } } },
    "/api/v1/collections": { status: 200, body: { collections: [] } },
  });
  renderAt("/admin/import");
  const file = new File(['{"openapi":"3.0.0"}'], "acme-api.json", { type: "application/json" });
  fireEvent.change(screen.getByTestId("file-input"), { target: { files: [file] } });
  await waitFor(() => expect(screen.getByDisplayValue("acme-api")).toBeInTheDocument());
  expect(await screen.findByText("OpenAPI 3")).toBeInTheDocument();
  fireEvent.submit(screen.getByRole("form", { name: "import" }));
  expect(await screen.findByRole("status")).toHaveTextContent("Created Acme with 19 endpoints");
  const upload = calls.find((c) => c.method === "POST" && c.url.includes("/admin/collections"));
  expect(upload?.headers.get("authorization")).toBe("Bearer t0k");
});

test("search page lists hits grouped by collection with Scalar anchors", async () => {
  mockFetch({
    "/api/v1/collections": { status: 200, body: { collections: [petstore] } },
    "/api/v1/search": {
      status: 200,
      body: {
        hits: [
          { id: "o1", collectionId: "c1", collectionSlug: "petstore", collectionName: "Petstore", method: "GET", path: "/pet/findByStatus", operationId: "findPetsByStatus", summary: "Finds Pets by status", description: "", tags: ["pet"], deprecated: false, score: 1 },
        ],
      },
    },
  });
  renderAt("/search?q=status");
  expect(await screen.findByText("/pet/findByStatus")).toBeInTheDocument();
  expect(screen.getByRole("link", { name: /GET.*findByStatus/ })).toHaveAttribute("href", "/docs/petstore#tag/pet/GET/pet/findByStatus");
  expect(screen.getByText(/1 result/)).toBeInTheDocument();
});

test("docs page offers LLM exports", async () => {
  mockFetch({ "/api/v1/collections/petstore": { status: 200, body: petstore } });
  renderAt("/docs/petstore");
  await screen.findByTestId("scalar");
  fireEvent.click(screen.getByRole("button", { name: /Export for LLMs/ }));
  expect(screen.getByRole("menuitem", { name: "Download Markdown (.md)" })).toHaveAttribute("href", "/docs/petstore.md");
  expect(screen.getByRole("menuitem", { name: "OpenAPI document (JSON)" })).toHaveAttribute("href", "/docs/petstore/openapi.json");
});

test("unknown route shows not found", () => {
  mockFetch({});
  renderAt("/nope");
  expect(screen.getByRole("heading", { name: "Page not found" })).toBeInTheDocument();
});
