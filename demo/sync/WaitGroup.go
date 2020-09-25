/*
一个WaitGroup可以在它的一个Wait方法返回之后被重用。
但是请注意，当一个WaitGroup值维护的基数为零时，它的带有正整数实参的Add方法调用不能和它的Wait方法调用并发运行，否则将可能出现数据竞争
请注意wg.Add(delta)、wg.Done()和wg.Wait()分别是(&wg).Add(delta)、(&wg).Done()和(&wg).Wait()的简写形式
*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags)
	rand.Seed(time.Now().UnixNano())

	const N = 5
	var values [N]int32

	var wgA, wgB sync.WaitGroup
	wgA.Add(N)
	wgB.Add(1)

	for i := 0; i < N; i++ {
		i := i
		go func() {
			wgB.Wait() // 等待广播通知
			log.Printf("values[%v]=%v \n", i, values[i])
			wgA.Done()
		}()
	}

	// 下面这个循环保证将在上面的任何一个
	// wg.Wait调用结束之前执行。
	for i := 0; i < N; i++ {
		values[i] = 50 + rand.Int31n(50)
	}
	wgB.Done() // 发出一个广播通知
	wgA.Wait()

	// 所有的元素都保证被初始化了。
	fmt.Println("1---------values:", values)

	//重用 wgA
	for i := 0; i < N; i++ {
		wgA.Add(1) // 将被执行5次
		i := i
		go func() {
			values[i] = 100 + rand.Int31n(100)
			wgA.Done()
		}()
	}

	wgA.Wait()
	// 所有的元素都保证被初始化了。
	fmt.Println("2---------values:", values)
}
