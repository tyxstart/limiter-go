package leaky_bucket_limiter

import (
	"testing"
	"time"
)

func TestLeakyBucketLimiter(t *testing.T) {
	rate := time.Second / 5
	capacity := 5
	limiter := NewLeakyBucketLimiter(rate, capacity)

	for i := 0; i < capacity; i++ {
		if !limiter.allow() {
			t.Errorf("expected true but got false at iteration %d", i)
		}
	}

	if limiter.allow() {
		t.Error("expected false but got true")
	}

	time.Sleep(rate)

	if !limiter.allow() {
		t.Error("expected true but got false")
	}
}
