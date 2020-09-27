/*
嵌套的延迟函数调用可以修改外层函数的返回结果
*/

package main

import "fmt"

func F() (r int) {
	defer func() {
		r = 789
	}()

	return 123 // <=> r = 123; return
}

func main() {
	fmt.Println(F()) // 789
}
