/*
1、从nil映射中读取元素不会导致崩溃，读取结果是一个零元素值
2、从一个nil映射中删除一个条目不会导致崩溃，这是一个空操作

*/

package main

import "fmt"

//函数Foo1和Foo2是等价的，但是函数Foo2比函数Foo1简洁得多
func Foo1(m map[string]int) int {
	if m != nil {
		return m["foo"]
	}
	return 0
}

func Foo2(m map[string]int) int {
	return m["foo"]
}

func main() {
	var m map[string]int // nil
	fmt.Println(Foo1(m))
	fmt.Println(Foo2(m))
	delete(m, "foo")
}

/*
Outpt:
0
0

*/
