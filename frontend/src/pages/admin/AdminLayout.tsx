import { Navigate, NavLink, Outlet, useLocation, useNavigate } from "react-router";
import Shell from "@/components/Shell";
import { clearSession, isAuthenticated } from "@/lib/auth";
import { site } from "@/lib/site";

const tab = ({ isActive }: { isActive: boolean }) =>
  `rounded-md px-3 py-1.5 text-sm font-medium ${isActive ? "bg-gray-900 text-white" : "text-gray-600 hover:bg-gray-100"}`;

/** Guards every /admin route: unauthenticated visitors go to the login page. */
export default function AdminLayout() {
  const location = useLocation();
  const navigate = useNavigate();
  if (!isAuthenticated()) {
    return <Navigate to="/admin/login" replace state={{ from: location.pathname }} />;
  }
  function signOut() {
    clearSession();
    void fetch("/api/v1/auth/logout", { method: "POST" }).catch(() => undefined);
    navigate("/admin/login", { replace: true });
  }
  // The editor owns the viewport width like the docs page does.
  const fullBleed = location.pathname.startsWith("/admin/editor");
  return (
    <Shell fullBleed={fullBleed}>
      <div className={`flex flex-wrap items-center gap-2 ${fullBleed ? "mb-2 px-4 pt-4" : "mb-6"}`}>
        <h1 className="mr-4 text-2xl font-semibold">Admin</h1>
        <NavLink to="/admin" end className={tab}>
          Collections
        </NavLink>
        <NavLink to="/admin/import" className={tab}>
          Import
        </NavLink>
        <NavLink to="/admin/editor" className={tab}>
          Write spec
        </NavLink>
        <button onClick={signOut} className="ml-auto text-sm text-gray-600 hover:underline" style={{ color: site.accent }}>
          Sign out
        </button>
      </div>
      <Outlet />
    </Shell>
  );
}
