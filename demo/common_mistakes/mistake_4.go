package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var x int32 = 0
	for i := 0; i < 100; i++ {
		//wg.Add(1)
		go func() {
			wg.Add(1) //我们应该将对Add方法的调用移出匿名协程之外，使得任何一个Done方法调用都确保发生在唯一的Wait方法调用返回之前。
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
	}

	fmt.Println("等待片刻...")
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&x))
}
