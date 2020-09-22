package main

import "fmt"

//管道形式
func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	//制造完毕就关闭管道
	close(c)
}

func main() {
	maxNum := 10
	c := make(chan int, maxNum)
	go fibonacci(maxNum, c)

	for i := 0; i < maxNum; i++ {
		fmt.Printf("%d ", <-c)
	}

	fmt.Println()
}

//Output:
//0 1 1 2 3 5 8 13 21 34 55
