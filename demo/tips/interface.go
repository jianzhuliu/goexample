/*
1、确保并表明一个自定义类型实现了指定的接口类型
2、如何在不导入reflect标准库包的情况下检查一个值是否拥有某个方法
*/

package main

import "fmt"
import "io"

//将一个自定义类型的一个值赋给指定接口类型的一个变量来确保此自定义类型实现了指定接口类型
type MyReader uint16

func NewMyReader() *MyReader {
	var mr MyReader
	return &mr
}

func (mr *MyReader) Read(data []byte) (int, error) {
	switch len(data) {
	default:
		*mr = MyReader(data[0])<<8 | MyReader(data[1])
		return 2, nil
	case 2:
		*mr = MyReader(data[0])<<8 | MyReader(data[1])
	case 1:
		*mr = MyReader(data[0])
	case 0:
	}
	return len(data), io.EOF
}

// 下面三行中的任一行都可以保证类型*MyReader实现
// 了接口io.Reader。
var _ io.Reader = NewMyReader()
var _ io.Reader = (*MyReader)(nil)

func _() { _ = io.Reader(nil).(*MyReader) }

type A int
type B int

func (b B) M(x int) string {
	return fmt.Sprint(b, ": ", x)
}

func check(v interface{}) bool {
	_, has := v.(interface{ M(int) string })
	return has
}

func main() {
	var a A = 123
	var b B = 789
	fmt.Println(check(a)) // false
	fmt.Println(check(b)) // true
}
