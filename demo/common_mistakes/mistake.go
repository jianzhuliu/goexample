package main

import "sync"

import (
	"fmt"
	"time"
)

type Counter struct {
	sync.Mutex
	n int64
}

// 此方法实现是没问题的。
func (c *Counter) Increase(d int64) (r int64) {
	c.Lock()
	c.n += d
	r = c.n
	c.Unlock()
	return
}

// 此方法的实现是有问题的。当它被调用时，
// 一个Counter属主值将被复制。
func (c Counter) Value() (r int64) {
	c.Lock()
	r = c.n
	c.Unlock()
	return
}

func request1() int {
	c := make(chan int) //改成带有缓存的通道， c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			c <- i // 4个协程将永久阻塞在这里
		}()
	}
	return <-c
}

func request2() int {
	c := make(chan int) //改成带有缓存的通道， c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			select {
			case c <- i:
			default:
			}
		}()
	}
	return <-c // 有可能永久阻塞在此
}

// 如果某两个连续的消息的间隔大于一分钟，此函数将返回。
//如果longRunning1函数被调用并且在一分钟内有一百万条消息到达，
//那么在某个特定的很小时间段（大概若干秒）内将存在一百万个活跃的Timer值，即使其中只有一个是真正有用的
func longRunning1(messages <-chan string) {
	for {
		select {
		case <-time.After(time.Minute):
			return
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}

//调整后，避免太多的Timer值被创建，我们应该只使用（并复用）一个Timer值
func longRunning2(messages <-chan string) {
	timer := time.NewTimer(time.Minute)
	defer timer.Stop()

	for {
		select {
		case <-timer.C: // 过期了
			return
		case msg := <-messages:
			fmt.Println(msg)

			// 此if代码块很重要。
			// 用来舍弃一个可能在执行第二个分支代码块的时候发送过来的超时通知
			if !timer.Stop() {
				<-timer.C
			}
		}

		// 必须重置以复用。
		timer.Reset(time.Minute)
	}
}
