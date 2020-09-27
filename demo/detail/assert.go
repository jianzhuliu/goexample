/*
断言
1、类型断言可以用于将一个接口值转换为另一个接口类型，即使此接口值的类型并未实现另一个接口类型
2、一个失败的类型断言的可选的第二个结果是否被舍弃将影响此类型断言的行为

*/

package main

//import "fmt"

type Foo interface {
	foo()
}

type T int

func (T) foo() {}

func main() {
	var x interface{} = T(123)
	// 下面这两行将编译失败。
	/*
		var _ Foo = x   // error: interface{}类型没有实现Foo类型
		var _ = Foo(x)  // error: interface{}类型没有实现Foo类型
	*/
	// 但是下面这行可以编译通过。
	var _ = x.(Foo) // okay

	var y interface{} = true
	_, _ = y.(int) // 断言失败，但不会导致恐慌。
	//_ = y.(int)    // 断言失败，并导致一个恐慌。
}
