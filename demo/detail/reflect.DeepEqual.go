/*
reflect.DeepEqual(x, y)和x == y的结果可能会不同

1、如果表达式x和y的类型不相同，则函数调用DeepEqual(x, y)的结果总为false，但x == y的估值结果有可能为true
2、如果x和y为（同类型的）两个引用着不同其它值的指针值，则x == y的估值结果总为false，
但函数调用DeepEqual(x, y)的结果可能为true，因为函数reflect.DeepEqual将比较x和y所引用的两个值

3、第三个区别是当x和y处于一个循环引用链中时，DeepEqual调用的结果可能未必正确

4、第四个区别是一个DeepEqual(x, y)调用无论如何不应该产生一个恐慌，
但是如果x和y是两个动态类型相同的接口值并且它们的动态类型是不可比较类型的时候，x == y将产生一个恐慌


*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	func1()

	type Book struct{ page int }
	x := struct{ page int }{123}
	y := Book{123}
	fmt.Println(reflect.DeepEqual(x, y)) // false
	fmt.Println(x == y)                  // true

	z := Book{123}
	fmt.Println(reflect.DeepEqual(&z, &y)) // true
	fmt.Println(&z == &y)                  // false

	type T struct{ p *T }
	t := &T{&T{nil}}
	t.p.p = t                              // form a cyclic reference chain.
	fmt.Println(reflect.DeepEqual(t, t.p)) // true
	fmt.Println(t == t.p)                  // false

	var f1, f2 func() = nil, func() {}
	fmt.Println(reflect.DeepEqual(f1, f1)) // true
	fmt.Println(reflect.DeepEqual(f2, f2)) // false

	var a, b interface{} = []int{1, 2}, []int{1, 2}
	fmt.Println(reflect.DeepEqual(a, b)) // true
	fmt.Println(a == b)                  // panic
}

func func1() {
	a := [1]func(){func() {}}
	b := a
	fmt.Println(reflect.DeepEqual(a, a))       // false
	fmt.Println(reflect.DeepEqual(a[:], a[:])) // true
	fmt.Println(reflect.DeepEqual(a[:], b[:])) // false
	a[0] = nil
	fmt.Println(reflect.DeepEqual(a, a)) // true
	fmt.Println()
}
