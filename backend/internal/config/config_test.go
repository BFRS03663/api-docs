package config

import (
	"strings"
	"testing"
)

func setValid(t *testing.T) {
	t.Helper()
	t.Setenv("PORT", "8080")
	t.Setenv("MONGO_URI", "mongodb://localhost:27017")
	t.Setenv("MONGO_DB", "apidocs")
	t.Setenv("ADMIN_USERNAME", "admin")
	t.Setenv("ADMIN_PASSWORD_HASH", "$2a$10$abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRS")
	t.Setenv("JWT_SECRET", strings.Repeat("s", 32))
	t.Setenv("CORS_ORIGINS", "http://localhost:5173, http://localhost:3000")
	t.Setenv("PROXY_ALLOW_PRIVATE", "")
	t.Setenv("PROXY_ALLOWED_HOSTS", "")
}

func TestLoadValid(t *testing.T) {
	setValid(t)
	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cfg.CORSOrigins) != 2 || cfg.CORSOrigins[1] != "http://localhost:3000" {
		t.Fatalf("CORS origins not split/trimmed: %v", cfg.CORSOrigins)
	}
	if cfg.ProxyAllowPrivate {
		t.Fatal("ProxyAllowPrivate should default to false")
	}
	if cfg.GinMode != "release" {
		t.Fatalf("GinMode default = %q, want release", cfg.GinMode)
	}
}

func TestLoadMissingNamesEveryVar(t *testing.T) {
	setValid(t)
	t.Setenv("JWT_SECRET", "")
	t.Setenv("MONGO_URI", "")
	_, err := Load()
	if err == nil {
		t.Fatal("expected error")
	}
	for _, want := range []string{"JWT_SECRET", "MONGO_URI"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not name %s", err, want)
		}
	}
}

func TestLoadRejectsShortSecretAndBadHash(t *testing.T) {
	setValid(t)
	t.Setenv("JWT_SECRET", "short")
	if _, err := Load(); err == nil || !strings.Contains(err.Error(), "JWT_SECRET") {
		t.Fatalf("short secret: got %v", err)
	}

	setValid(t)
	t.Setenv("ADMIN_PASSWORD_HASH", "plaintext")
	if _, err := Load(); err == nil || !strings.Contains(err.Error(), "bcrypt") {
		t.Fatalf("bad hash: got %v", err)
	}
}
