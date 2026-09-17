/**
 * Site branding. Override per deployment with Vite env vars:
 *   VITE_SITE_NAME   header title
 *   VITE_SITE_LOGO   image URL or /public path; empty shows a plain mark
 *   VITE_ACCENT      accent colour used by the shell and Scalar
 */
export const site = {
  name: import.meta.env.VITE_SITE_NAME ?? "API Docs",
  logo: import.meta.env.VITE_SITE_LOGO ?? "",
  accent: import.meta.env.VITE_ACCENT ?? "#EF5B25",
};
