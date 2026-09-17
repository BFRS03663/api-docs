import { useState, type FormEvent } from "react";
import { Navigate, useLocation, useNavigate } from "react-router";
import Shell from "@/components/Shell";
import { login } from "@/api/admin";
import { isAuthenticated } from "@/lib/auth";
import { site } from "@/lib/site";

export default function AdminLogin() {
  const navigate = useNavigate();
  const location = useLocation();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [busy, setBusy] = useState(false);

  if (isAuthenticated()) {
    return <Navigate to="/admin" replace />;
  }
  const from = (location.state as { from?: string } | null)?.from ?? "/admin";

  async function onSubmit(e: FormEvent) {
    e.preventDefault();
    setBusy(true);
    setError(null);
    try {
      await login(username, password);
      navigate(from, { replace: true });
    } catch (err) {
      setError(err instanceof Error ? err.message : "login failed");
    } finally {
      setBusy(false);
    }
  }

  return (
    <Shell>
      <div className="mx-auto mt-16 max-w-sm">
        <h1 className="text-2xl font-semibold">Admin sign in</h1>
        <p className="mt-1 text-sm text-gray-600">Manage the collections published on {site.name}.</p>
        <form onSubmit={onSubmit} className="mt-6 space-y-4" aria-label="sign in">
          <label className="block text-sm">
            <span className="font-medium">Username</span>
            <input
              className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2"
              autoComplete="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </label>
          <label className="block text-sm">
            <span className="font-medium">Password</span>
            <input
              type="password"
              className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2"
              autoComplete="current-password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </label>
          {error && (
            <p role="alert" className="rounded-md border border-red-200 bg-red-50 p-2 text-sm text-red-700">
              {error}
            </p>
          )}
          <button
            type="submit"
            disabled={busy}
            className="w-full rounded-md px-4 py-2 text-sm font-medium text-white disabled:opacity-60"
            style={{ background: site.accent }}
          >
            {busy ? "Signing in…" : "Sign in"}
          </button>
        </form>
      </div>
    </Shell>
  );
}
