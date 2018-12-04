package main

import "fmt"

func main() {
	// string
	var address string = "深圳"
	fmt.Println(address)

	// 字符串两种表现形式
	//1.双引号 会识别转义字符
	//2.反引号 义字符串原生输出 可以实现防止攻击 输出源代码等效果
	var str string = `func main (){
	// string
	var address string ="深圳"
	fmt.Println(address)`
	fmt.Println(str)

	// 当拼接很长字符串时 加号需要注意（go默认最后面加；号）
	var str2 string = "1" +
		"1" +
		"1"
	str2 ="haha"
	fmt.Println(str2)
}
