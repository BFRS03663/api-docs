import { useQuery } from "@tanstack/react-query";

export type SourceType = "openapi2" | "openapi3" | "postman";

export interface Collection {
  id: string;
  slug: string;
  name: string;
  description: string;
  version: string;
  source: { type: SourceType; filename: string; importedAt: string };
  servers: string[];
  operationCount: number;
  createdAt: string;
  updatedAt: string;
}

export interface Operation {
  id: string;
  collectionId: string;
  method: string;
  path: string;
  operationId: string;
  summary: string;
  description: string;
  tags: string[];
  deprecated: boolean;
  order: number;
}

export class ApiError extends Error {
  constructor(
    public status: number,
    message: string,
  ) {
    super(message);
  }
}

async function getJSON<T>(path: string): Promise<T> {
  const res = await fetch(path, { headers: { Accept: "application/json" } });
  if (!res.ok) {
    let message = res.statusText;
    try {
      const body = (await res.json()) as { error?: string };
      if (body.error) message = body.error;
    } catch {
      // non-JSON error body; keep statusText
    }
    throw new ApiError(res.status, message);
  }
  return (await res.json()) as T;
}

export function useCollections() {
  return useQuery({
    queryKey: ["collections"],
    queryFn: () => getJSON<{ collections: Collection[] }>("/api/v1/collections").then((b) => b.collections),
  });
}

export function useCollection(slug: string | undefined) {
  return useQuery({
    queryKey: ["collection", slug],
    queryFn: () => getJSON<Collection>(`/api/v1/collections/${slug}`),
    enabled: Boolean(slug),
    retry: (count, err) => !(err instanceof ApiError && err.status === 404) && count < 2,
  });
}

export function useOperations(slug: string | undefined) {
  return useQuery({
    queryKey: ["operations", slug],
    queryFn: () => getJSON<{ operations: Operation[] }>(`/api/v1/collections/${slug}/operations`).then((b) => b.operations),
    enabled: Boolean(slug),
  });
}
