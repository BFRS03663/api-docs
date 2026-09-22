import { createReadStream, cpSync, existsSync, statSync } from "node:fs";
import path from "node:path";
import type { Plugin } from "vite";

/**
 * Serves the hand-written reference committed at the repository root
 * (`llms.txt`, `llms-full.txt` and `docs/shiprocket-api/*.md`) under one URL
 * prefix, so the relative links inside llms.txt keep resolving:
 *
 *   /reference/llms.txt
 *   /reference/llms-full.txt
 *   /reference/docs/shiprocket-api/<file>.md
 *
 * Dev: a middleware streams the files straight from the repo. Build: the
 * same files are copied into dist/<prefix>/ so the Go binary embeds and
 * serves them as static assets. `/llms.txt` itself stays with the backend,
 * which generates it from the imported collections.
 *
 * Source directory: REPO_DOCS_DIR, defaulting to the parent of frontend/.
 */
export const mountPrefix = "/reference";

/** Entries (relative to the repo root) that are exposed. Nothing else is. */
export const exposedEntries = ["llms.txt", "llms-full.txt", "docs/shiprocket-api"];

const contentTypes: Record<string, string> = {
  ".md": "text/markdown; charset=utf-8",
  ".txt": "text/plain; charset=utf-8",
};

function sourceRoot(frontendDir: string): string {
  return path.resolve(process.env.REPO_DOCS_DIR ?? path.join(frontendDir, ".."));
}

/** Path inside the repo for a request URL, or null when it is not exposed. */
export function resolveExposed(root: string, url: string): string | null {
  const [pathname] = url.split("?");
  if (!pathname.startsWith(mountPrefix + "/")) return null;
  let rel: string;
  try {
    rel = decodeURIComponent(pathname.slice(mountPrefix.length + 1));
  } catch {
    return null;
  }
  // Normalise first so ".." segments cannot smuggle a path past the allow-list.
  const file = path.resolve(root, rel);
  const normalized = path.relative(root, file).split(path.sep).join("/");
  if (normalized.startsWith("..") || path.isAbsolute(normalized)) return null;
  const allowed = exposedEntries.some((e) => normalized === e || normalized.startsWith(e + "/"));
  if (!allowed) return null;
  if (!existsSync(file) || !statSync(file).isFile()) return null;
  return file;
}

export default function repoDocs(frontendDir: string): Plugin {
  const root = sourceRoot(frontendDir);
  let outDir = "dist";
  let isBuild = false;
  return {
    name: "repo-docs",
    configResolved(config) {
      outDir = config.build.outDir;
      // Vitest also runs closeBundle, against a throwaway outDir; only copy on a real build.
      isBuild = config.command === "build" && config.mode !== "test";
    },
    configureServer(server) {
      for (const entry of exposedEntries) {
        if (!existsSync(path.join(root, entry))) {
          server.config.logger.warn(
            `repo-docs: ${path.join(root, entry)} not found; ${mountPrefix}/${entry} will 404 (set REPO_DOCS_DIR)`,
          );
        }
      }
      server.middlewares.use((req, res, next) => {
        const file = resolveExposed(root, req.url ?? "");
        if (!file) return next();
        res.setHeader("Content-Type", contentTypes[path.extname(file)] ?? "application/octet-stream");
        res.setHeader("Cache-Control", "no-cache");
        createReadStream(file).pipe(res);
      });
    },
    closeBundle() {
      if (!isBuild) return;
      const target = path.join(path.resolve(frontendDir, outDir), mountPrefix.slice(1));
      for (const entry of exposedEntries) {
        const from = path.join(root, entry);
        if (!existsSync(from)) {
          this.warn(`repo-docs: ${from} not found; ${mountPrefix}/${entry} will not be in the bundle`);
          continue;
        }
        cpSync(from, path.join(target, entry), { recursive: true });
      }
    },
  };
}
