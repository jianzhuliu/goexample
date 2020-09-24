package main

//按照声明时的顺序，初始化。处在依赖关系的，后初始化
var (
	_ = f("w", x)
	x = f("x", z)
	y = f("y")
	z = f("z")
)

func f(s string, deps ...int) int {
	println(s)
	return 0
}

func main() {
	f("main done")
}

/*
Output:
y
z
x
w
main done
*/
