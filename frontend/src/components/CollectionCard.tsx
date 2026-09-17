import { Link } from "react-router";
import type { Collection } from "@/api/client";

const sourceLabel: Record<Collection["source"]["type"], string> = {
  openapi2: "Swagger 2.0",
  openapi3: "OpenAPI 3",
  postman: "Postman",
};

export default function CollectionCard({ collection }: { collection: Collection }) {
  const imported = new Date(collection.source.importedAt);
  return (
    <Link
      to={`/docs/${collection.slug}`}
      className="group flex flex-col rounded-xl border border-gray-200 p-5 transition-shadow hover:shadow-md focus:outline-none focus-visible:ring-2 focus-visible:ring-gray-900"
    >
      <div className="flex items-start justify-between gap-3">
        <h2 className="text-lg font-semibold group-hover:underline">{collection.name}</h2>
        <span className="shrink-0 rounded-full bg-gray-100 px-2 py-0.5 text-xs font-medium text-gray-700">
          {sourceLabel[collection.source.type] ?? collection.source.type}
        </span>
      </div>
      {collection.description && <p className="mt-2 line-clamp-3 text-sm text-gray-600">{collection.description}</p>}
      <dl className="mt-4 flex flex-wrap gap-x-5 gap-y-1 text-xs text-gray-500">
        <div>
          <dt className="inline">Version </dt>
          <dd className="inline font-mono text-gray-700">{collection.version || "n/a"}</dd>
        </div>
        <div>
          <dt className="inline">Endpoints </dt>
          <dd className="inline font-mono text-gray-700">{collection.operationCount}</dd>
        </div>
        <div>
          <dt className="inline">Imported </dt>
          <dd className="inline text-gray-700">{isNaN(imported.getTime()) ? "unknown" : imported.toLocaleDateString()}</dd>
        </div>
      </dl>
    </Link>
  );
}
