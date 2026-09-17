import type { ApiReferenceConfiguration } from "@scalar/api-reference-react";
import { site } from "./site";

/**
 * Scalar's in-page anchor for an operation: #tag/<first tag>/<METHOD>/<path>.
 * Untagged operations fall back to the top of the page.
 */
export function operationAnchor(op: { method: string; path: string; tags: string[] }): string {
  const tag = op.tags[0];
  if (!tag) return "";
  return `#tag/${tag}/${op.method.toUpperCase()}${op.path}`;
}

/** URL the backend serves the canonical OpenAPI document from. */
export function specUrl(slug: string): string {
  return `/docs/${encodeURIComponent(slug)}/openapi.json`;
}

/**
 * Styles layered on Scalar's default theme so the reference matches the
 * shell: brand accent, white chrome, and a dark example column like
 * Postman's published docs.
 */
export function themeCss(accent: string): string {
  return `
.light-mode, .dark-mode {
  --scalar-color-accent: ${accent};
  --scalar-background-accent: ${accent}1a;
  --scalar-border-radius: 6px;
}
.light-mode {
  --scalar-sidebar-background-1: #ffffff;
  --scalar-sidebar-color-active: ${accent};
  --scalar-sidebar-item-active-background: ${accent}14;
  --scalar-sidebar-search-background: #f6f6f6;
}
/* Right-hand example column: dark like Postman's published docs. */
.light-mode .section-columns > .section-column:last-child:not(:first-child) {
  --scalar-background-1: #303030;
  --scalar-background-2: #3a3a3a;
  --scalar-background-3: #474747;
  --scalar-color-1: #f5f5f5;
  --scalar-color-2: #d4d4d4;
  --scalar-color-3: #a8a8a8;
  --scalar-border-color: #4a4a4a;
  background: #303030;
  color: var(--scalar-color-1);
  border-radius: 10px;
  padding: 12px;
}
.light-mode .section-columns > .section-column:last-child:not(:first-child) .scalar-card {
  background: var(--scalar-background-2);
}
`;
}

/** Backend endpoint implementing Scalar's proxy contract (?scalar_url=). */
export const defaultProxyUrl = "/api/v1/proxy";

/**
 * Scalar configuration for one collection. Try-it requests go through the
 * backend proxy so APIs that block cross-origin calls still work; pass an
 * empty string to call target APIs directly from the browser.
 */
export function scalarConfig(slug: string, proxyUrl: string = defaultProxyUrl): Partial<ApiReferenceConfiguration> {
  return { ...baseConfig(proxyUrl), url: specUrl(slug) };
}

/**
 * Renders an unsaved document from the admin editor. Scalar parses the YAML or
 * JSON string itself, so the preview needs no server round trip. The classic
 * single-column layout fits the half-width preview pane.
 */
export function scalarPreviewConfig(content: string, proxyUrl: string = defaultProxyUrl): Partial<ApiReferenceConfiguration> {
  return { ...baseConfig(proxyUrl), content, layout: "classic", showSidebar: false };
}

function baseConfig(proxyUrl: string): Partial<ApiReferenceConfiguration> {
  return {
    proxyUrl: proxyUrl || undefined,
    layout: "modern",
    showSidebar: true,
    hideModels: false,
    hideDownloadButton: false,
    hideDarkModeToggle: true,
    forceDarkModeState: "light",
    searchHotKey: "k",
    theme: "default",
    customCss: themeCss(site.accent),
    // Accepted at runtime but missing from the published type: hides Scalar's
    // own "Developer Tools / Share / Deploy" bar so only our shell header shows.
    ...({ showToolbar: "never" } as Record<string, unknown>),
  };
}
