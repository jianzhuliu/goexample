package main

import "fmt"

//多路复用，管道模式
func fibonacci(a chan int, c chan struct{}) {
	x, y := 0, 1
	for {
		select {
		case a <- x:
			x, y = y, x+y
		case <-c:
			return
		}
	}
}

func main() {
	maxNum := 10
	a, c := make(chan int, maxNum), make(chan struct{})

	go fibonacci(a, c)

	for i := 0; i < maxNum; i++ {
		fmt.Printf("%d ", <-a)
	}

	//使用完毕，主动关闭通道
	close(c)

	fmt.Println()
}

//Output:
//0 1 1 2 3 5 8 13 21 34 55
