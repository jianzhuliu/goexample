/*
尽量避免将大尺寸的值包裹在接口值中
*/
package main

import "fmt"

func main() {
	var a [1000]int

	// 这两行的开销相对较大，因为数组a中的元素都将被复制。
	fmt.Println(a)
	fmt.Printf("Type of a: %T\n", a)

	// 这两行的开销较小，数组a中的元素没有被复制。
	fmt.Printf("%v\n", a[:])
	fmt.Println("Type of a:", fmt.Sprintf("%T", &a)[1:])
}
