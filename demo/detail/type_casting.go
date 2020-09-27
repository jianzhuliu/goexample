/*
如果两个指针的类型具有不同的底层类型但是它们的基类型却共享相同的底层类型，则这两个指针值可以间接相互转换为对方的类型
*/

package main

type MyInt int64
type Ta *int64
type Tb *MyInt

func main() {
	var a Ta
	var b Tb

	//a = Ta(b) // error: 直接转换是不允许的。

	// 但是间接转换是允许的。
	y := (*MyInt)(b)
	x := (*int64)(y)
	a = x           // 等价于下一行
	a = (*int64)(y) // 等价于下一行
	a = (*int64)((*MyInt)(b))
	_ = a
}
