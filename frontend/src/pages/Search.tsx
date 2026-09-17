import { useEffect, useState, type FormEvent } from "react";
import { Link, useSearchParams } from "react-router";
import { useQuery } from "@tanstack/react-query";
import Shell from "@/components/Shell";
import { useCollections } from "@/api/client";
import { operationAnchor } from "@/lib/scalar";

export interface SearchHit {
  id: string;
  collectionId: string;
  collectionSlug: string;
  collectionName: string;
  method: string;
  path: string;
  operationId: string;
  summary: string;
  description: string;
  tags: string[];
  deprecated: boolean;
  score: number;
}

async function search(q: string, collection: string): Promise<SearchHit[]> {
  const params = new URLSearchParams({ q, limit: "50" });
  if (collection) params.set("collection", collection);
  const res = await fetch(`/api/v1/search?${params.toString()}`, { headers: { Accept: "application/json" } });
  if (!res.ok) throw new Error(`search failed (${res.status})`);
  return ((await res.json()) as { hits: SearchHit[] }).hits;
}

const methodColor: Record<string, string> = {
  GET: "bg-blue-100 text-blue-800",
  POST: "bg-green-100 text-green-800",
  PUT: "bg-amber-100 text-amber-800",
  PATCH: "bg-purple-100 text-purple-800",
  DELETE: "bg-red-100 text-red-800",
};

export default function Search() {
  const [params, setParams] = useSearchParams();
  const q = params.get("q") ?? "";
  const collection = params.get("collection") ?? "";
  const [draft, setDraft] = useState(q);
  useEffect(() => setDraft(q), [q]);

  const collections = useCollections();
  const results = useQuery({
    queryKey: ["search", q, collection],
    queryFn: () => search(q, collection),
    enabled: q.trim().length > 0,
  });

  function submit(e: FormEvent) {
    e.preventDefault();
    const next = new URLSearchParams();
    if (draft.trim()) next.set("q", draft.trim());
    if (collection) next.set("collection", collection);
    setParams(next);
  }

  const grouped = new Map<string, SearchHit[]>();
  for (const hit of results.data ?? []) {
    const key = hit.collectionSlug || hit.collectionId;
    grouped.set(key, [...(grouped.get(key) ?? []), hit]);
  }

  return (
    <Shell>
      <h1 className="text-3xl font-semibold tracking-tight">Search endpoints</h1>
      <form onSubmit={submit} className="mt-6 flex flex-wrap gap-2" role="search">
        <input
          type="search"
          value={draft}
          onChange={(e) => setDraft(e.target.value)}
          placeholder="e.g. tracking, create order, awb"
          aria-label="Search query"
          className="min-w-[240px] flex-1 rounded-md border border-gray-300 px-3 py-2"
        />
        <select
          aria-label="Collection"
          value={collection}
          onChange={(e) => {
            const next = new URLSearchParams(params);
            if (e.target.value) next.set("collection", e.target.value);
            else next.delete("collection");
            setParams(next);
          }}
          className="rounded-md border border-gray-300 px-3 py-2"
        >
          <option value="">All collections</option>
          {collections.data?.map((c) => (
            <option key={c.slug} value={c.slug}>
              {c.name}
            </option>
          ))}
        </select>
        <button type="submit" className="rounded-md bg-gray-900 px-4 py-2 text-sm font-medium text-white">
          Search
        </button>
      </form>

      <div className="mt-8">
        {!q && <p className="text-sm text-gray-500">Type a word from an endpoint name, path, or description.</p>}
        {results.isLoading && <p className="text-sm text-gray-500">Searching…</p>}
        {results.isError && (
          <p role="alert" className="rounded-md border border-red-200 bg-red-50 p-3 text-sm text-red-700">
            {results.error.message}
          </p>
        )}
        {results.data && results.data.length === 0 && <p className="text-sm text-gray-500">No endpoints match “{q}”.</p>}
        {[...grouped.entries()].map(([slug, hits]) => (
          <section key={slug} className="mb-8">
            <h2 className="mb-2 text-sm font-semibold uppercase tracking-wide text-gray-500">
              <Link to={`/docs/${slug}`} className="hover:underline">
                {hits[0].collectionName || slug}
              </Link>{" "}
              · {hits.length} result{hits.length === 1 ? "" : "s"}
            </h2>
            <ul className="divide-y divide-gray-100 rounded-xl border border-gray-200">
              {hits.map((hit) => (
                <li key={hit.id}>
                  <Link to={`/docs/${slug}${operationAnchor(hit)}`} className="block px-4 py-3 hover:bg-gray-50">
                    <div className="flex items-center gap-3">
                      <span className={`rounded px-1.5 py-0.5 font-mono text-xs font-semibold ${methodColor[hit.method] ?? "bg-gray-100 text-gray-700"}`}>
                        {hit.method}
                      </span>
                      <span className="font-mono text-sm">{hit.path}</span>
                      {hit.deprecated && <span className="text-xs text-red-600">deprecated</span>}
                    </div>
                    {hit.summary && <p className="mt-1 text-sm text-gray-700">{hit.summary}</p>}
                  </Link>
                </li>
              ))}
            </ul>
          </section>
        ))}
      </div>
    </Shell>
  );
}
