/*
在某些很少见的场景中，圆括号是必需的
*/

package main

type T struct{ x, y int }

func main() {
	// 因为{}的烦扰，下面这三行均编译失败。
	/*
		if T{} == T{123, 789} {}
		if T{} == (T{123, 789}) {}
		if (T{}) == T{123, 789} {}
		var _ = func()(nil) // nil被认为是一个类型
	*/

	// 必须加上一对小括号()才能编译通过。
	if (T{} == T{123, 789}) {
	}
	if (T{}) == (T{123, 789}) {
	}
	var _ = (func())(nil) // nil被认为是一个值
}
