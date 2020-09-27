/*
以相同实参调用两次errors.New函数返回的两个error值是不相等的。
原因是errors.New函数会复制输入的字符串实参至一个局部变量并取此局部变量的指针作为返回error值的动态值。 两次调用会产生两个不同的指针。
*/

package main

import "fmt"
import "errors"

func main() {
	notfound := "not found"
	a, b := errors.New(notfound), errors.New(notfound)
	fmt.Println(a == b) // false
}
