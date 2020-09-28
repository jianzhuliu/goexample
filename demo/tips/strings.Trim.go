/*
标准包strings和bytes里有多个修剪（trim）函数。 这些函数可以被分类为两组：
1、Trim、TrimLeft、TrimRight、TrimSpace、TrimFunc、TrimLeftFunc和TrimRightFunc。
这些函数将修剪首尾所有满足指定（或隐含）条件的utf-8编码的Unicode码点(即rune)。
（TrimSpace隐含了修剪各种空格符。）
这些函数将检查每个开头或结尾的rune值，直到遇到一个不满足条件的rune值为止。

2、TrimPrefix和TrimSuffix。 这两个函数会把指定前缀或后缀的子字符串（或子切片）作为一个整体进行修剪

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "abaay森z众xbbab"
	o := fmt.Println
	o(strings.TrimPrefix(s, "ab")) // aay森z众xbbab
	o(strings.TrimSuffix(s, "ab")) // abaay森z众xbb
	o(strings.TrimLeft(s, "ab"))   // y森z众xbbab
	o(strings.TrimRight(s, "ab"))  // abaay森z众x
	o(strings.Trim(s, "ab"))       // y森z众x
	o(strings.TrimFunc(s, func(r rune) bool {
		return r < 128 // trim all ascii chars
	})) // 森z众
}
