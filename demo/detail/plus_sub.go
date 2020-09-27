/*
递增运算符++和递减运算符--的优先级低于解引用运算符*和取地址运算符&，
解引用运算符和取地址运算符的优先级低于选择器.中的属性选择操作符
*/

package main

import "fmt"

type T struct {
	x int
	y *int
}

func main() {
	var t T
	p := &t.x             // <=> p := &(t.x)
	fmt.Printf("%T\n", p) // *int

	*p++ // <=> (*p)++
	*p-- // <=> (*p)--

	t.y = p
	a := *t.y             // <=> *(t.y)
	fmt.Printf("%T\n", a) // int
}
