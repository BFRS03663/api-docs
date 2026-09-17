import { Link } from "react-router";
import { useCollections } from "@/api/client";
import { useDeleteCollection } from "@/api/admin";

const sourceLabel: Record<string, string> = { openapi2: "Swagger 2.0", openapi3: "OpenAPI 3", postman: "Postman" };

export default function AdminCollections() {
  const { data, isLoading, isError, error } = useCollections();
  const remove = useDeleteCollection();

  function onDelete(slug: string, name: string) {
    if (!window.confirm(`Delete "${name}" (${slug})? This removes the collection and its endpoints.`)) return;
    remove.mutate(slug);
  }

  return (
    <section>
      {isLoading && <p className="text-sm text-gray-500">Loading…</p>}
      {isError && (
        <p role="alert" className="rounded-md border border-red-200 bg-red-50 p-3 text-sm text-red-700">
          {error.message}
        </p>
      )}
      {remove.isError && (
        <p role="alert" className="mb-3 rounded-md border border-red-200 bg-red-50 p-3 text-sm text-red-700">
          Delete failed: {remove.error.message}
        </p>
      )}
      {data && data.length === 0 && (
        <p className="rounded-xl border border-dashed border-gray-300 p-8 text-center text-gray-500">
          No collections yet. <Link to="/admin/import" className="underline">Import one</Link>.
        </p>
      )}
      {data && data.length > 0 && (
        <div className="overflow-x-auto rounded-xl border border-gray-200">
          <table className="w-full text-left text-sm">
            <thead className="bg-gray-50 text-xs uppercase tracking-wide text-gray-500">
              <tr>
                <th className="px-4 py-3">Name</th>
                <th className="px-4 py-3">Slug</th>
                <th className="px-4 py-3">Source</th>
                <th className="px-4 py-3">Endpoints</th>
                <th className="px-4 py-3">Imported</th>
                <th className="px-4 py-3 text-right">Actions</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-100">
              {data.map((c) => (
                <tr key={c.id}>
                  <td className="px-4 py-3 font-medium">
                    <Link to={`/docs/${c.slug}`} className="hover:underline">
                      {c.name}
                    </Link>
                  </td>
                  <td className="px-4 py-3 font-mono text-xs">{c.slug}</td>
                  <td className="px-4 py-3">
                    {sourceLabel[c.source.type] ?? c.source.type}
                    <span className="block text-xs text-gray-500">{c.source.filename}</span>
                  </td>
                  <td className="px-4 py-3">{c.operationCount}</td>
                  <td className="px-4 py-3 text-gray-600">{new Date(c.source.importedAt).toLocaleString()}</td>
                  <td className="px-4 py-3 text-right whitespace-nowrap">
                    <Link to={`/admin/import/${c.slug}`} className="mr-3 text-gray-700 hover:underline">
                      Re-import
                    </Link>
                    <button
                      onClick={() => onDelete(c.slug, c.name)}
                      disabled={remove.isPending}
                      className="text-red-600 hover:underline disabled:opacity-50"
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </section>
  );
}
