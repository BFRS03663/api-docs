// Package config loads application settings from environment variables and
// fails fast when a required value is missing.
package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config holds every runtime setting the server needs.
type Config struct {
	Port              string
	MongoURI          string
	MongoDB           string
	AdminUsername     string
	AdminPasswordHash string
	JWTSecret         string
	CORSOrigins       []string
	ProxyAllowedHosts []string
	ProxyAllowPrivate bool
	ShutdownTimeout   time.Duration
	GinMode           string
	SiteName          string
}

// requiredVars lists every variable that must be present at startup.
var requiredVars = []string{
	"PORT",
	"MONGO_URI",
	"MONGO_DB",
	"ADMIN_USERNAME",
	"ADMIN_PASSWORD_HASH",
	"JWT_SECRET",
	"CORS_ORIGINS",
}

// Load reads the environment and returns a validated Config. The returned
// error names every missing variable so operators can fix them in one pass.
func Load() (*Config, error) {
	var missing []string
	for _, name := range requiredVars {
		if strings.TrimSpace(os.Getenv(name)) == "" {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	allowPrivate, err := parseBool(os.Getenv("PROXY_ALLOW_PRIVATE"), false)
	if err != nil {
		return nil, fmt.Errorf("PROXY_ALLOW_PRIVATE: %w", err)
	}

	cfg := &Config{
		Port:              os.Getenv("PORT"),
		MongoURI:          os.Getenv("MONGO_URI"),
		MongoDB:           os.Getenv("MONGO_DB"),
		AdminUsername:     os.Getenv("ADMIN_USERNAME"),
		AdminPasswordHash: os.Getenv("ADMIN_PASSWORD_HASH"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		CORSOrigins:       splitList(os.Getenv("CORS_ORIGINS")),
		ProxyAllowedHosts: splitList(os.Getenv("PROXY_ALLOWED_HOSTS")),
		ProxyAllowPrivate: allowPrivate,
		ShutdownTimeout:   10 * time.Second,
		GinMode:           envOr("GIN_MODE", "release"),
		SiteName:          envOr("SITE_NAME", "API Docs"),
	}

	if len(cfg.JWTSecret) < 32 {
		return nil, errors.New("JWT_SECRET must be at least 32 characters")
	}
	if !strings.HasPrefix(cfg.AdminPasswordHash, "$2") {
		return nil, errors.New("ADMIN_PASSWORD_HASH must be a bcrypt hash (run `make hash-password`)")
	}
	return cfg, nil
}

func splitList(raw string) []string {
	var out []string
	for part := range strings.SplitSeq(raw, ",") {
		if p := strings.TrimSpace(part); p != "" {
			out = append(out, p)
		}
	}
	return out
}

func parseBool(raw string, def bool) (bool, error) {
	if strings.TrimSpace(raw) == "" {
		return def, nil
	}
	return strconv.ParseBool(raw)
}

func envOr(name, def string) string {
	if v := strings.TrimSpace(os.Getenv(name)); v != "" {
		return v
	}
	return def
}
