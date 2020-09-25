package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const N = 10
	var values [N]string

	cond := sync.NewCond(&sync.RWMutex{})
	cond.L.Lock()

	for i := 0; i < N; i++ {
		d := time.Second * time.Duration(rand.Intn(10)) / 10
		go func(i int) {
			time.Sleep(d)
			cond.L.(*sync.RWMutex).RLock()
			values[i] = string('a' + i)
			cond.L.(*sync.RWMutex).RUnlock()
			cond.Signal()
		}(i)
	}

	// 此函数必须在cond.L被锁定的时候调用。
	checkCondition := func() bool {
		log.Println(values)
		for i := 0; i < N; i++ {
			if values[i] == "" {
				return false
			}
		}
		return true
	}
	for !checkCondition() {
		cond.Wait() // 必须在cond.L被锁定的时候调用
	}
	cond.L.Unlock()
}
