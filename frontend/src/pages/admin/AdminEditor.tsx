import { useEffect, useMemo, useState, type FormEvent } from "react";
import { Link, useNavigate, useParams } from "react-router";
import { ApiReferenceReact } from "@scalar/api-reference-react";
import "@scalar/api-reference-react/style.css";
import { useCollectionSource, useImportCollection } from "@/api/admin";
import { useCollection } from "@/api/client";
import YamlEditor from "@/components/YamlEditor";
import { formatLabel, slugify, slugPattern } from "@/lib/detect";
import { exampleTemplate, parseDraft, skeletonTemplate, specToYaml } from "@/lib/draft";
import { scalarPreviewConfig } from "@/lib/scalar";
import { site } from "@/lib/site";

/** Pause after the last keystroke before the preview re-renders. */
const previewDelayMs = 400;

function useDebounced<T>(value: T, delay: number): T {
  const [debounced, setDebounced] = useState(value);
  useEffect(() => {
    const timer = setTimeout(() => setDebounced(value), delay);
    return () => clearTimeout(timer);
  }, [value, delay]);
  return debounced;
}

/**
 * Swagger-Editor style page: the OpenAPI document on the left, the rendered
 * reference on the right, updated as you type. Publishing sends the text
 * through the same import endpoint as a file upload.
 */
