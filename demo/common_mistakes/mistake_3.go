package main

import (
	"fmt"
	"time"
)

var x = 0

func main() {
	var num = 123
	var p = &num

	c := make(chan int)

	go func() {
		c <- *p + x //123
		//c <- *p 		//789
	}()

	time.Sleep(time.Second)
	num = 789
	fmt.Println(<-c)
}

/*
此程序中存在数据竞争。表达式*p的估值可能发生在赋值num = 789之前、之后、或者同时。 time.Sleep调用并不能保证*p的估值发生在此赋值之后
*/
