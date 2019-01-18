package main

import "fmt"

// defer (函数中),一种延时机制，类似于完成某事后再调用的回调函数.

// defer执行时机，在函数执行完之后，如果存在多个defer语句，则按 先进后出(栈) 的方式执行.
// ps:在defer的进栈过程中，需要使用到的值都会进行一次拷贝(备份).
func main() {
	v1 := 7
	v2 := 77
	defer fmt.Println("第一个defer", v1)
	defer fmt.Println("第二个defer", v2)

	// ps:在defer的进栈过程中，需要使用到的值都会进行一次拷贝(备份).
	v1++
	v2++
	fmt.Println("函数执行语句->", v1, v2)

	// 常见使用场景
	// 创建资源如 链接数据库 打开文件等...
	// 交给defer来处理 断开数据库 关闭文件等操作
}
