import { Link, NavLink, useNavigate } from "react-router";
import { useState, type FormEvent, type ReactNode } from "react";
import { site } from "@/lib/site";

const linkClass = ({ isActive }: { isActive: boolean }) =>
  `rounded-md px-3 py-1.5 text-sm font-medium transition-colors ${
    isActive ? "text-white" : "text-gray-600 hover:bg-gray-100 hover:text-gray-900"
  }`;

const activeStyle = ({ isActive }: { isActive: boolean }) => (isActive ? { background: site.accent } : undefined);

/** Top navigation plus page body. Docs pages pass `fullBleed` so Scalar owns the viewport. */
export default function Shell({ children, fullBleed = false }: { children: ReactNode; fullBleed?: boolean }) {
  const navigate = useNavigate();
  const [q, setQ] = useState("");
  function onSearch(e: FormEvent) {
    e.preventDefault();
    if (q.trim()) navigate(`/search?q=${encodeURIComponent(q.trim())}`);
  }
  return (
    <div className="flex min-h-screen flex-col bg-white text-gray-900" style={{ ["--accent" as string]: site.accent }}>
      <header className="sticky top-0 z-20 border-b border-gray-200 bg-white/95 backdrop-blur">
        <div className="mx-auto flex h-14 max-w-screen-2xl items-center gap-4 px-4">
          <Link to="/" className="flex items-center gap-2 text-base font-semibold tracking-tight">
            {site.logo ? (
              <img src={site.logo} alt={site.name} className="h-7 w-auto" />
            ) : (
              <>
                <span className="inline-block h-6 w-6 rounded" style={{ background: site.accent }} aria-hidden />
                {site.name}
              </>
            )}
          </Link>
          <nav className="ml-4 flex items-center gap-1">
            <NavLink to="/" end className={linkClass} style={activeStyle}>
              Collections
            </NavLink>
            <NavLink to="/admin" className={linkClass} style={activeStyle}>
              Admin
            </NavLink>
          </nav>
          <div className="ml-auto flex items-center gap-3">
            <form onSubmit={onSearch} role="search" className="hidden sm:block">
              <input
                type="search"
                value={q}
                onChange={(e) => setQ(e.target.value)}
                placeholder="Search endpoints…"
                aria-label="Search endpoints"
                className="w-56 rounded-md border border-gray-300 px-3 py-1.5 text-sm"
              />
            </form>
            <a href="/llms.txt" className="text-xs text-gray-500 hover:text-gray-900" title="Machine-readable index of all documentation">
              llms.txt
            </a>
          </div>
        </div>
      </header>
      <main className={fullBleed ? "flex-1" : "mx-auto w-full max-w-7xl flex-1 px-4 py-8"}>{children}</main>
    </div>
  );
}
