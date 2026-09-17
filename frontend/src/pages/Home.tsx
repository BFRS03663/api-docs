import { useCollections } from "@/api/client";
import CollectionCard from "@/components/CollectionCard";
import Shell from "@/components/Shell";

export default function Home() {
  const { data, isLoading, isError, error } = useCollections();
  return (
    <Shell>
      <div className="mb-8">
        <h1 className="text-3xl font-semibold tracking-tight">API reference</h1>
        <p className="mt-2 max-w-2xl text-gray-600">
          Browse every documented API. Each collection includes endpoint reference, code samples, and a request runner.
        </p>
      </div>

      {isLoading && <p className="text-sm text-gray-500">Loading collections…</p>}
      {isError && (
        <p role="alert" className="rounded-md border border-red-200 bg-red-50 p-3 text-sm text-red-700">
          Could not load collections: {error.message}
        </p>
      )}
      {data && data.length === 0 && (
        <div className="rounded-xl border border-dashed border-gray-300 p-10 text-center text-gray-500">
          No collections yet. Import an OpenAPI spec or Postman collection from the admin area.
        </div>
      )}
      {data && data.length > 0 && (
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {data.map((c) => (
            <CollectionCard key={c.id} collection={c} />
          ))}
        </div>
      )}
    </Shell>
  );
}
