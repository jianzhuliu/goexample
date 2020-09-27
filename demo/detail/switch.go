/*
1、在switch和select流程控制代码块中，default分支可以放在所有的case分支之前或者所有的case分支之后，
也可以放在case分支之间

2、switch流程控制代码块中的数字常量case表达式不能重复，但是布尔常量case表达式可以重复

3、switch流程控制代码块里的switch表达式总是被估值为类型确定值

4、switch流程控制代码块中的switch表达式的缺省默认值为类型确定值true（其类型为预定义类型bool）

5、有些case分支代码块必须是显式的

*/
package main

import (
	"fmt"
	"math/rand"
)

//在switch和select流程控制代码块中，default分支可以放在所有的case分支之前或者所有的case分支之后，
//也可以放在case分支之间
func func1() {
	switch n := rand.Intn(3); n {
	case 0:
		fmt.Println("n == 0")
	case 1:
		fmt.Println("n == 1")
	default:
		fmt.Println("n == 2")
	}

	switch n := rand.Intn(3); n {
	default:
		fmt.Println("n == 2")
	case 0:
		fmt.Println("n == 0")
	case 1:
		fmt.Println("n == 1")
	}

	switch n := rand.Intn(3); n {
	case 0:
		fmt.Println("n == 0")
	default:
		fmt.Println("n == 2")
	case 1:
		fmt.Println("n == 1")
	}

	var x, y chan int

	select {
	case <-x:
	case y <- 1:
	default:
	}

	select {
	case <-x:
	default:
	case y <- 1:
	}

	select {
	default:
	case <-x:
	case y <- 1:
	}
}

////////////switch流程控制代码块中的数字常量case表达式不能重复，但是布尔常量case表达式可以重复
func func2() {
	switch false {
	case false:
	case false:
	}

	/*
		switch 123 {
		case 123:
		case 123: // error: duplicate case
		}
	*/
}

///////////////switch流程控制代码块里的switch表达式总是被估值为类型确定值
/*
func func3(){
	switch 123 {
	case int64(123):  // error: 类型不匹配
	case uint32(789): // error: 类型不匹配
	}

}
*/

////////////switch流程控制代码块中的switch表达式的缺省默认值为类型确定值true（其类型为预定义类型bool）
func func4_1() {
	switch { // <=> switch true {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}
}

func False() bool {
	return false
}

func func4_2() {
	switch False(); // <=> switch False();true
	{
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}
}

//////有些case分支代码块必须是显式的
/*
func func5(n, m int) (r int) {
	switch n {
	case 123:
		if m > 0 {
			goto End
		}
		r++

		End: // syntax error: 标签后缺少语句
	default:
		r = 1
	}
	return
}
*/

//case分支代码块必须改成显式的
func func5_fix1(n, m int) (r int) {
	switch n {
	case 123:
		{
			if m > 0 {
				goto End
			}
			r++

		End: // syntax error: 标签后缺少语句
		}
	default:
		r = 1
	}
	return
}

//我们可以在标签End:之后加一个分号
func func5_fix2(n, m int) (r int) {
	switch n {
	case 123:
		if m > 0 {
			goto End
		}
		r++

	End:
		;
	default:
		r = 1
	}
	return
}

func main() {
	func1()
	func2()
	func4_1()
	func4_2()
	func5_fix1(123, 1)
	func5_fix2(123, 1)
}
