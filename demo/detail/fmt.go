/*
[Sp|Fp|P]rintf函数支持位置参数
*/
package main

import (
	"fmt"
)

func main() {
	// The next line prints: coco
	fmt.Printf("%[2]v%[1]v%[2]v%[1]v", "o", "c")
}
