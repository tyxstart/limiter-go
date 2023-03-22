package sliding_log_limiter

import (
	"testing"
	"time"
)

func TestSlidingLogLimiter(t *testing.T) {
	windowSize := time.Second
	limit := 2
	limiter := NewSlidingLogLimiter(windowSize, limit)

	if !limiter.allow() {
		t.Error("expected true but got false")
	}

	if !limiter.allow() {
		t.Error("expected true but got false")
	}

	if limiter.allow() {
		t.Error("expected false but got true")
	}

	time.Sleep(windowSize)

	if !limiter.allow() {
		t.Error("expected true but got false")
	}
}
