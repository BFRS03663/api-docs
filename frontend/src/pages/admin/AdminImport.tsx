import { useEffect, useState, type DragEvent, type FormEvent } from "react";
import { Link, useParams } from "react-router";
import { useCollection } from "@/api/client";
import { useImportCollection, type ImportSource } from "@/api/admin";
import { detectFormat, formatLabel, slugify, slugPattern, type DetectedFormat } from "@/lib/detect";
import { site } from "@/lib/site";

type Mode = "file" | "url";

/** Reads the first 64 KB of a file as text (FileReader works in every browser and in jsdom). */
function readHead(f: File): Promise<string> {
  return new Promise((resolve) => {
    const reader = new FileReader();
    reader.onload = () => resolve(typeof reader.result === "string" ? reader.result : "");
    reader.onerror = () => resolve("");
    reader.readAsText(f.slice(0, 64 * 1024));
  });
}

export default function AdminImport() {
  const { slug: existingSlug } = useParams<{ slug: string }>();
  const existing = useCollection(existingSlug);
  const importer = useImportCollection();

  const [mode, setMode] = useState<Mode>("file");
  const [file, setFile] = useState<File | null>(null);
  const [detected, setDetected] = useState<DetectedFormat | null>(null);
  const [url, setUrl] = useState("");
  const [slug, setSlug] = useState(existingSlug ?? "");
  const [slugTouched, setSlugTouched] = useState(Boolean(existingSlug));
  const [name, setName] = useState("");
  const [dragging, setDragging] = useState(false);

  useEffect(() => {
    if (existing.data) setName(existing.data.name);
  }, [existing.data]);

  async function pick(f: File | null) {
    setFile(f);
    setDetected(null);
    if (!f) return;
    if (!slugTouched) setSlug(slugify(f.name));
    setDetected(detectFormat(await readHead(f)));
  }

  function onDrop(e: DragEvent<HTMLDivElement>) {
    e.preventDefault();
    setDragging(false);
    void pick(e.dataTransfer.files[0] ?? null);
  }

  function onUrlChange(v: string) {
    setUrl(v);
    if (!slugTouched) {
      try {
        setSlug(slugify(new URL(v).hostname.replace(/^(www|apidocs|docs|api)\./, "")));
      } catch {
        // not a URL yet
      }
    }
  }

  const slugOk = slugPattern.test(slug);
  const source: ImportSource | null = mode === "file" ? (file ? { kind: "file", file } : null) : url ? { kind: "url", url } : null;
  const canSubmit = Boolean(source) && slugOk && !importer.isPending;

  function onSubmit(e: FormEvent) {
    e.preventDefault();
    if (!source) return;
    importer.mutate({ source, slug, name, existing: Boolean(existingSlug) });
  }

  const input = "mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-sm";
  const tabBtn = (active: boolean) =>
    `rounded-md px-3 py-1.5 text-sm font-medium ${active ? "bg-gray-900 text-white" : "bg-gray-100 text-gray-700"}`;

  return (
    <section className="max-w-2xl">
      <h2 className="text-lg font-semibold">{existingSlug ? `Re-import "${existingSlug}"` : "Import a collection"}</h2>
      <p className="mt-1 text-sm text-gray-600">
        Upload an OpenAPI 2.0/3.x or Postman v2.1 file, or paste the URL of a spec or of a Postman published-docs page.
      </p>

      <form onSubmit={onSubmit} className="mt-6 space-y-5" aria-label="import">
        <div className="flex gap-2">
          <button type="button" className={tabBtn(mode === "file")} onClick={() => setMode("file")}>
            Upload file
          </button>
          <button type="button" className={tabBtn(mode === "url")} onClick={() => setMode("url")}>
            From URL
          </button>
        </div>

        {mode === "file" ? (
          <div
            onDragOver={(e) => {
              e.preventDefault();
              setDragging(true);
            }}
            onDragLeave={() => setDragging(false)}
            onDrop={onDrop}
            className={`rounded-xl border-2 border-dashed p-8 text-center ${dragging ? "border-gray-900 bg-gray-50" : "border-gray-300"}`}
          >
            <p className="text-sm text-gray-600">Drag a file here, or</p>
            <label className="mt-2 inline-block cursor-pointer rounded-md bg-gray-100 px-3 py-1.5 text-sm font-medium">
              choose a file
              <input
                type="file"
                accept=".json,.yaml,.yml,application/json,application/yaml"
                className="sr-only"
                data-testid="file-input"
                onChange={(e) => void pick(e.target.files?.[0] ?? null)}
              />
            </label>
            {file && (
              <p className="mt-3 text-sm">
                <span className="font-mono">{file.name}</span> · {(file.size / 1024).toFixed(0)} KB
                {detected && (
                  <span
                    className={`ml-2 rounded-full px-2 py-0.5 text-xs font-medium ${detected === "unknown" ? "bg-red-100 text-red-700" : "bg-green-100 text-green-800"}`}
                  >
                    {formatLabel[detected]}
                  </span>
                )}
              </p>
            )}
          </div>
        ) : (
          <label className="block text-sm">
            <span className="font-medium">URL</span>
            <input className={input} type="url" placeholder="https://apidocs.example.com/" value={url} onChange={(e) => onUrlChange(e.target.value)} />
            <span className="mt-1 block text-xs text-gray-500">Postman published-docs pages are detected and their collection is fetched automatically.</span>
          </label>
        )}

        <div className="grid gap-4 sm:grid-cols-2">
          <label className="block text-sm">
            <span className="font-medium">Slug</span>
            <input
              className={`${input} font-mono ${slug && !slugOk ? "border-red-400" : ""}`}
              value={slug}
              disabled={Boolean(existingSlug)}
              onChange={(e) => {
                setSlugTouched(true);
                setSlug(e.target.value);
              }}
              placeholder="my-api"
            />
            <span className="mt-1 block text-xs text-gray-500">Lowercase letters, digits and hyphens. Appears in the URL: /docs/{slug || "my-api"}</span>
          </label>
          <label className="block text-sm">
            <span className="font-medium">Display name</span>
            <input className={input} value={name} onChange={(e) => setName(e.target.value)} placeholder="Defaults to the spec title" />
          </label>
        </div>

        {importer.isError && (
          <p role="alert" className="rounded-md border border-red-200 bg-red-50 p-3 text-sm text-red-700">
            {importer.error.message}
          </p>
        )}
        {importer.isSuccess && (
          <p role="status" className="rounded-md border border-green-200 bg-green-50 p-3 text-sm text-green-800">
            {importer.data.created ? "Created" : "Updated"} <strong>{importer.data.collection.name}</strong> with{" "}
            {importer.data.collection.operationCount} endpoints.{" "}
            <Link to={`/docs/${importer.data.collection.slug}`} className="underline">
              Open documentation
            </Link>
          </p>
        )}

        <button
          type="submit"
          disabled={!canSubmit}
          className="rounded-md px-4 py-2 text-sm font-medium text-white disabled:opacity-50"
          style={{ background: site.accent }}
        >
          {importer.isPending ? "Importing…" : existingSlug ? "Re-import" : "Import"}
        </button>
      </form>
    </section>
  );
}
