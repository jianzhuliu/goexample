/*
defer 返回值
*/

package main

import "fmt"

//匿名返回值且非指针类型, defer 内不可修改返回值变量
func func1() int {
	var i int
	defer func() {
		i++
		fmt.Println("func1 defer1:", i) //func1 defer1:2
	}()

	defer func() {
		i++
		fmt.Println("func1 defer2:", i) //func1 defer2:1
	}()

	return i //i初始化为0,且 ret
}

//有名返回值，defer 内可改变返回值
func func2() (i int) { //i初始化为0
	defer func() {
		i++
		fmt.Println("func2 defer1:", i) //func2 defer1:2
	}()

	defer func() {
		i++
		fmt.Println("func2 defer2:", i) //func2 defer2:1
	}()

	return i
}

//返回指针变量，defer 可修改指针变量指向的变量值
func func3() *int {
	var i int
	defer func() {
		i++
		fmt.Println("func3 defer1:", i, &i) //func3 defer1:2 0xc0000a2090
	}()

	defer func() {
		i++
		fmt.Println("func3 defer2:", i, &i) //func3 defer2:1 0xc0000a2090
	}()

	return &i
}

func main() {
	v1 := func1()
	fmt.Println("main func1:", v1) //main func1:0
	fmt.Println()

	v2 := func2()
	fmt.Println("main func2:", v2) //main func2:2
	fmt.Println()

	v3 := func3()
	fmt.Println("main func3:", *v3, v3) //main func3:2 0xc0000a2090
	fmt.Println()
}

/*
基于 go 1.15.2 版本
Output:
func1 defer2:1
func1 defer1:2
main func1:0

func2 defer2:1
func2 defer1:2
main func2:2

func3 defer2:1 0xc0000a2090
func3 defer1:2 0xc0000a2090
main func3:2 0xc0000a2090

*/
