/*
defer 声明时会先计算确定参数的值，defer 推迟执行的仅是其函数体
*/
package main

import "fmt"

func main() {
	deferIncludeSubFunc()
	fmt.Println()

	i := 0
	defer func(j int) {
		i++
		fmt.Println("main defer1:", i, j) //main defer1:2 0
	}(i)

	defer func(j int) {
		i++
		fmt.Println("main defer2:", i, j) //main defer2:1 0
	}(i)

	fmt.Println("main:", i) //main:0
}

func deferIncludeSubFunc() {
	defer show(1, show(3, 0))
	defer show(2, show(4, 0))
	fmt.Println("deferIncludeSubFunc done")
}

func show(x, y int) int {
	fmt.Println(x)
	return x
}

/*
Output:
3
4
deferIncludeSubFunc done
2
1

main:0
main defer2:1 0
main defer1:2 0



*/
