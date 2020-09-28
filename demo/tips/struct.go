/*
1、如何使一个结构体类型不可比较
要避免一个自定义的结构体类型被用做一个映射的键值类型，可以放置一个非导出的零尺寸的不可比较类型的字段在结构体类型中以使此结构体类型不可比较


*/

package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	dummy [0]func()
	//_	[0]func()
	AnotherField int
}

//var x map[T]int // 编译错误：非法的键值类型

func main() {
	type T1 struct {
		a struct{}
		x int64
	}
	fmt.Println(unsafe.Sizeof(T1{})) // 8

	type T2 struct {
		x int64
		a struct{}
	}
	fmt.Println(unsafe.Sizeof(T2{})) // 16

	//var a, b T
	//_ = a == b // 编译错误：非法的比较
}
