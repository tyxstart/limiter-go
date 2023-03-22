# 滑动日志限流算法

## 算法介绍

滑动日志限流算法是一种常用的限流算法。它通过记录请求日志并控制在一个固定时间窗口内允许通过的请求数量来控制请求速率。如果在当前时间窗口内已经达到了限制数量，则新的请求将被拒绝。

## 结构

在Go中，我们可以使用结构体来实现滑动日志限流器。下面是一个简单的例子：

```go
type SlidingLogLimiter struct {
	windowSize time.Duration // 窗口大小
	limit      int           // 限制数量
	logs       []int64       // 请求日志
}
```
## 方法及参数
我们可以定义一个NewSlidingLogLimiter函数来创建一个新的滑动日志限流器。该函数接受两个参数：windowSize和limit，分别表示窗口大小和限制数量。

此外，我们还需要定义一个allow()方法来判断是否允许通过请求。该方法不需要任何参数。

## 实现思路
在创建新的滑动日志限流器时，我们需要初始化结构体中的各个字段。然后，在调用allow()方法时，我们首先计算当前时间窗口的截止时间，并删除所有过期的请求日志。

接下来，判断剩余请求数量是否超过了限制数量。如果超过，则返回false；否则添加新的请求日志并返回true。

## 测试实现思路
在这个测试用例中，我们首先创建了一个新的 SlidingLogLimiter 实例，窗口大小为 1 秒，限制数量为 2。然后，我们调用 allow 方法来测试限流器的行为。

前两次调用 allow 方法应该返回 true，因为此时窗口内已经通过的请求数量还没有达到限制数量。
第三次调用 allow 方法应该返回 false，因为此时窗口内已经通过的请求数量已经达到了限制数量。

接下来，我们等待一个窗口大小的时间，让当前窗口过期。然后再次调用 allow 方法，此时应该返回 true。