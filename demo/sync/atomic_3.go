/*
SwapT函数调用和StoreT函数调用类似，但是返回修改之前的旧值（因此称为置换操作）。

一个CompareAndSwapT函数调用仅在新值和旧值相等的情况下才会执行修改操作，并返回true；否则立即返回false
*/
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var n int64 = 123
	var old = atomic.SwapInt64(&n, 789)
	fmt.Println(n, old) // 789 123
	swapped := atomic.CompareAndSwapInt64(&n, 123, 456)
	fmt.Println(swapped) // false
	fmt.Println(n)       // 789
	swapped = atomic.CompareAndSwapInt64(&n, 789, 456)
	fmt.Println(swapped) // true
	fmt.Println(n)       // 456
}
