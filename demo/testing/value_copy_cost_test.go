package main

import "testing"

type S struct{ a, b, c, d, e int64 }

var sX = make([]S, 1000)
var sY = make([]S, 1000)
var sZ = make([]S, 1000)
var sumX, sumY, sumZ int64

func Benchmark_Loop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumX = 0
		for j := 0; j < len(sX); j++ {
			sumX += sX[j].a
		}
	}
}

func Benchmark_Range_OneIterVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumZ = 0
		for j := range sY {
			sumZ += sY[j].a
		}
	}
}

func Benchmark_Range_TwoIterVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumY = 0
		for _, v := range sY {
			sumY += v.a
		}
	}
}

/*
在32位系统架构中，一个word为4个字节；而在64位系统架构中，一个word为8个字节
在实践中，我们可以将尺寸不大于4个原生字并且字段数不超过4个的结构体值看作是小尺寸值。复制小尺寸值的代价是比较小的

1byte	bool,int8,uint8(byte)
2byte 	int16,uint16
4byte	int32(rune),uint32,float32
8byte	int64,uint64,float64,complex64
16byte	complex128

1word 	int,uint,uintptr,指针，映射，通道，函数
2word	字符串，接口
3word	切片
数组	元素类型的尺寸 * 长度
结构体 	所有字段类型尺寸之和 + 所有填充的字节数(保证内存地址对齐)

一般来说，在实践中，我们很少使用基类型为切片类型、映射类型、通道类型、函数类型、字符串类型和接口类型的指针类型，因为复制这些类型的值的代价很小。

如果一个数组或者切片的元素类型是一个大尺寸类型，我们应该避免在for-range循环中使用双循环变量来遍历这样的数组或者切片类型的值中的元素。 因为，在遍历过程中，每个元素将被复制给第二个循环变量一次
*/
