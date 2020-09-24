package main

import "fmt"
import "strconv"

func parseInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		// 一些新手Go程序员会认为下一行中声明
		// 的err变量已经在外层声明过了。然而其
		// 实下一行中的b和err都是新声明的变量。
		// 此新声明的err遮挡了外层声明的err。
		b, err := strconv.ParseBool(s)
		if err != nil {
			return 0, err
		}

		// 如果代码运行到这里，一些新手Go程序员
		// 期望着内层的nil err将被返回。但是其实
		// 返回是外层的非nil err。因为内层的err
		// 的作用域到外层if代码块结尾就结束了。
		if b {
			n = 1
		}
	}
	return n, err
}

func main() {
	fmt.Println(parseInt("TRUE"))
}

/*
Output:
1 strconv.Atoi:parsing "TRUE": invalid syntax
*/
