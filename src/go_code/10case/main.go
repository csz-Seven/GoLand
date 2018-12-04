package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num2 float64 = 3.1415926
	var b bool = true
	var myChar byte = 'l'
	var str string

	//基本数据类型转string
	// 1.fmt.Sprintf 转换
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T, str = %q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T, str = %q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T, str = %q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T, str = %q\n", str, str)
	fmt.Println("---------------------")

	// strconv 转换
	var num11 int = 99
	var num12 float64 = 3.1415926
	var b2 bool = true

	str = strconv.FormatInt(int64(num11), 10)
	fmt.Printf("str type %T,str = %q\n", str, str)

	str = strconv.FormatFloat(num12, 'f', 10, 64)
	fmt.Printf("str type %T,str = %q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T,str = %q\n", str, str)
}
