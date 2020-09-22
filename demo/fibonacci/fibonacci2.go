package main

import "fmt"

//匿名函数
func fibonacci() func() int {
	x, y := 0, 1

	return func() int {
		tmp := x
		x, y = y, x+y
		return tmp
	}
}

func main() {
	maxNum := 10
	lambFunc := fibonacci()

	for i := 0; i < maxNum; i++ {
		fmt.Printf("%d ", lambFunc())
	}

	fmt.Println()

}

//Output:
//0 1 1 2 3 5 8 13 21 34 55
