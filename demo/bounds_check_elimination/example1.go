// go build -gcflags="-d=ssa/check_bce/debug=1" example1.go
package main

func f1(s []int) {
	_ = s[0] // 第5行： 需要边界检查
	_ = s[1] // 第6行： 需要边界检查
	_ = s[2] // 第7行： 需要边界检查
}

func f2(s []int) {
	_ = s[2] // 第11行： 需要边界检查
	_ = s[1] // 第12行： 边界检查消除了！
	_ = s[0] // 第13行： 边界检查消除了！
}

func f3(s []int, index int) {
	_ = s[index] // 第17行： 需要边界检查
	_ = s[index] // 第18行： 边界检查消除了！
}

func f4(a [5]int) {
	_ = a[4] // 第22行： 边界检查消除了！
}

func main() {}

/*
go build -gcflags="-d=ssa/check_bce/debug=1" example1.go
# command-line-arguments
.\example1.go:5:7: Found IsInBounds
.\example1.go:6:7: Found IsInBounds
.\example1.go:7:7: Found IsInBounds
.\example1.go:11:7: Found IsInBounds
.\example1.go:17:7: Found IsInBounds
*/