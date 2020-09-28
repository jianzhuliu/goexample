/*
strings.Builder类型以及在sync标准库包里的类型的值不推荐被复制

复制bytes.Buffer的值不会在运行时被检查到，也不会被go vet命令所检测到。 千万要小心不要随意这样做
*/

package main

import "strings"
import "sync"

func f(m sync.Mutex) { // warning: f passes lock by value: sync.Mutex
	m.Lock()
	defer m.Unlock()
	// do something ...
}

func main() {
	var b strings.Builder
	b.WriteString("hello ")

	/*
		var b2 = b
		b2.WriteString("world!") // 一个恐慌将在这里产生
	*/

	//go vet noCopy.go
	var sm sync.Mutex
	f(sm)
}
