package fixed_window_limiter

import (
	"testing"
	"time"
)

func TestFixedWindowLimiter(t *testing.T) {
	limit := 2
	interval := time.Second
	limiter := NewFixedWindowRateLimiter(limit, interval)

	if !limiter.Allow() {
		t.Error("expected true but got false")
	}

	if !limiter.Allow() {
		t.Error("expected true but got false")
	}

	if limiter.Allow() {
		t.Error("expected false but got true")
	}

	time.Sleep(interval)

	if !limiter.Allow() {
		t.Error("expected true but got false")
	}
}
