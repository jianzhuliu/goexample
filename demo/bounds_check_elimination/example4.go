// go build -gcflags="-d=ssa/check_bce/debug=1" example4.go
package main

import "math/rand"

func fb2(s []int, index int) {
	_ = s[index:] // 第7行： 需要边界检查
	_ = s[:index] // 第8行： 边界检查消除了！
}

func fc2() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	s = s[:4]
	index := rand.Intn(7)
	_ = s[index:] // 第15行： 需要边界检查
	_ = s[:index] // 第16行： 边界检查消除了！
}

func main() {}

/*
go build -gcflags="-d=ssa/check_bce/debug=1" example4.go
# command-line-arguments
.\example4.go:7:7: Found IsSliceInBounds
.\example4.go:15:7: Found IsSliceInBounds
*/