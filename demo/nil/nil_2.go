/*
nil不是一个关键字
预声明标识符nil可以被更内层的同名标识符所遮挡。
*/

package main

import "fmt"

func main() {
	nil := 123
	fmt.Println(nil) // 123

	// 下面这行编译报错，因为此行中的nil是一个int值。
	var _ map[string]int = nil
}
