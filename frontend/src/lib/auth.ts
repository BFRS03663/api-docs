import { ApiError } from "@/api/client";

const KEY = "apidocs.admin";

interface Session {
  token: string;
  expiresAt: string;
}

// Fallback when localStorage is unavailable (private mode, blocked storage).
let memory: Session | null = null;

function read(): Session | null {
  let s: Session | null = memory;
  try {
    const raw = localStorage.getItem(KEY);
    if (raw) s = JSON.parse(raw) as Session;
  } catch {
    // storage blocked; rely on memory
  }
  if (!s?.token) return null;
  if (new Date(s.expiresAt).getTime() <= Date.now()) {
    clearSession();
    return null;
  }
  return s;
}

export function getToken(): string | null {
  return read()?.token ?? null;
}

export function setSession(token: string, expiresAt: string): void {
  memory = { token, expiresAt };
  try {
    localStorage.setItem(KEY, JSON.stringify(memory));
  } catch {
    // keep the in-memory copy only
  }
}

export function clearSession(): void {
  memory = null;
  try {
    localStorage.removeItem(KEY);
  } catch {
    // ignore
  }
}

export function isAuthenticated(): boolean {
  return getToken() !== null;
}

/**
 * fetch with the admin bearer token. A 401 clears the session so route
 * guards send the user back to the login page.
 */
export async function authFetch(input: string, init: RequestInit = {}): Promise<Response> {
  const token = getToken();
  const headers = new Headers(init.headers);
  if (token) headers.set("Authorization", `Bearer ${token}`);
  const res = await fetch(input, { ...init, headers });
  if (res.status === 401) {
    clearSession();
    throw new ApiError(401, "session expired, please sign in again");
  }
  return res;
}
