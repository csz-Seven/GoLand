package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// -128  ~ 127
	var i int8 = -128
	var j int8 = 127
	fmt.Println(i, j)

	// 0~255
	var a uint8 = 0
	var b uint8 = 255
	fmt.Println(a, b)

	// int uint(32（4）/64（8）系统)
	// rune byte

	// golang区分符号和无符号
	// 声明整型默认为 int类型
	fmt.Printf("b 数据类型%T 字节数%d", b, unsafe.Sizeof(b))
}
// 38
