/*
移位运算中的左类型不确定操作数的类型推断规则取决于右操作数是否是常量
*/

package main

func main() {
}

const M = 2

var _ = 1.0 << M // 编译没问题。1.0将被推断为一个int值。

var N = 2
var _ = 1.0 << N // 编译失败。1.0将被推断为一个float64值。
