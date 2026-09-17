package httpapi

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// ipLimiter is a per-client-IP token bucket. It is deliberately small and
// in-memory: the portal runs as a single instance and the limits protect
// the login and proxy endpoints from brute force, not from volumetric DoS.
type ipLimiter struct {
	mu        sync.Mutex
	buckets   map[string]*entry
	limit     rate.Limit
	burst     int
	lastPurge time.Time
	now       func() time.Time
}

type entry struct {
	lim  *rate.Limiter
	seen time.Time
}

func newIPLimiter(perMinute int) *ipLimiter {
	return &ipLimiter{
		buckets: map[string]*entry{},
		limit:   rate.Every(time.Minute / time.Duration(perMinute)),
		burst:   perMinute,
		now:     time.Now,
	}
}

func (l *ipLimiter) allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := l.now()
	if now.Sub(l.lastPurge) > 10*time.Minute {
		for k, e := range l.buckets {
			if now.Sub(e.seen) > 10*time.Minute {
				delete(l.buckets, k)
			}
		}
		l.lastPurge = now
	}
	e, ok := l.buckets[key]
	if !ok {
		e = &entry{lim: rate.NewLimiter(l.limit, l.burst)}
		l.buckets[key] = e
	}
	e.seen = now
	return e.lim.AllowN(now, 1)
}

// middleware returns 429 once a client exceeds its bucket.
func (l *ipLimiter) middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !l.allow(c.ClientIP()) {
			c.Header("Retry-After", "60")
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			return
		}
		c.Next()
	}
}
