import { useMutation, useQueryClient } from "@tanstack/react-query";
import { ApiError, type Collection } from "./client";
import { authFetch, setSession } from "@/lib/auth";

async function parseError(res: Response): Promise<ApiError> {
  let message = res.statusText || `request failed (${res.status})`;
  try {
    const body = (await res.json()) as { error?: string };
    if (body.error) message = body.error;
  } catch {
    // non-JSON body
  }
  return new ApiError(res.status, message);
}

export async function login(username: string, password: string): Promise<void> {
  const res = await fetch("/api/v1/auth/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password }),
  });
  if (!res.ok) throw await parseError(res);
  const body = (await res.json()) as { token: string; expiresAt: string };
  setSession(body.token, body.expiresAt);
}

export interface ImportResult {
  collection: Collection;
  created: boolean;
}

export type ImportSource = { kind: "file"; file: File } | { kind: "url"; url: string };

/**
 * Creates a collection (POST) or re-imports an existing one (PUT /:slug).
 * File uploads go as multipart; URL imports as JSON.
 */
export async function importCollection(source: ImportSource, slug: string, name: string, existing: boolean): Promise<ImportResult> {
  const path = existing ? `/api/v1/admin/collections/${encodeURIComponent(slug)}` : "/api/v1/admin/collections";
  const method = existing ? "PUT" : "POST";
  let res: Response;
  if (source.kind === "file") {
    const form = new FormData();
    form.append("file", source.file, source.file.name);
    form.append("slug", slug);
    form.append("name", name);
    res = await authFetch(path, { method, body: form });
  } else {
    res = await authFetch(path, {
      method,
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ url: source.url, slug, name }),
    });
  }
  if (!res.ok) throw await parseError(res);
  return (await res.json()) as ImportResult;
}

export async function deleteCollection(slug: string): Promise<void> {
  const res = await authFetch(`/api/v1/admin/collections/${encodeURIComponent(slug)}`, { method: "DELETE" });
  if (!res.ok && res.status !== 404) throw await parseError(res);
}

export function useDeleteCollection() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: deleteCollection,
    onSuccess: () => qc.invalidateQueries({ queryKey: ["collections"] }),
  });
}

export function useImportCollection() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (args: { source: ImportSource; slug: string; name: string; existing: boolean }) =>
      importCollection(args.source, args.slug, args.name, args.existing),
    onSuccess: () => qc.invalidateQueries({ queryKey: ["collections"] }),
  });
}
