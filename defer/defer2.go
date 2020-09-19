/*
defer 与 return, return 先执行，defer后执行
*/
package main

import "fmt"

func deferAndReturn() int {
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(2)
	}()

	return returnFunc()
}

func returnFunc() int {
	fmt.Println("return")
	return 0
}

func main() {
	deferAndReturn()
}

/*
Output:
return
2
1
*/
