package main

import "fmt"

func main() {
	var price float32 = 89.12
	fmt.Println(price)
	// 浮点型类型
	// 单精度 float32（4）
	// 双精度 float64（8）
	//浮点数=符号位+指数位+尾数位

	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	fmt.Println(num1, num2)

	var num3 = 3.1415926
	fmt.Printf("num3数据类型 %T", num3)

	num4 := 7.77777e2
	num5 := 7.77777E2
	fmt.Println(num4, num5)
}
