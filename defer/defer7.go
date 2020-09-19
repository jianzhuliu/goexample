/*
panic defer recover
*/
package main

import "fmt"

func func1() {

	defer func() {
		fmt.Println("func1 defer1 recover")
		if err := recover(); err != nil {
			fmt.Println("defer1 catch error:", err)
		}
	}()

	defer func() {
		fmt.Println("func1 defer2 panic")
		panic("func1 defer2 new error") //覆盖掉已有的 panic 信息
	}()

	defer fmt.Println("func1 defer3 do nothing")

	panic("panic error")
}

func main() {
	func1()
	fmt.Println("main done")
}

/*
Output:
func1 defer3 do nothing
func1 defer2 panic
func1 defer1 recover
defer1 catch error:func1 defer2 new error
main done

*/
