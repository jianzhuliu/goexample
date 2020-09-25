package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	"sync"
)

//数据的生成
func RandomGenerator() <-chan uint64 {
	c := make(chan uint64)
	go func() {
		rnds := make([]byte, 8)
		for {
			_, err := rand.Read(rnds)
			if err != nil {
				close(c)
				break
			}
			c <- binary.BigEndian.Uint64(rnds)
		}
	}()
	return c
}

//数据聚合 old
func AggregatorOld(inputs ...<-chan uint64) <-chan uint64 {
	out := make(chan uint64)
	for _, in := range inputs {
		go func(in <-chan uint64) {
			for {
				out <- <-in // <=> out <- (<-in)
			}
		}(in)
	}
	return out
}

//数据聚合 new
func Aggregator(inputs ...<-chan uint64) <-chan uint64 {
	output := make(chan uint64)
	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		go func(int <-chan uint64) {
			defer wg.Done()
			// 如果通道in被关闭，此循环将最终结束。
			for x := range in {
				output <- x
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

//数据分流
func Divisor(input <-chan uint64, outputs ...chan<- uint64) {
	for _, out := range outputs {
		go func(o chan<- uint64) {
			for {
				o <- <-input // <=> o <- (<-input)
			}
		}(out)
	}
}

//数据合成
func Composor(inA, inB <-chan uint64) <-chan uint64 {
	output := make(chan uint64)
	go func() {
		for {
			a1, b, a2 := <-inA, <-inB, <-inA
			output <- a1 ^ b&a2
		}
	}()
	return output
}

//数据复制
func Duplicator(in <-chan uint64) (<-chan uint64, <-chan uint64) {
	outA, outB := make(chan uint64), make(chan uint64)
	go func() {
		for x := range in {
			outA <- x
			outB <- x
		}
	}()
	return outA, outB
}

//数据计算/分析
func Calculator(in <-chan uint64, out chan uint64) <-chan uint64 {
	if out == nil {
		out = make(chan uint64)
	}
	go func() {
		for x := range in {
			out <- ^x
		}
	}()
	return out
}

//数据验证/过滤
func Filter0(input <-chan uint64, output chan uint64) <-chan uint64 {
	if output == nil {
		output = make(chan uint64)
	}
	go func() {
		bigInt := big.NewInt(0)
		for x := range input {
			bigInt.SetUint64(x)
			if bigInt.ProbablyPrime(1) {
				output <- x
			}
		}
	}()
	return output
}

func Filter(input <-chan uint64) <-chan uint64 {
	return Filter0(input, nil)
}

//数据服务/存盘
func Printer(input <-chan uint64) {
	for x := range input {
		fmt.Println(x)
	}
}

////////////////////////////////

func func1() {
	Printer(
		Filter(
			Calculator(
				RandomGenerator(), nil,
			),
		),
	)
}

func func2() {
	filterA := Filter(RandomGenerator())
	filterB := Filter(RandomGenerator())
	filterC := Filter(RandomGenerator())
	filter := Aggregator(filterA, filterB, filterC)
	calculatorA := Calculator(filter, nil)
	calculatorB := Calculator(filter, nil)
	calculator := Aggregator(calculatorA, calculatorB)
	Printer(calculator)
}

func func3() {
	c1 := make(chan uint64, 100)
	Filter0(RandomGenerator(), c1) // filterA
	Filter0(RandomGenerator(), c1) // filterB
	Filter0(RandomGenerator(), c1) // filterC
	c2 := make(chan uint64, 100)
	Calculator(c1, c2) // calculatorA
	Calculator(c1, c2) // calculatorB
	Printer(c2)
}

func main() {
	//func1()
	//func2()
	func3()
}
