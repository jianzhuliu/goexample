package main

func main() {
	s0 := make([]int, 5, 10)
	// len(s0) == 5, cap(s0) == 10

	index := 8

	// 在Go中，对于一个子切片表达式s[a:b]，a和b须满足
	// 0 <= a <= b <= cap(s);否则，将产生一个恐慌。

	_ = s0[:index]
	// 上一行是安全的不能保证下一行是无需边界检查的。
	// 事实上，下一行将产生一个恐慌，因为起始下标
	// index大于终止下标（即切片s0的长度）。
	_ = s0[index:] // panic
}