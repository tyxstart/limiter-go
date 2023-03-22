# 限流算法

本文档汇总了各种常用的限流算法，包括令牌桶、漏桶、滑动窗口、滑动日志和固定窗口限流算法。下面将分别介绍这些算法的使用示例、优缺点和应用方向。

## 令牌桶算法

### 使用示例

```go
package main

import (
	"fmt"
	"time"

	"github.com/yourusername/yourproject/token_bucket_limiter"
)

func main() {
	rate := time.Second
	capacity := 10
	tb := token_bucket_limiter.NewTokenBucketLimiter(rate, capacity)

	for i := 0; i < capacity+1; i++ {
		fmt.Println(tb.Allow())
	}
}
```
### 优缺点
令牌桶算法能够很好地处理突发流量，因为它允许请求在短时间内快速通过。但是，它也可能导致请求在短时间内过多地通过，从而对系统造成压力。

### 应用方向
令牌桶算法适用于需要处理突发流量的场景，例如Web服务器或API接口。

## 漏桶算法
### 使用示例
```go
package main

import (
	"fmt"
	"time"

	"github.com/yourusername/yourproject/leaky_bucket_limiter"
)

func main() {
	rate := time.Second
	capacity := 10
	lb := leaky_bucket_limiter.NewLeakyBucketLimiter(rate, capacity)

	for i := 0; i < capacity+1; i++ {
		fmt.Println(lb.Allow())
	}
}
```

### 优缺点
漏桶算法能够很好地控制请求速率，因为它强制请求以固定速率通过。但是，它不能很好地处理突发流量，因为新的请求将被拒绝直到漏桶中有足够的空间。

### 应用方向
漏桶算法适用于需要严格控制请求速率的场景，例如消息队列或数据库操作。

## 滑动窗口算法
### 使用示例
```go
package main

import (
	"fmt"
	"time"

	"github.com/yourusername/yourproject/sliding_window_limiter"
)

func main() {
	windowSize := time.Second * 5
	limit := 100
	sw := sliding_window_limiter.NewSlidingWindowLimiter(windowSize, limit)

	for i := 0; i < limit+1; i++ {
		fmt.Println(sw.Allow())
```

### 优缺点

滑动窗口算法能够很好地控制请求速率，因为它在一个固定时间窗口内限制了允许通过的请求数量。但是，它不能很好地处理突发流量，因为新的请求将被拒绝直到当前时间窗口结束。

### 应用方向

滑动窗口算法适用于需要在一个固定时间窗口内控制请求速率的场景，例如API接口或Web服务器。

## 滑动日志算法

### 使用示例

```go
package main

import (
	"fmt"
	"time"

	"github.com/yourusername/yourproject/sliding_log_limiter"
)

func main() {
	windowSize := time.Second * 5
	limit := 100
	sl := sliding_log_limiter.NewSlidingLogLimiter(windowSize, limit)

	for i := 0; i < limit+1; i++ {
		fmt.Println(sl.Allow())
	}
}
```
### 优缺点
滑动日志算法能够很好地控制请求速率，因为它记录了所有通过的请求并在一个固定时间窗口内限制了允许通过的请求数量。但是，由于需要记录所有通过的请求，该算法可能会占用较多的内存。

### 应用方向
滑动日志算法适用于需要在一个固定时间窗口内控制请求速率且对内存占用不敏感的场景，例如API接口或Web服务器。

## 固定窗口算法
### 使用示例
```go
package main

import (
	"fmt"
	"time"

	"github.com/yourusername/yourproject/fixed_window_limiter"
)

func main() {
	interval := time.Second * 5
	limit := 100
	fw := fixed_window_limiter.NewFixedWindowRateLimiter(limit, interval)

	for i := 0; i < limit+1; i++ {
		fmt.Println(fw.Allow())
	}
}
```

### 优缺点
固定窗口算法能够很好地控制请求速率，因为它在一个固定时间窗口内限制了允许通过的请求数量。但是，由于它使用了固定时间窗口而非滑动时间窗口，在某些情况下可能会导致流量不均衡。

### 应用方向
固定窗口算法适用于需要在一个固定时间段内控制请求速率且对流量均衡性要求不高的场景，例如API接口或Web服务器。