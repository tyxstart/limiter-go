package token_bucket_limiter

import (
	"testing"
	"time"
)

func TestTokenBucketLimiter(t *testing.T) {
	rate := time.Second
	capacity := 10
	tb := NewTokenBucketLimiter(rate, capacity)

	for i := 0; i < capacity; i++ {
		if !tb.allow() {
			t.Errorf("expected allow to return true")
		}
	}

	if tb.allow() {
		t.Errorf("expected allow to return false")
	}

	time.Sleep(rate)

	if !tb.allow() {
		t.Errorf("expected allow to return true")
	}
}

/*
这段测试代码用于测试TokenBucketLimiter结构体的功能。它首先创建一个新的令牌桶限流器，然后在循环中调用allow()方法，
检查是否允许通过请求。如果在容量范围内，allow()方法应返回true。当超出容量时，allow()方法应返回false。

接下来，测试代码等待一段时间（与产生令牌速率相同），然后再次调用allow()方法。由于此时已经产生了新的令牌，因此allow()方法应返回true。

这段测试代码旨在验证令牌桶限流器的基本功能是否正常运行。
*/
