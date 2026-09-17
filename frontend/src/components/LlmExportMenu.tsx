import { useEffect, useRef, useState } from "react";

/**
 * Links to the machine-readable renderings of a collection: Markdown, plain
 * text, the OpenAPI document, and the site-wide llms.txt. "Copy as Markdown"
 * fetches the .md rendering and puts it on the clipboard.
 */
export default function LlmExportMenu({ slug }: { slug: string }) {
  const [open, setOpen] = useState(false);
  const [copied, setCopied] = useState<"idle" | "copying" | "done" | "failed">("idle");
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!open) return;
    function onClick(e: MouseEvent) {
      if (ref.current && !ref.current.contains(e.target as Node)) setOpen(false);
    }
    document.addEventListener("mousedown", onClick);
    return () => document.removeEventListener("mousedown", onClick);
  }, [open]);

  async function copyMarkdown() {
    setCopied("copying");
    try {
      const res = await fetch(`/docs/${encodeURIComponent(slug)}.md`);
      if (!res.ok) throw new Error(String(res.status));
      await navigator.clipboard.writeText(await res.text());
      setCopied("done");
    } catch {
      setCopied("failed");
    } finally {
      setTimeout(() => setCopied("idle"), 2000);
    }
  }

  const base = `/docs/${encodeURIComponent(slug)}`;
  const item = "block w-full px-3 py-2 text-left text-sm hover:bg-gray-50";

  return (
    <div ref={ref} className="relative">
      <button
        type="button"
        onClick={() => setOpen((v) => !v)}
        aria-haspopup="menu"
        aria-expanded={open}
        className="rounded-md border border-gray-300 bg-white px-3 py-1.5 text-sm font-medium text-gray-700 hover:bg-gray-50"
      >
        Export for LLMs ▾
      </button>
      {open && (
        <div role="menu" className="absolute right-0 z-30 mt-1 w-64 overflow-hidden rounded-md border border-gray-200 bg-white shadow-lg">
          <button role="menuitem" className={item} onClick={copyMarkdown} disabled={copied === "copying"}>
            {copied === "done" ? "Copied!" : copied === "failed" ? "Copy failed" : copied === "copying" ? "Copying…" : "Copy as Markdown"}
          </button>
          <a role="menuitem" className={item} href={`${base}.md`} download={`${slug}.md`}>
            Download Markdown (.md)
          </a>
          <a role="menuitem" className={item} href={`${base}.txt`} download={`${slug}.txt`}>
            Download plain text (.txt)
          </a>
          <a role="menuitem" className={item} href={`${base}/openapi.json`} target="_blank" rel="noreferrer">
            OpenAPI document (JSON)
          </a>
          <a role="menuitem" className={item} href="/llms.txt" target="_blank" rel="noreferrer">
            Site index (llms.txt)
          </a>
        </div>
      )}
    </div>
  );
}
