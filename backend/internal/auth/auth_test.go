package auth

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func newService(t *testing.T) *Service {
	t.Helper()
	hash, err := bcrypt.GenerateFromPassword([]byte("correct horse"), bcrypt.MinCost)
	if err != nil {
		t.Fatal(err)
	}
	return New("admin", string(hash), strings.Repeat("s", 32), time.Hour)
}

func TestLoginAndVerify(t *testing.T) {
	s := newService(t)
	token, exp, err := s.Login("admin", "correct horse")
	if err != nil {
		t.Fatalf("login: %v", err)
	}
	if time.Until(exp) < 59*time.Minute {
		t.Errorf("expiry too soon: %v", exp)
	}
	user, err := s.Verify(token)
	if err != nil || user != "admin" {
		t.Fatalf("verify = %q, %v", user, err)
	}

	for _, bad := range [][2]string{{"admin", "wrong"}, {"root", "correct horse"}, {"", ""}} {
		if _, _, err := s.Login(bad[0], bad[1]); !errors.Is(err, ErrInvalidCredentials) {
			t.Errorf("login(%q,%q) err = %v", bad[0], bad[1], err)
		}
	}
}

func TestVerifyRejectsTamperedExpiredAndForeignTokens(t *testing.T) {
	s := newService(t)
	token, _, _ := s.Login("admin", "correct horse")

	if _, err := s.Verify(token + "x"); err == nil {
		t.Error("tampered token accepted")
	}
	other := New("admin", string(s.hash), strings.Repeat("o", 32), time.Hour)
	if _, err := other.Verify(token); err == nil {
		t.Error("token signed with another secret accepted")
	}
	s.now = func() time.Time { return time.Now().Add(2 * time.Hour) }
	if _, err := s.Verify(token); err == nil {
		t.Error("expired token accepted")
	}
}

func TestMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	s := newService(t)
	token, _, _ := s.Login("admin", "correct horse")

	r := gin.New()
	r.GET("/secret", s.Middleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "hello %s", c.GetString(ContextKey))
	})

	cases := map[string]struct {
		header string
		want   int
	}{
		"missing":     {"", http.StatusUnauthorized},
		"not bearer":  {"Basic abc", http.StatusUnauthorized},
		"garbage":     {"Bearer nope", http.StatusUnauthorized},
		"valid token": {"Bearer " + token, http.StatusOK},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/secret", nil)
			if tc.header != "" {
				req.Header.Set("Authorization", tc.header)
			}
			r.ServeHTTP(rec, req)
			if rec.Code != tc.want {
				t.Fatalf("status = %d, want %d (%s)", rec.Code, tc.want, rec.Body)
			}
			if tc.want == http.StatusOK && rec.Body.String() != "hello admin" {
				t.Errorf("body = %s", rec.Body)
			}
		})
	}
}
