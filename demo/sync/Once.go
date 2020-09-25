/*
一个sync.Once值被用来确保一段代码在一个并发程序中被执行且仅被执行一次
sync.Once值o，o.Do()（即(&o).Do()的简写形式）
*/
package main

import (
	"log"
	"sync"
)

func main() {
	log.SetFlags(0)

	x := 0
	doSomething := func() {
		x++
		log.Println("Hello")
	}

	var wg sync.WaitGroup
	var once sync.Once
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(doSomething)
			log.Println("world!")
		}()
	}

	wg.Wait()
	log.Println("x =", x) // x = 1
}
