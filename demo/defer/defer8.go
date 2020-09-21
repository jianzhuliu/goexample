/*
test example
*/

package main

import "fmt"

func func1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()

	return t
}

func func2(i int) int {
	t := i
	defer func() {
		t += 3
	}()

	return t
}

func func3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func func4() (t int) {
	defer func(i int) {
		fmt.Println("func4 i:", i)
		fmt.Println("func4 t:", t)
	}(t)
	t = 1
	return 2
}

func main() {
	fmt.Println("func1:", func1(1))
	fmt.Println("func2:", func2(1))
	fmt.Println("func3:", func3(1))
	func4()
}

/*
Output:
func1:4
func2:1
func3:3
func4 i:0
func4 j:2

*/
