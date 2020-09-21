/*
defer 作用域
*/
package main

import (
	"fmt"
	"os"
	"time"
)

func func1() {
	defer func() {
		fmt.Println("func1 defer1")
	}()

	defer func() {
		fmt.Println("func1 defer2")
	}()

	panic("func1 error")
}

func main() {
	//panic("error3")	//(3) 导致 defer 不执行
	//os.Exit(1)		//(4) 导致 defer 不执行

	defer func() {
		fmt.Println("main defer")
	}()

	//func1()

	//go func() {panic("error2")}()	//(2) 导致 defer 不执行
	//panic("error1")				//(1) defer 可执行
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main")
	os.Exit(1) //(5) 导致 defer 不执行
}

/*
main
exit status 1
*/
