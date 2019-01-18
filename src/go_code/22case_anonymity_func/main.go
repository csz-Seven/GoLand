package main

import "fmt"

// 匿名函数
// 没有名字的函数，匿名函数可以一次性调用，或者多次调用.

// 全局匿名函数
var Fun1 = func(n1 int) int {
	return n1
}

func main() {
	// 一次性调用
	v1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(66, 11)
	fmt.Println("v1=", v1)

	// 多次调用
	// 把匿名函数赋值给一个变量，再通过函数变量多次调用匿名函数.
	f1 := func(n1 int) int {
		return n1
	}
	v2 := f1(37)
	fmt.Println("v2=", v2)
}
