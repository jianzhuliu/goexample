/*
nil是一个预声明的标识符
预声明的nil标识符可以表示很多种类型的零值
在Go中，预声明的nil可以表示下列种类（kind）的类型的零值：
1\指针类型（包括类型安全和非类型安全指针）
2\映射类型
3\切片类型
4\函数类型
5\通道类型
6\接口类型

预声明标识符nil没有默认类型
Go中其它的预声明标识符都有各自的默认类型，比如
1\预声明标识符true和false的默认类型均为内置类型bool。
2\预声明标识符iota的默认类型为内置类型int。

*/

package main

//事实上，预声明标识符nil是Go中唯一一个没有默认类型的类型不确定值。
// 我们必须在代码中提供足够的信息以便让编译器能够推断出一个类型不确定的nil值的期望类型
func main() {
	// 代码中必须提供充足的信息来让编译器推断出某个nil的类型。
	_ = (*struct{})(nil)
	_ = []int(nil)
	_ = map[int]bool(nil)
	_ = chan string(nil)
	_ = (func())(nil)
	_ = interface{}(nil)

	// 下面这一组和上面这一组等价。
	var _ *struct{} = nil
	var _ []int = nil
	var _ map[int]bool = nil
	var _ chan string = nil
	var _ func() = nil
	var _ interface{} = nil

	// 下面这行编译不通过。
	var _ = nil
}
