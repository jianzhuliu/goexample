/*
flag标准库包对待布尔命令选项不同于数值和字符串选项
传递程序选项有三种形式。
1、-flag：仅适用于布尔选项。
2、-flag=x：用于任何类型的选项。.
3、-flag x：仅用于非布尔选项。

请注意，使用第一种形式的布尔选项将被视为最后一个选项，其后面的所有项都被视为参数。
*/

package main

import "fmt"
import "flag"

var b = flag.Bool("b", true, "一个布尔选项")
var i = flag.Int("i", 123, "一个整数选项")
var s = flag.String("s", "hi", "一个字符串选项")

func main() {
	flag.Parse()
	fmt.Print("b=", *b, ", i=", *i, ", s=", *s, "\n")
	fmt.Println("arguments:", flag.Args())
}

/*
Output:
go run flag.go -b false -i 789 -s bye arg0 arg1
b=true, i=123, s=hi
arguments: [false -i 789 -s bye arg0 arg1]

这个输出显然不是我们所期望的,修正
go run flag.go -b=false -i 789 -s bye arg0 arg1
b=false, i=789, s=bye
arguments: [arg0 arg1]

或者
go run flag.go -i 789 -s bye -b arg0 arg1
b=true, i=789, s=bye
arguments: [arg0 arg1]



*/
