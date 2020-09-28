/*
编译时刻断言技巧
*/
package main

const N = 35
const M = 10

//为了避免包级变量消耗太多的内存，我们可以把断言代码放在一个名为空标识符的函数体中
func _() {
	var _ = map[bool]int{false: 0, N >= M: 1}
	var _ [N - M]int
}

//用来在编译时刻保证常量N不小于另一个常量M
// 下面任一行均可保证N >= M
func _(x []int)    { _ = x[N-M] }
func _()           { _ = []int{N - M: 0} }
func _([N - M]int) {}

var _ [N - M]int

const _ uint = N - M

type _ [N - M]int

// 如果M和N都是正整数常量，则我们也可以使用下一行所示的方法。
var _ uint = N/M - 1

//此点子利用了容器组合字面量中不能出现重复的常量键值这一规则。
var _ = map[bool]struct{}{false: struct{}{}, N >= M: struct{}{}}

//它也可以不必很冗长，但需要多消耗一点（完全可以忽略的）内存
var _ = map[bool]int{false: 0, N >= M: 1}

/*
//下面是断言两个整数常量相等的方法
var _ [N-M]int; var _ [M-N]int
type _ [N-M]int; type _ [M-N]int
const _, _ uint = N-M, M-N
func _([N-M]int, [M-N]int) {}

var _ = map[bool]int{false: 0, M==N: 1}

var _ = [1]int{M-N: 0} // 唯一被允许的元素索引下标为0
var _ = [1]int{}[M-N]  // 唯一被允许的元素索引下标为0

var _ [N-M]int = [M-N]int{}
*/

func main() {
}
