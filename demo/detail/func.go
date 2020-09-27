/*
有返回值的函数的调用是一种表达式
一个有且只有一个返回值的函数的每个调用总可以被当成一个单值表达式使用。 比如，它可以被内嵌在其它函数调用中当作实参使用，或者可以被当作其它表达式中的操作数使用。

如果一个有多个返回结果的函数的一个调用的返回结果没有被舍弃，则此调用可以当作一个多值表达式使用在两种场合：
1、此调用可以在一个赋值语句中当作源值来使用，但是它不能和其它源值掺和到一块。
2、此调用可以内嵌在另一个函数调用中当作实参来使用，但是它不能和其它实参掺和到一块。

即， 一个多返回值函数调用表达式不能和其它表达式混用在一个赋值语句的右侧或者另一个函数调用的实参列表中


某些函数调用是在在编译时刻被估值的
unsafe.Sizeof
unsafe.Alignof
unsafe.Offsetof

len,cap,real,imag,complex  常量表达式时才在编译时刻估值
*/

package main

func HalfAndNegative(n int) (int, int) {
	return n / 2, -n
}

func AddSub(a, b int) (int, int) {
	return a + b, a - b
}

func Dummy(values ...int) {}

func main() {
	// 这几行编译没问题。
	AddSub(HalfAndNegative(6))
	AddSub(AddSub(AddSub(7, 5)))
	AddSub(AddSub(HalfAndNegative(6)))
	Dummy(HalfAndNegative(6))
	_, _ = AddSub(7, 5)

	// 下面这几行编译不通过。
	/*
		_, _, _ = 6, AddSub(7, 5)
		Dummy(AddSub(7, 5), 9)
		Dummy(AddSub(7, 5), HalfAndNegative(6))
	*/
}
