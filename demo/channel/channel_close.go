/*
1、关闭一个nil通道或者一个已经关闭的通道将产生一个恐慌
2、向一个已关闭的通道发送数据也将导致一个恐慌
3、向一个nil通道发送数据或者从一个nil通道接收数据将使当前协程永久阻塞

*/

package main

import (
	"fmt"
	"sync"
)

type T int

//粗鲁地关闭
func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// 一个函数的返回结果可以在defer调用中修改。
			justClosed = false
		}
	}()

	// 假设ch != nil。
	close(ch)   // 如果ch已关闭，则产生一个恐慌。
	return true // <=> justClosed = true; return
}

//粗鲁地发送
func SafeSend(ch chan T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()

	ch <- value  // 如果ch已关闭，则产生一个恐慌。
	return false // <=> closed = false; return
}

////////////////////////////begin sync.Once
type MyChannel struct {
	C    chan T
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

////////////////////////////end sync.Once

////////////////////////////begin sync.Mutex
type MyChannelSM struct {
	C      chan T
	closed bool
	mutex  sync.Mutex
}

func NewMyChannelSM() *MyChannelSM {
	return &MyChannelSM{C: make(chan T)}
}

func (mc *MyChannelSM) SafeClose() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
}

func (mc *MyChannelSM) IsClosed() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}

////////////////////////////end sync.Mutex

func main() {
	var c chan struct{} // nil
	select {
	case <-c: // 阻塞操作
	case c <- struct{}{}: // 阻塞操作
	default:
		fmt.Println("Go here.")
	}
}
