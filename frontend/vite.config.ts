/// <reference types="vitest/config" />
import { defineConfig } from "vite";
import { fileURLToPath, URL } from "node:url";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";
import repoDocs from "./vite-plugin-repo-docs";

const backend = process.env.VITE_BACKEND_URL ?? "http://localhost:8080";
const frontendDir = fileURLToPath(new URL(".", import.meta.url));

export default defineConfig({
  plugins: [react(), tailwindcss(), repoDocs(frontendDir)],
  resolve: { alias: { "@": fileURLToPath(new URL("./src", import.meta.url)) } },
  server: {
    port: 5173,
    proxy: {
      "/api": backend,
      "/docs": { target: backend, bypass: (req) => (req.headers.accept?.includes("text/html") ? "/index.html" : undefined) },
      "/llms.txt": backend,
      "/llms-full.txt": backend,
      "/healthz": backend,
    },
  },
  build: { outDir: "dist", emptyOutDir: true },
  test: {
    environment: "jsdom",
    globals: true,
    setupFiles: ["./src/test/setup.ts"],
  },
});
