package main

import "fmt"

// 每一个源文件都可以包含一个 init 函数，该函数在main执行之前，被Go运行框架调用.
//init -> main

// 如果文件还包含 全局变量的定义
// 执行流程 全局变量定义->init ->main

// 如果存在包的引用，则根据 谁先被引入 谁就先进行 全局->init->main的流程,以此类推.
var caseInit = variable()

func variable() int {
	fmt.Println("全局变量")
	return 77
}
func init() {
	fmt.Println("init()执行")
}
func main() {
	fmt.Println("main()执行")
}
