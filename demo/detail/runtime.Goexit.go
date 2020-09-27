/*
通过调用runtime.Goexit函数退出一个goroutine。 runtime.Goexit函数没有参数
*/

package main

import "fmt"
import "runtime"

func main() {
	c := make(chan int)
	go func() {
		defer func() { c <- 1 }()
		defer fmt.Println("Go")
		func() {
			defer fmt.Println("C")
			runtime.Goexit()
		}()
		fmt.Println("Java")
	}()
	<-c
}

/*
Output:
C
Go
*/
