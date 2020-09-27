/*
我们应该使用os.IsNotExist(err)而不是err == os.ErrNotExist来检查文件是否存在
更推荐使用errors.Is(err, os.ErrNotExist)来检查文件是否存在。
*/

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("a-nonexistent-file.abcxyz")
	fmt.Println(os.IsNotExist(err))    // true
	fmt.Println(err == os.ErrNotExist) // false

	fmt.Println(errors.Is(err, os.ErrNotExist)) // true
}
