/*
一个Go源文件可以多次引入同一个包。但是每次的引入名称必须不同。这些相同的包引入引用着同一个包实例
*/
package main

import "fmt"
import "io"
import inout "io"

func main() {
	fmt.Println(&inout.EOF == &io.EOF) // true
}
