/*
很多编译器优化（在编译时刻）和CPU处理器优化（在运行时刻）会常常调整指令执行顺序，从而使得指令执行顺序和代码中指定的顺序不太一致
源文件中的代码行在运行时刻并非总是按照它们的出现次序被执行
*/
package main

import "log"
import "runtime"

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
	// a 与 done 赋值语句，在内存中的指令执行顺序有可能是 done 在 a 之前
	// 没有做同步处理，存在数据竞争，使用 go run -race 打印查看

	if done {
		log.Println(len(a)) // 如果被打印出来，它总是12
	}
}

func main() {
	go setup()

	for !done {
		runtime.Gosched()
	}
	log.Println(a) // 期待的打印结果：hello, world
}
