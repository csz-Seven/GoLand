package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string 转基本数据类型
	var str, str2, str3 string = "true", "123456", "7.77777"
	var b bool
	var i int64
	var f float64

	b, _ = strconv.ParseBool(str)
	i, _ = strconv.ParseInt(str2, 10, 64)
	f, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("b type %T, b=%v\n", b, b)
	fmt.Printf("i type %T, i=%v\n", i, i)
	fmt.Printf("f type %T, f=%v\n", f, f)
}
// 51
