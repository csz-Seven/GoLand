package main

import "fmt"

//全局变量声明
var (
	a1 = 1
	a2 = 2
	a3 = 3
)

func main() {
	// 默认值->0
	var i int
	fmt.Println(i)
	i = 10
	fmt.Println(i)

	// 变量类型推导 自动判定类型
	var num = 10.10
	fmt.Println(num)

	// 省略var :=左侧不能为已声明的变量
	name := "seven"
	fmt.Println(name)

	// 多变量声明
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	var n4, n5, n6 = 1, "seven", 7
	fmt.Println(n4, n5, n6)

	n7, n8, n9 := 2, "seven-2", 14
	fmt.Println(n7, n8, n9)

	fmt.Println(a1, a2, a3)
}
