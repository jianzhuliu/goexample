package main

import "fmt"

func main() {
	f := func(n int) int {
		fmt.Printf("f(%v) is called.\n", n)
		return n
	}

	//各个case关键字后跟随的表达式将按照从上到下和从左到右的顺序进行估值，直到某个比较结果为true为止
	switch x := f(3); x + f(4) {
	default:
	case f(5):
	case f(6), f(7), f(8):
	case f(9), f(10):
	}
}

/*
Output:
f(3) is called.
f(4) is called.
f(5) is called.
f(6) is called.
f(7) is called.
*/
