/*
内存泄漏
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

///////////////////////////////////////////////////////子字符串造成的暂时性内存泄露 begin
var s0 string // 一个包级变量

func demo() {
	//s := createStringWithLengthOnHeap(1 << 20) // 1M bytes
	s := strings.Repeat("Hello World", 50)
	f(s)
}

// 一个演示目的函数。
func f(s1 string) {
	s0 = s1[:50]
	// 目前，s0和s1共享着承载它们的字节序列的同一个内存块。
	// 虽然s1到这里已经不再被使用了，但是s0仍然在使用中，
	// 所以它们共享的内存块将不会被回收。虽然此内存块中
	// 只有50字节被真正使用，而其它字节却无法再被使用。
}

//为防止上面的f函数产生临时性内存泄露，我们可以将子字符串表达式的结果转换为一个字节切片，然后再转换回来
//此种防止临时性内存泄露的方法不是很高效，因为在此过程中底层的字节序列被复制了两次，其中一次是不必要的
func f_fix1(s1 string) {
	s0 = string([]byte(s1[:50]))
}

//我们可以利用官方Go编译器对字符串衔接所做的优化来防止一次不必要的复制，代价是有一个字节的浪费
//此第二种防止临时性内存泄露的方法有可能在将来会失效，并且它对于其它编译器来说很可能是无效的
func f_fix2(s1 string) {
	s0 = (" " + s1[:50])[1:]
}

//第三种防止临时性内存泄露的方法是使用在Go 1.10种引入的strings.Builder类型来防止一次不必要的复制
func f_fix3(s1 string) {
	var b strings.Builder
	b.Grow(50)
	b.WriteString(s1[:50])
	s0 = b.String()
}

///////////////////////////////////////////////////////子字符串造成的暂时性内存泄露 end

///////////////////////////////////////////////////////子切片造成的暂时性内存泄露 begin

var s02 []int

//当函数g被调用之后，承载着切片s1的元素的内存块的开头大段内存将不再可用（假设没有其它值引用着此内存块）
//同时因为s0仍在引用着此内存块，所以此内存块得不到释放。

func g(s1 []int) {
	// 假设s1的长度远大于30。
	s02 = s1[len(s1)-30:]
}

//在函数g中将30个元素均复制一份，使得切片s0和s1不共享承载底层元素的内存块
func g_fix(s1 []int) {
	s02 = append(s1[:0:0], s1[len(s1)-30:]...)
	// 现在，如果再没有其它值引用着承载着s1元素的内存块，
	// 则此内存块可以被回收了。
}

//只要h函数调用返回的切片仍在被使用中，它的各个元素就不会回收，包括首尾两个已经丢失的元素
// 因此这两个已经丢失的元素引用着的两个int值也不会被回收，即使我们再也无法使用这两个int值。
func h() []*int {
	s := []*int{new(int), new(int), new(int), new(int)}
	// 使用此s切片 ...

	return s[1:3:3]
}

//防止这样的暂时性内存泄露，我们必须重置丢失的元素中的指针
func h_fix() []*int {
	s := []*int{new(int), new(int), new(int), new(int)}
	// 使用此s切片 ...

	s[0], s[len(s)-1] = nil, nil // 重置首尾元素指针
	return s[1:3:3]
}

///////////////////////////////////////////////////////子切片造成的暂时性内存泄露 end

///////////////////////////////////////////////////////不正确地使用终结器（finalizer）而造成的永久性内存泄露 begin

//不要为一个循环引用值组中的值设置终结器。
//顺便说一下，我们不应该把终结器用做析构函数
func memoryLeaking() {
	type T struct {
		v [1 << 20]int
		t *T
	}

	var finalizer = func(t *T) {
		fmt.Println("finalizer called")
	}

	var x, y T

	// 此SetFinalizer函数调用将使x逃逸到堆上。
	runtime.SetFinalizer(&x, finalizer)

	// 下面这行将形成一个包含x和y的循环引用值组。
	// 这有可能造成x和y不可回收。
	x.t, y.t = &y, &x // y也逃逸到了堆上。
}

///////////////////////////////////////////////////////不正确地使用终结器（finalizer）而造成的永久性内存泄露 end

///////////////////////////////////////////////////////延迟调用函数导致的临时性内存泄露 begin

//一个较大的延迟调用堆栈可能会消耗很多内存，而且延迟调用堆栈中尚未执行的延迟调用可能会导致某些资源未被及时释放
//处理大量的文件，则在此函数推出之前，将有大量的文件句柄得不到释放
func writeManyFiles(files []File) error {
	for _, file := range files {
		f, err := os.Open(file.path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(file.content)
		if err != nil {
			return err
		}

		err = f.Sync()
		if err != nil {
			return err
		}
	}

	return nil
}

func writeManyFiles_fix(files []File) error {
	for _, file := range files {
		if err := func() error {
			f, err := os.Open(file.path)
			if err != nil {
				return err
			}
			defer f.Close() // 将在此循环步步尾执行

			_, err = f.WriteString(file.content)
			if err != nil {
				return err
			}

			return f.Sync()
		}(); err != nil {
			return err
		}
	}

	return nil
}

///////////////////////////////////////////////////////延迟调用函数导致的临时性内存泄露 end

/*
1、因为协程被永久阻塞而造成的永久性内存泄露
Go运行时出于两个原因并不杀掉处于永久阻塞状态的协程。
一是有时候Go运行时很难分辨出一个处于阻塞状态的协程是永久阻塞还是暂时性阻塞；
二是有时我们可能故意永久阻塞某些协程。

2、因为没有停止不再使用的time.Ticker值而造成的永久性内存泄露
当一个time.Timer值不再被使用，一段时间后它将被自动垃圾回收掉。
但对于一个不再使用的time.Ticker值，我们必须调用它的Stop方法结束它，否则它将永远不会得到回收


*/
