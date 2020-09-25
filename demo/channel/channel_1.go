/*
管道实现加锁，解锁

在一个通知用例中，我们并不关心回应的值，我们只关心回应是否已发生，
所以常使用空结构体类型struct{}来做为通道的元素类型，因为空结构体类型的尺寸为零，能够节省一些内存（虽然常常很少量）
*/
package main

import "fmt"

var initNum int

//容量为1，才能保证实现加解锁
var mutex chan struct{} = make(chan struct{}, 1)

//加锁
func lock() {
	mutex <- struct{}{}
}

//解锁
func unlock() {
	<-mutex
}

//自增操作
func incr() {
	lock()
	initNum++
	unlock()
}

func main() {
	f := func(done chan struct{}) {
		for i := 0; i < 1000; i++ {
			incr()
		}

		//执行完毕，通知
		done <- struct{}{}
	}

	num := 2
	done := make(chan struct{}, num)
	for i := 0; i < num; i++ {
		go f(done)
	}

	for i := 0; i < num; i++ {
		<-done
	}

	fmt.Println(initNum)
}
