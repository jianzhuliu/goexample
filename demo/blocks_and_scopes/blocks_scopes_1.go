package main

import "fmt"

var f = func(b bool) {
	fmt.Print("Goat")
}

func func1() {
	var f = func(b bool) {
		fmt.Print("Sheep")
		if b {
			fmt.Print(" ")
			f(!b) // 此f乃包级变量f也。
		}
	}
	f(true) // 此f为刚声明的局部变量f。
}

func func2() {
	var f func(b bool)
	f = func(b bool) {
		fmt.Print("Sheep")
		if b {
			fmt.Print(" ")
			f(!b) // 现在，此f变为局部变量f了。
		}
	}
	f(true)
}

func main() {
	func1()
	fmt.Println()
	func2()
}

/*
Output:
Sheep Goat
Sheep Sheep
*/
