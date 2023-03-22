package sliding_window_limiter

import (
	"testing"
	"time"
)

func TestSlidingWindowLimiter(t *testing.T) {
	windowSize := time.Second
	limit := 2
	limiter := NewSlidingWindowLimiter(windowSize, limit)

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

/*
在这个测试用例中，我们首先创建了一个新的 SlidingWindowLimiter 实例，窗口大小为 1 秒，限制数量为 2。然后，我们调用 allow 方法来测试限流器的行为。

前两次调用 allow 方法应该返回 true，因为此时窗口内已经通过的请求数量还没有达到限制数量。
第三次调用 allow 方法应该返回 false，因为此时窗口内已经通过的请求数量已经达到了限制数量。

接下来，我们等待一个窗口大小的时间，让当前窗口过期。然后再次调用 allow 方法，此时应该返回 true。
*/
