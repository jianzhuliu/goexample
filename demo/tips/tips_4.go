/*
是否每一个零值在内存中占据的字节都是零
对于大部分类型，答案是肯定的。不过事实上，这依赖于编译器。
例如，对于标准编译器，对于某些字符串类型的零值，此结论并不十分正确
*/

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s1 string
	fmt.Println(s1 == "")                         // true
	fmt.Println(*(*uintptr)(unsafe.Pointer(&s1))) // 0
	var s2 = "abc"[0:0]
	fmt.Println(s2 == "")                         // true
	fmt.Println(*(*uintptr)(unsafe.Pointer(&s2))) // 4869856
	fmt.Println(s1 == s2)                         // true
}
