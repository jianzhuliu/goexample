/*
defer panic
*/

package main

import "fmt"

func main() {
	defer fmt.Println("defer:before panic 1")
	defer fmt.Println("defer:before panic 2")
	panic("panic error") // panic 之后语句不再执行
	defer fmt.Println("defer:after panic, never execute")
}

/*
Output:
defer:before panic 2
defer:before panic 1
panic:panic error

//....stack info
exit status 2

*/
