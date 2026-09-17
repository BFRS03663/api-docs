import { lazy, Suspense } from "react";
import { Route, Routes } from "react-router";
import Home from "@/pages/Home";
import NotFound from "@/pages/NotFound";
import Search from "@/pages/Search";
import Shell from "@/components/Shell";
import AdminLogin from "@/pages/admin/AdminLogin";
import AdminLayout from "@/pages/admin/AdminLayout";
import AdminCollections from "@/pages/admin/AdminCollections";
import AdminImport from "@/pages/admin/AdminImport";

// Scalar is the largest dependency; only load it when a docs page is opened.
const Docs = lazy(() => import("@/pages/Docs"));
const AdminEditor = lazy(() => import("@/pages/admin/AdminEditor"));

function DocsFallback() {
  return (
    <Shell>
      <p className="text-sm text-gray-500">Loading documentation…</p>
    </Shell>
  );
}

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route
        path="/docs/:slug"
        element={
          <Suspense fallback={<DocsFallback />}>
            <Docs />
          </Suspense>
        }
      />
      <Route path="/search" element={<Search />} />
      <Route path="/admin/login" element={<AdminLogin />} />
      <Route path="/admin" element={<AdminLayout />}>
        <Route index element={<AdminCollections />} />
        <Route path="import" element={<AdminImport />} />
        <Route path="import/:slug" element={<AdminImport />} />
        <Route
          path="editor/:slug?"
          element={
            <Suspense fallback={<p className="text-sm text-gray-500">Loading editor…</p>}>
              <AdminEditor />
            </Suspense>
          }
        />
      </Route>
      <Route path="*" element={<NotFound />} />
    </Routes>
  );
}
