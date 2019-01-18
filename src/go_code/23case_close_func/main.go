package main

import "fmt"

// 闭包
// 个人理解 函数与 相关引用环境 的整体 ，与作用域相关,防止全局变量污染，也可以进行一些函数的初始化.

func AddNumber() func(int) int {
	var n int = 7
	return func(num int) int {
		n = n + num
		return n
	}
}

func main() {
	// 初始化
	func1 := AddNumber();
	fmt.Println(func1(1)) // 8
	fmt.Println(func1(2)) // 10
	fmt.Println(func1(3)) // 13
}
