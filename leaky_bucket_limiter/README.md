# 漏桶限流算法

## 算法介绍

漏桶限流算法是一种常用的限流算法。它通过控制水流出漏桶的速率来控制请求速率。每个请求都相当于向漏桶中加入一滴水。如果漏桶已满，则新的请求将被拒绝。

## 结构

在Go中，我们可以使用结构体来实现漏桶限流器。下面是一个简单的例子：

```go
type LeakyBucketLimiter struct {
	mu        sync.Mutex    // 互斥锁
	rate      time.Duration // 漏水速率
	capacity  int           // 桶容量
	remaining int           // 剩余水量
	last      time.Time     // 上次漏水时间
}
```
## 方法及参数
我们可以定义一个NewLeakyBucketLimiter函数来创建一个新的漏桶限流器。该函数接受两个参数：rate和capacity，分别表示漏水速率和桶容量。

此外，我们还需要定义一个allow()方法来判断是否允许通过请求。该方法不需要任何参数。

## 实现思路
在创建新的漏桶限流器时，我们需要初始化结构体中的各个字段。然后，在调用allow()方法时，我们首先计算自上次调用以来已经漏出了多少水，并更新剩余水量。如果剩余数量小于0，则将其设置为0。

接下来，判断剩余数量是否超过了容量。如果超过，则返回false；否则增加剩余数量并返回true。

## 测试实现思路
在这个测试用例中，我们首先创建了一个新的 LeakyBucketLimiter 实例，漏水速率为每秒 5 次，桶容量为 5。然后，我们调用 allow 方法来测试限流器的行为。

前五次调用 allow 方法应该返回 true，因为此时桶内剩余水量还没有达到桶容量。
第六次调用 allow 方法应该返回 false，因为此时桶内剩余水量已经达到了桶容量。

接下来，我们等待一个漏水速率的时间，让桶内水量减少一个单位。然后再次调用 allow 方法，此时应该返回 true。