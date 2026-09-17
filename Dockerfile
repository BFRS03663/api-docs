# syntax=docker/dockerfile:1

# ---- frontend bundle ------------------------------------------------------
FROM node:22-alpine AS ui
WORKDIR /ui
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci --no-audit --no-fund
COPY frontend/ ./
# Branding is baked in at build time (Vite env); pass with --build-arg.
ARG VITE_SITE_NAME="API Docs"
ARG VITE_SITE_LOGO=""
ARG VITE_ACCENT="#EF5B25"
ENV VITE_SITE_NAME=$VITE_SITE_NAME VITE_SITE_LOGO=$VITE_SITE_LOGO VITE_ACCENT=$VITE_ACCENT
RUN npm run build

# ---- go binary with the bundle embedded -----------------------------------
FROM golang:1.26 AS build
WORKDIR /src
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ ./
COPY --from=ui /ui/dist ./internal/ui/dist
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/server ./cmd/server

# ---- runtime ----------------------------------------------------------------
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/server /server
ENV PORT=8080 GIN_MODE=release
EXPOSE 8080
USER nonroot
ENTRYPOINT ["/server"]
