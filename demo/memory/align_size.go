package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

var (
	b bool

	i   int
	i8  int8
	i16 int16
	i32 int32
	i64 int64

	ui   uint
	ui8  uint8
	ui16 uint16
	ui32 uint32
	ui64 uint64

	f32 float32
	f64 float64

	str string
)

//打印格式化处理
var (
	formatOne string = "%-10s%-25v%-25v"
	formatTwo string = "%-20v%-20v\n"
	format    string = fmt.Sprintf("%s%s", formatOne, formatTwo)
)

//通过 反射获取对齐尺寸及字节大小
func showAlignSize(m interface{}) {
	t := reflect.TypeOf(m)
	fmt.Printf(formatOne, t.Kind(), t.Size(), t.Align())
}

//基础类型对齐保证及字节大小
func showBasicTypeAlignAndSize() {
	fmt.Printf(format, "kind", "reflect.TypeOf().Size()", "reflect.TypeOf().Align()", "unsafe.Sizeof()", "unsafe.Alignof()")

	showAlignSize(b)
	fmt.Printf(formatTwo, unsafe.Sizeof(b), unsafe.Alignof(b))

	showAlignSize(i8)
	fmt.Printf(formatTwo, unsafe.Sizeof(i8), unsafe.Alignof(i8))

	showAlignSize(ui8)
	fmt.Printf(formatTwo, unsafe.Sizeof(ui8), unsafe.Alignof(ui8))

	showAlignSize(i16)
	fmt.Printf(formatTwo, unsafe.Sizeof(i16), unsafe.Alignof(i16))

	showAlignSize(ui16)
	fmt.Printf(formatTwo, unsafe.Sizeof(ui16), unsafe.Alignof(ui16))

	showAlignSize(i32)
	fmt.Printf(formatTwo, unsafe.Sizeof(i32), unsafe.Alignof(i32))

	showAlignSize(ui32)
	fmt.Printf(formatTwo, unsafe.Sizeof(ui32), unsafe.Alignof(ui32))

	showAlignSize(f32)
	fmt.Printf(formatTwo, unsafe.Sizeof(f32), unsafe.Alignof(f32))

	showAlignSize(i64)
	fmt.Printf(formatTwo, unsafe.Sizeof(i64), unsafe.Alignof(i64))

	showAlignSize(ui64)
	fmt.Printf(formatTwo, unsafe.Sizeof(ui64), unsafe.Alignof(ui64))

	showAlignSize(f64)
	fmt.Printf(formatTwo, unsafe.Sizeof(f64), unsafe.Alignof(f64))

	showAlignSize(ui)
	fmt.Printf(formatTwo, unsafe.Sizeof(ui), unsafe.Alignof(ui))

	showAlignSize(i)
	fmt.Printf(formatTwo, unsafe.Sizeof(i), unsafe.Alignof(i))

	showAlignSize(str)
	fmt.Printf(formatTwo, unsafe.Sizeof(str), unsafe.Alignof(str))

}

func main() {
	showBasicTypeAlignAndSize()
	fmt.Println()
	showStructAlign()
}

type T1 struct {
	a int8

	// 在64位架构上，为了让字段b的地址为8字节对齐，
	// 需在这里填充7个字节。在32位架构上，为了让
	// 字段b的地址为4字节对齐，需在这里填充3个字节。

	b int64
	c int16

	// 为了让类型T1的尺寸为T1的对齐保证的倍数，
	// 在64位架构上需在这里填充6个字节，在32架构
	// 上需在这里填充2个字节。
}

// 类型T1的尺寸在64位架构上位24个字节（1+7+8+2+6），
// 在32位架构上为16个字节（1+3+8+2+2）。

type T2 struct {
	a int8

	// 为了让字段c的地址为2字节对齐，
	// 需在这里填充1个字节。

	c int16

	// 在64位架构上，为了让字段b的地址为8字节对齐，
	// 需在这里填充4个字节。在32位架构上，不需填充
	// 字节即可保证字段b的地址为4字节对齐的。

	b int64
}

// 类型T2的尺寸在64位架构上位16个字节（1+1+2+4+8），
// 在32位架构上为12个字节（1+1+2+8）。

func showStructAlignAndSize(m interface{}) {
	fmt.Printf("%-10s%-25v%-25v\n", "kind", "reflect.TypeOf().Size()", "reflect.TypeOf().Align()")
	t := reflect.TypeOf(m)
	fmt.Printf("%-10s%-25v%-25v\n", t.Kind(), t.Size(), t.Align())

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%-10s%-25v%-25v\n", "-", f.Name, f.Offset)
	}
}

func showStructAlign() {
	var t1 T1
	showStructAlignAndSize(t1)

	var t2 T2
	showStructAlignAndSize(t2)
}

/*
kind      reflect.TypeOf().Size()  reflect.TypeOf().Align() unsafe.Sizeof()     unsafe.Alignof()
bool      1                        1                        1                   1
int8      1                        1                        1                   1
uint8     1                        1                        1                   1
int16     2                        2                        2                   2
uint16    2                        2                        2                   2
int32     4                        4                        4                   4
uint32    4                        4                        4                   4
float32   4                        4                        4                   4
int64     8                        8                        8                   8
uint64    8                        8                        8                   8
float64   8                        8                        8                   8
uint      8                        8                        8                   8
int       8                        8                        8                   8
string    16                       8                        16                  8



*/
