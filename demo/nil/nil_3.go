/*
不同种类的类型的nil值的尺寸很可能不相同
一个类型的所有值的内存布局都是一样的，此类型nil值也不例外（假设此类型的零值使用nil表示）。
所以同一个类型的nil值和非nil值的尺寸是一样的。但是不同类型的nil值的尺寸可能是不一样的。

*/

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var p *struct{} = nil
	fmt.Println(unsafe.Sizeof(p)) // 8

	var s []int = nil
	fmt.Println(unsafe.Sizeof(s)) // 24

	var m map[int]bool = nil
	fmt.Println(unsafe.Sizeof(m)) // 8

	var c chan string = nil
	fmt.Println(unsafe.Sizeof(c)) // 8

	var f func() = nil
	fmt.Println(unsafe.Sizeof(f)) // 8

	var i interface{} = nil
	fmt.Println(unsafe.Sizeof(i)) // 16
}
