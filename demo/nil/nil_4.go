/*
 */

package main

import "fmt"

///////////////////////////1\两个不同类型的nil值可能不能相互比较
// error: 类型不匹配
var _ = (*int)(nil) == (*bool)(nil)

// error: 类型不匹配
var _ = (chan int)(nil) == (chan bool)(nil)

//////////////////////////////2\以下合法
type IntPtr *int

// 类型IntPtr的底层类型为*int。
var _ = IntPtr(nil) == (*int)(nil)

// 任何类型都实现了interface{}类型。
var _ = (interface{})(nil) == (*int)(nil)

// 一个双向通道可以隐式转换为和它的
// 元素类型一样的单项通道类型。
var _ = (chan int)(nil) == (chan<- int)(nil)
var _ = (chan int)(nil) == (<-chan int)(nil)

//////////////////////////////3\同一个类型的两个nil值可能不能相互比较

//映射类型、切片类型和函数类型是不支持比较类型
var _ = ([]int)(nil) == ([]int)(nil)
var _ = (map[string]int)(nil) == (map[string]int)(nil)
var _ = (func())(nil) == (func())(nil)

//但是，映射类型、切片类型和函数类型的任何值都可以和类型不确定的裸nil标识符比较
// 这几行编译都没问题。
var _ = ([]int)(nil) == nil
var _ = (map[string]int)(nil) == nil
var _ = (func())(nil) == nil

func main() {
	//////////////////////////////////4\两个nil值可能并不相等
	//如果可被比较的两个nil值中的一个的类型为接口类型，而另一个不是，则比较结果总是false
	//原因是，在进行此比较之前，此非接口nil值将被转换为另一个nil值的接口类型，从而将此比较转化为两个接口值的比较
	//一个nil接口值中什么也没包裹，但是一个包裹了nil非接口值的接口值并非什么都没包裹
	fmt.Println((interface{})(nil) == (*int)(nil)) // false

	/////////////////////////////////////5\访问nil映射值的条目不会产生恐慌
	fmt.Println((map[string]int)(nil)["key"]) // 0
	fmt.Println((map[int]bool)(nil)[123])     // false
	fmt.Println((map[int]*int64)(nil)[123])   // <nil>

	///////////////////////////////////6\range关键字后可以跟随nil通道、nil映射、nil切片和nil数组指针

	//遍历nil映射和nil切片的循环步数均为零。

	//遍历一个nil数组指针的循环步数为对应数组类型的长度。 //但是，如果此数组类型的长度不为零并且第二个循环变量未被舍弃或者忽略，则对应for-range循环将导致一个恐慌。）

	//遍历一个nil通道将使当前协程永久阻塞。

	//比如，下面的代码将输出0、1、2、3和4后进入阻塞状态。 Hello、world和Bye不会被输出。

	for range []int(nil) {
		fmt.Println("Hello")
	}

	for range map[string]string(nil) {
		fmt.Println("world")
	}

	for i := range (*[5]int)(nil) {
		fmt.Println(i)
	}

	for range chan bool(nil) { // 阻塞在此
		fmt.Println("Bye")
	}

}
