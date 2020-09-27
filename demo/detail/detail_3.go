/*
栈溢出不可被挽救，它将使程序崩溃
*/
package main

func f() {
	f()
}

func main() {
	defer func() {
		recover() // 无法防止程序崩溃
	}()
	f()
}
