/*
defer 执行顺序，如同栈，先进后出|后进先出
*/
package main

import "fmt"

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	fmt.Println("main done")
}

/*
Output:
main done
4
3
2
1
*/
