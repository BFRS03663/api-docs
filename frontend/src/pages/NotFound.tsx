import { Link } from "react-router";
import Shell from "@/components/Shell";

export default function NotFound() {
  return (
    <Shell>
      <div className="py-20 text-center">
        <p className="font-mono text-sm text-gray-500">404</p>
        <h1 className="mt-2 text-2xl font-semibold">Page not found</h1>
        <p className="mt-2 text-gray-600">The collection or page you asked for does not exist.</p>
        <Link to="/" className="mt-6 inline-block rounded-md bg-gray-900 px-4 py-2 text-sm font-medium text-white">
          Back to collections
        </Link>
      </div>
    </Shell>
  );
}
