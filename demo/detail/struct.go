/*
在结构体值的比较中，名为空标识符的字段将被忽略
*/

package main

import "fmt"

type T struct {
	_ int
	_ bool
}

func main() {
	var t1 = T{123, true}
	var t2 = T{789, false}
	fmt.Println(t1 == t2) // true
}
