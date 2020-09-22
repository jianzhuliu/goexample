package main

import "fmt"

//递归处理
func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	maxNum := 10
	for i := 0; i < maxNum; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}

	fmt.Println()
}

//Output:
//0 1 1 2 3 5 8 13 21 34 55