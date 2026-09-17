package httpapi

import (
	"testing"
	"time"
)

func TestIPLimiterBucketsPerKeyAndRefills(t *testing.T) {
	l := newIPLimiter(2)
	clock := time.Now()
	l.now = func() time.Time { return clock }

	if !l.allow("a") || !l.allow("a") {
		t.Fatal("first two requests should pass")
	}
	if l.allow("a") {
		t.Fatal("third request within the minute should be limited")
	}
	if !l.allow("b") {
		t.Fatal("other client must have its own bucket")
	}
	clock = clock.Add(31 * time.Second)
	if !l.allow("a") {
		t.Fatal("bucket should refill one token after half a minute")
	}
}
