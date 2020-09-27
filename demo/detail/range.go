/*
1、用range遍历nil映射或者nil切片是没问题的，这属于空操作
2、用range遍历nil数组指针时，如果忽略或省略第二个迭代变量，则此遍历是没问题的。遍历中的循环步数为相应数组类型的长度

*/

package main

import "fmt"

var i int

func fa(s []int, n int) int {
	i = n
	for i = 0; i < len(s); i++ {
	}
	return i
}

func fb(s []int, n int) int {
	i = n
	for i = range s {
	}
	return i
}

func func1() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(fa(s, -1), fb(s, -1)) // 6 5
	s = nil
	fmt.Println(fa(s, -1), fb(s, -1)) // 0 -1
}

func main() {
	var s []int // nil
	for range s {
	}

	var m map[string]int // nil
	for range m {
	}

	var a *[5]int // nil
	for i, _ := range a {
		fmt.Print(i)
	}

	fmt.Println()
	func1()
}

/*
Output:
01234
6 5
0 -1
*/
