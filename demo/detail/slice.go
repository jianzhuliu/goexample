/*
1、从一个基础切片派生出的子切片的长度可能大于基础切片的长度。
2、从一个nil切片中派生子切片是允许的，只要子切片表达式中使用的所有索引都为零，则不会有恐慌产生，结果子切片同样是一个nil切片
3、切片的长度和容量可以被单独修改

*/

package main

import "fmt"
import "reflect"

func func1() {
	s := make([]int, 2, 6)
	fmt.Println(len(s), cap(s)) // 2 6

	reflect.ValueOf(&s).Elem().SetLen(3)
	fmt.Println(len(s), cap(s)) // 3 6

	reflect.ValueOf(&s).Elem().SetCap(5)
	fmt.Println(len(s), cap(s)) // 3 5
}

func main() {
	s := make([]int, 3, 9)
	fmt.Println(len(s)) // 3
	s2 := s[2:7]
	fmt.Println(len(s2)) // 5

	var x []int // nil
	a := x[:]
	b := x[0:0]
	c := x[:0:0]
	// 下一行将打印出三个true。
	fmt.Println(a == nil, b == nil, c == nil)

	fmt.Println()
	func1()
}

/*
Output:
3
5
true true true

2 6
3 6
3 5
*/
