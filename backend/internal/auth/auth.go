// Package auth implements the single-admin login: bcrypt password check and
// short-lived HS256 JWTs, plus the Gin middleware that guards admin routes.
package auth

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// ErrInvalidCredentials is returned for any login failure; callers must not
// reveal which part was wrong.
var ErrInvalidCredentials = errors.New("invalid credentials")

const issuer = "apidocs"

// Service issues and verifies admin tokens.
type Service struct {
	username string
	hash     []byte
	secret   []byte
	ttl      time.Duration
	now      func() time.Time
}

// New builds a Service. hash is the bcrypt hash of the admin password.
func New(username, hash, secret string, ttl time.Duration) *Service {
	return &Service{username: username, hash: []byte(hash), secret: []byte(secret), ttl: ttl, now: time.Now}
}

// Login checks the credentials and returns a signed token with its expiry.
// The bcrypt comparison always runs so a wrong username costs the same time
// as a wrong password.
func (s *Service) Login(username, password string) (string, time.Time, error) {
	userOK := subtle.ConstantTimeCompare([]byte(username), []byte(s.username)) == 1
	passOK := bcrypt.CompareHashAndPassword(s.hash, []byte(password)) == nil
	if !userOK || !passOK {
		return "", time.Time{}, ErrInvalidCredentials
	}
	now := s.now()
	exp := now.Add(s.ttl)
	claims := jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   s.username,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(exp),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.secret)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("sign token: %w", err)
	}
	return token, exp, nil
}

// Verify parses a token and returns the admin username it belongs to.
func (s *Service) Verify(token string) (string, error) {
	claims := &jwt.RegisteredClaims{}
	parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return s.secret, nil
	}, jwt.WithIssuer(issuer), jwt.WithExpirationRequired(), jwt.WithTimeFunc(func() time.Time { return s.now() }))
	if err != nil || !parsed.Valid {
		return "", ErrInvalidCredentials
	}
	if subtle.ConstantTimeCompare([]byte(claims.Subject), []byte(s.username)) != 1 {
		return "", ErrInvalidCredentials
	}
	return claims.Subject, nil
}

// ContextKey is where the middleware stores the authenticated username.
const ContextKey = "admin"

// Middleware rejects requests without a valid bearer token.
func (s *Service) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		const prefix = "Bearer "
		if !strings.HasPrefix(header, prefix) {
			unauthorized(c)
			return
		}
		user, err := s.Verify(strings.TrimSpace(header[len(prefix):]))
		if err != nil {
			unauthorized(c)
			return
		}
		c.Set(ContextKey, user)
		c.Next()
	}
}

func unauthorized(c *gin.Context) {
	c.Header("WWW-Authenticate", `Bearer realm="admin"`)
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
}
