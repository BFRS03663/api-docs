import { useMemo } from "react";
import { useParams } from "react-router";
import { ApiReferenceReact } from "@scalar/api-reference-react";
import "@scalar/api-reference-react/style.css";
import { ApiError, useCollection } from "@/api/client";
import Shell from "@/components/Shell";
import LlmExportMenu from "@/components/LlmExportMenu";
import { scalarConfig } from "@/lib/scalar";
import NotFound from "./NotFound";

export default function Docs() {
  const { slug } = useParams<{ slug: string }>();
  const { data, isLoading, isError, error } = useCollection(slug);
  const configuration = useMemo(() => (slug ? scalarConfig(slug) : undefined), [slug]);

  if (!slug || (isError && error instanceof ApiError && error.status === 404)) {
    return <NotFound />;
  }
  if (isError) {
    return (
      <Shell>
        <p role="alert" className="rounded-md border border-red-200 bg-red-50 p-3 text-sm text-red-700">
          Could not load this collection: {error.message}
        </p>
      </Shell>
    );
  }
  if (isLoading || !data || !configuration) {
    return (
      <Shell>
        <p className="text-sm text-gray-500">Loading documentation…</p>
      </Shell>
    );
  }
  return (
    <Shell fullBleed>
      <title>{`${data.name} · API Docs`}</title>
      <div className="flex items-center gap-3 border-b border-gray-200 bg-gray-50 px-4 py-2 text-sm">
        <span className="font-medium">{data.name}</span>
        <span className="text-gray-500">{data.operationCount} endpoints</span>
        <div className="ml-auto">
          <LlmExportMenu slug={slug} />
        </div>
      </div>
      <ApiReferenceReact configuration={configuration} />
    </Shell>
  );
}
