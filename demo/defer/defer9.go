package main

import  "fmt"

func func1(){
	s := []string{"a", "b", "c", "d"}
	defer fmt.Println(s) // [a x y d]
	// defer append(s[:1], "x", "y") // 编译错误
	defer func() {
		_ = append(s[:1], "x", "y")
	}()
	fmt.Println("func1 body done")
}

type T int

func (t T) M(n int) T {
  fmt.Println(n)
  return t
}

func func2() {
	var t T
	// t.M(1)是方法调用M(2)的属主实参，因此它
	// 将在M(2)调用被推入延迟调用堆栈之前被估值。
	defer t.M(1).M(2)
	t.M(3).M(4)
	fmt.Println("func2 body done")
}

func func3(){
	i := 1
	defer fmt.Println("func3 defer1 worked",i)
	var f func()	//f == nil 
	defer f() 		//panic
	i++
	fmt.Println("func3 body worked", i)
	f = func(){}	//此处无法改变 panic 发生，defer 声明时，参数已经被估值
}

func main(){
	//内置函数，除 copy, recover 外，调用的返回值不可舍弃
	func1()
	
	fmt.Println("--------------------")
	
	//一个延迟调用的实参也是在此调用被推入延迟调用堆栈之前估值的
	func2()
	
	fmt.Println("--------------------")
	
	// defer 调用 nil 函数时，在 defer 延迟调用时 panic
	func3()
	
}

/*
Output:
func1 body done
[a x y d]
--------------------
1
3
4
func2 body done
2
--------------------
func3 body worked 2
func3 defer1 worked 1
...panic stack

*/

/*
// defer 延迟调用，避免延迟调用堆栈过大，导致资源未被及时释放
func writeManyFiles(files []File) error {
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
//*/