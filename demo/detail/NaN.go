/*
把以NaN做为键值的条目放如映射就宛如把条目放入黑洞一样
以NaN作为键值的元素也可以通过类似fmt.Print的函数打印出来
*/

package main

import "fmt"
import "math"

func func1(){
	fmt.Println()
	var a = math.Sqrt(-1.0)
	fmt.Println(a)      // NaN
	fmt.Println(a == a) // false

	var x = 0.0
	var y = 1.0 / x
	var z = 2.0 * y
	fmt.Println(y, z, y == z) // +Inf +Inf true
}

func main() {
	var a = math.NaN()
	fmt.Println(a) // NaN

	var m = map[float64]int{}
	m[a] = 123
	v, present := m[a]
	fmt.Println(v, present) // 0 false
	m[a] = 789
	v, present = m[a]
	fmt.Println(v, present) // 0 false

	fmt.Println(m) // map[NaN:789 NaN:123]
	delete(m, a)   // no-op
	fmt.Println(m) // map[NaN:789 NaN:123]

	for k, v := range m {
		fmt.Println(k, v)
	}
	// the above loop outputs:
	// NaN 123
	// NaN 789
	
	func1()
}
