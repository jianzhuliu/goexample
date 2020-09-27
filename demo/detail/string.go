/*

 */

package main

import "fmt"

func func1() {
	fmt.Println()
	h := []byte("Hello")
	hc := append(h, " C"...)
	_ = append(h, " Go"...)
	fmt.Printf("%d, %d, %s\n", cap(h), len(h), hc)
}

func func1_fix() {
	//fmt.Println()
	h := []byte("Hello")
	h = h[:len(h):len(h)] // need this line
	hc := append(h, " C"...)
	_ = append(h, " Go"...)
	fmt.Printf("%d, %d, %s\n", cap(h), len(h), hc)
}

func main() {
	s := "a"
	x := []byte(s)              // len(s) == 1
	fmt.Println(cap([]byte(s))) // 32
	fmt.Println(cap(x))         // 8
	fmt.Println(x)
	//如果最后一个fmt.Println行被删除，在其前面的两行会打印相同的值32，否则，一个打印32，一个打印8

	func1()
	func1_fix()
}