export default function AdminEditor() {
  const { slug: routeSlug } = useParams<{ slug: string }>();
  const navigate = useNavigate();
  const importer = useImportCollection();
  const collection = useCollection(routeSlug);

  // A new document starts as the skeleton; a stored one is loaded below.
  const [text, setText] = useState(routeSlug ? "" : skeletonTemplate);
  const [loadedSlug, setLoadedSlug] = useState<string | null>(null);
  const [publishedText, setPublishedText] = useState<string | null>(null);
  const [slug, setSlug] = useState(routeSlug ?? "");
  const [slugTouched, setSlugTouched] = useState(Boolean(routeSlug));
  const [name, setName] = useState("");

  // The stored document is fetched once per slug; a publish from this page
  // records the slug so the just-sent text is not reloaded over the editor.
  const source = useCollectionSource(routeSlug, loadedSlug !== routeSlug);
  useEffect(() => {
    if (!routeSlug || !source.data || loadedSlug === routeSlug) return;
    const initial = source.data.canonical ? specToYaml(source.data.content) : source.data.content;
    setText(initial);
    setPublishedText(initial);
    setLoadedSlug(routeSlug);
  }, [routeSlug, source.data, loadedSlug]);
  useEffect(() => {
    if (collection.data) setName(collection.data.name);
  }, [collection.data]);

  const debounced = useDebounced(text, previewDelayMs);
  const draft = useMemo(() => parseDraft(debounced), [debounced]);
  // While the user is mid-edit the preview keeps showing the last document that parsed.
  const [preview, setPreview] = useState<string | null>(null);
  useEffect(() => {
    if (draft.ok) setPreview(debounced);
  }, [draft, debounced]);
  useEffect(() => {
    if (draft.ok && !slugTouched) setSlug(slugify(draft.title));
  }, [draft, slugTouched]);
  const configuration = useMemo(() => (preview ? scalarPreviewConfig(preview) : null), [preview]);

  const baseline = publishedText ?? (routeSlug ? "" : skeletonTemplate);
  const dirty = text !== baseline;
  const untouched = !routeSlug && (text === skeletonTemplate || text.trim() === "");
  useEffect(() => {
    if (!dirty) return;
    const warn = (e: BeforeUnloadEvent) => e.preventDefault();
    window.addEventListener("beforeunload", warn);
    return () => window.removeEventListener("beforeunload", warn);
  }, [dirty]);

  function onChange(next: string) {
    setText(next);
    if (importer.isSuccess || importer.isError) importer.reset();
  }

  const slugOk = slugPattern.test(slug);
  const canPublish = draft.ok && slugOk && !importer.isPending;

  function onPublish(e: FormEvent) {
    e.preventDefault();
    if (!parseDraft(text).ok || !slugOk) return;
    const existing = Boolean(routeSlug);
    const file = new File([text], `${slug}.yaml`, { type: "application/yaml" });
    importer.mutate(
      { source: { kind: "file", file }, slug, name, existing },
      {
        onSuccess: (result) => {
          setPublishedText(text);
          if (!existing) {
            setLoadedSlug(result.collection.slug);
            navigate(`/admin/editor/${result.collection.slug}`, { replace: true });
          }
        },
      },
    );
  }

  const input = "rounded-md border border-gray-300 px-2 py-1 text-sm";
  const alert = "rounded-md border border-red-200 bg-red-50 px-3 py-2 text-sm text-red-700";
  const showBanners = source.isError || source.data?.canonical || (!draft.ok && text.trim()) || importer.isError || importer.isSuccess;

  return (
    <section className="flex flex-col" style={{ height: "calc(100vh - 7.75rem)" }}>
      <form onSubmit={onPublish} aria-label="publish" className="flex flex-wrap items-center gap-3 border-b border-gray-200 px-4 py-2">
        <h2 className="text-base font-semibold">{routeSlug ? `Edit "${routeSlug}"` : "Write a spec"}</h2>
        <label className="flex items-center gap-1 text-sm">
          <span className="text-gray-600">Slug</span>
          <input
            className={`${input} w-40 font-mono ${slug && !slugOk ? "border-red-400" : ""}`}
            aria-label="Slug"
            value={slug}
            disabled={Boolean(routeSlug)}
            placeholder="my-api"
            onChange={(e) => {
              setSlugTouched(true);
              setSlug(e.target.value);
            }}
          />
        </label>
        <label className="flex items-center gap-1 text-sm">
          <span className="text-gray-600">Display name</span>
          <input className={`${input} w-48`} aria-label="Display name" value={name} onChange={(e) => setName(e.target.value)} placeholder="Defaults to the title" />
        </label>
        {draft.ok && (
          <span className="rounded-full bg-green-100 px-2 py-0.5 text-xs font-medium text-green-800">
            {formatLabel[draft.format]} · {draft.pathCount} {draft.pathCount === 1 ? "path" : "paths"}
          </span>
        )}
        <div className="ml-auto flex items-center gap-2">
          {untouched && (
            <button type="button" onClick={() => onChange(exampleTemplate)} className="rounded-md bg-gray-100 px-3 py-1.5 text-sm font-medium text-gray-700">
              Load a full example
            </button>
          )}
          <button type="submit" disabled={!canPublish} className="rounded-md px-4 py-1.5 text-sm font-medium text-white disabled:opacity-50" style={{ background: site.accent }}>
            {importer.isPending ? "Publishing…" : routeSlug ? "Publish changes" : "Publish"}
          </button>
        </div>
      </form>

      {showBanners && (
        <div className="space-y-1 px-4 pt-2">
          {source.isError && (
            <p role="alert" className={alert}>
              Could not load the stored document: {source.error.message}
            </p>
          )}
          {source.data?.canonical && (
            <p className="rounded-md border border-blue-200 bg-blue-50 px-3 py-2 text-sm text-blue-800">
              This collection was imported from Postman. You are editing the converted OpenAPI document; publishing stores it as OpenAPI 3.
            </p>
          )}
          {!draft.ok && text.trim() && (
            <p role="alert" className={alert}>
              {draft.message}
            </p>
          )}
          {importer.isError && (
            <p role="alert" className={alert}>
              Publish failed: {importer.error.message}
            </p>
          )}
          {importer.isSuccess && (
            <p role="status" className="rounded-md border border-green-200 bg-green-50 px-3 py-2 text-sm text-green-800">
              {importer.data.created ? "Published" : "Updated"} <strong>{importer.data.collection.name}</strong> with {importer.data.collection.operationCount}{" "}
              endpoints.{" "}
              <Link to={`/docs/${importer.data.collection.slug}`} className="underline">
                Open documentation
              </Link>
            </p>
          )}
        </div>
      )}

      <div className="mt-2 grid min-h-0 flex-1 grid-cols-1 border-t border-gray-200 lg:grid-cols-2">
        <div className="min-h-0 overflow-hidden border-r border-gray-200">
          {routeSlug && source.isLoading ? <p className="p-4 text-sm text-gray-500">Loading document…</p> : <YamlEditor value={text} onChange={onChange} />}
        </div>
        <div className="min-h-0 overflow-y-auto bg-white" data-testid="preview">
          {configuration ? (
            <ApiReferenceReact configuration={configuration} />
          ) : (
            <p className="p-8 text-center text-sm text-gray-500">The rendered documentation appears here as soon as the document parses.</p>
          )}
        </div>
      </div>
    </section>
  );
}
