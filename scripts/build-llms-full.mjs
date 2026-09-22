#!/usr/bin/env node
// Builds llms-full.txt from llms.txt and the reference files it links to.
//
// llms.txt is the index (llmstxt.org). llms-full.txt is the same index
// followed by the complete contents of every docs/shiprocket-api/*.md file,
// in the order the "Reference files" section lists them, so a language model
// can read the whole Shiprocket API reference from one URL.
//
// Links into the reference files are rewritten to in-document anchors so the
// file is self-contained. Anchors follow the GitHub convention (slug, then
// slug-1, slug-2 for repeated headings), recomputed for the concatenated
// document because repeats across files shift the numbering. Everything else
// is copied verbatim.
//
// Usage: node scripts/build-llms-full.mjs [--check]
//   --check  exit 1 if llms-full.txt is not up to date instead of writing it.

import { readFileSync, writeFileSync } from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const repoRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
const indexPath = path.join(repoRoot, "llms.txt");
const outPath = path.join(repoRoot, "llms-full.txt");
const docsDir = "docs/shiprocket-api";

/** GitHub-style heading slug, before de-duplication. */
function slug(heading) {
  return heading
    .trim()
    .toLowerCase()
    .replace(/[^\p{L}\p{N}\s-]/gu, "")
    .replace(/\s/g, "-");
}

/** Heading texts in order, ignoring lines inside fenced code blocks. */
function headings(markdown) {
  const out = [];
  let inFence = false;
  for (const line of markdown.split("\n")) {
    if (/^```/.test(line)) inFence = !inFence;
    else if (!inFence && /^#{1,6}\s/.test(line)) out.push(line.replace(/^#+\s+/, ""));
  }
  return out;
}

/** Anchors for a heading sequence: slug, then slug-1, slug-2 ... for repeats. */
function anchorsFor(headingTexts) {
  const seen = new Map();
  return headingTexts.map((h) => {
    const base = slug(h);
    const n = seen.get(base) ?? 0;
    seen.set(base, n + 1);
    return n === 0 ? base : `${base}-${n}`;
  });
}

const index = readFileSync(indexPath, "utf8").replace(/\s+$/, "");

// Reference files in the order the index lists them under "## Reference files".
const refSection = index.split(/^## Reference files$/m)[1]?.split(/^## /m)[0] ?? "";
const files = [...refSection.matchAll(new RegExp(`\\]\\((${docsDir}/[a-z-]+\\.md)\\)`, "g"))].map((m) => m[1]);
if (files.length === 0) throw new Error(`no ${docsDir} links found under "## Reference files" in llms.txt`);

const docs = files.map((rel) => {
  const text = readFileSync(path.join(repoRoot, rel), "utf8").replace(/\s+$/, "");
  const texts = headings(text);
  if (texts.length === 0) throw new Error(`${rel} has no heading`);
  return { rel, base: path.basename(rel), text, headingTexts: texts, localAnchors: anchorsFor(texts) };
});

// Global anchors: the index's own headings come first, then each file's, so
// numbering matches what a Markdown renderer produces for the whole document.
const globalAnchors = anchorsFor([...headings(index), ...docs.flatMap((d) => d.headingTexts)]);
const target = new Map(); // "<file>" and "<file>#<local anchor>" -> global anchor
let offset = headings(index).length;
for (const doc of docs) {
  for (const key of [doc.rel, doc.base]) {
    target.set(key, globalAnchors[offset]);
    doc.localAnchors.forEach((a, i) => target.set(`${key}#${a}`, globalAnchors[offset + i]));
  }
  offset += doc.headingTexts.length;
}

/** `x.md#a` / `docs/shiprocket-api/x.md#a` -> `#<global anchor>`; `x.md` -> its title. */
function rewriteLinks(markdown, source) {
  return markdown.replace(/\]\(([a-z/-]+\.md)(#[^)]+)?\)/g, (_whole, file, anchor = "") => {
    const global = target.get(file + anchor);
    if (!global) throw new Error(`${source}: link to ${file}${anchor} does not resolve to a heading`);
    return `](#${global})`;
  });
}

const intro =
  "This file is the complete reference in one document: the index below, followed by the full contents of " +
  `every reference file (${docs.length} files, in the order listed under "Reference files"). ` +
  "The companion llms.txt is the index alone, linking to the same files individually.";

let body = rewriteLinks(index, "llms.txt");
// Add the explanatory paragraph right after the index's own "This index links to..." paragraph.
const marker = /^(This index links to [^\n]*)$/m;
if (!marker.test(body)) throw new Error('llms.txt no longer has the "This index links to" paragraph');
body = body.replace(marker, `$1\n\n${intro}`);

for (const doc of docs) {
  body += `\n\n---\n\n${rewriteLinks(doc.text, doc.rel)}`;
}
body += "\n";

// Verification against the assembled document itself.
const known = new Set(anchorsFor(headings(body)));
const anchors = [...body.matchAll(/\]\(#([^)]+)\)/g)].map((m) => m[1]);
const missing = anchors.filter((a) => !known.has(a));
if (missing.length) throw new Error(`unresolved anchors: ${[...new Set(missing)].join(", ")}`);
const endpoints = (index.match(/^- \[(GET|POST|PUT|PATCH|DELETE) /gm) ?? []).length;

if (process.argv.includes("--check")) {
  let current = "";
  try {
    current = readFileSync(outPath, "utf8");
  } catch {}
  if (current !== body) {
    console.error("llms-full.txt is out of date; run: node scripts/build-llms-full.mjs");
    process.exit(1);
  }
  console.log("llms-full.txt is up to date");
} else {
  writeFileSync(outPath, body);
  console.log(
    `wrote ${path.relative(repoRoot, outPath)}: ${body.length} bytes, ${docs.length} reference files, ` +
      `${endpoints} indexed endpoints, ${anchors.length} anchors verified`,
  );
}
