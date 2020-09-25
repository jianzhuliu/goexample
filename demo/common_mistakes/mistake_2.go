package main

import (
	"runtime"
	"time"
)

func main() {
	var a []int // nil
	var b bool  // false
	//c := make(chan struct{})

	// 一个匿名协程。
	go func() {
		a = make([]int, 3)
		b = true // 写入b
		//c <- struct{}{}
	}()

	for !b { // 读取b
		time.Sleep(time.Second)
		runtime.Gosched()
	}

	//<-c
	a[0], a[1], a[2] = 0, 1, 2 // 可能会发生恐慌
}

/*
源文件中的代码行在运行时刻并非总是按照它们的出现次序被执行
1、首先，主协程中对变量b的读取和匿名协程中的对变量b的写入可能会产生数据竞争
2、其次，在主协程中，条件b == true成立并不能确保条件a != nil也成立。
编译器和CPU可能会对调整此程序中匿名协程中的某些指令的顺序已获取更快的执行速度。
所以，站在主协程的视角看，对变量b的赋值可能会发生在对变量a的赋值之前，这将造成在修改a的元素时a依然为一个nil切片

修正方法，可以采用 sync或者管道做个同步
*/
