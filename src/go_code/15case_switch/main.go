package main

import "fmt"

func test(char int64) int64 {
	return char
}

var n1 int32 = 7
var n2 int64 = 77

func main() {
	// 1.case/switch 后是一个表达式( 即:常量值、变量、一个有返回值的函数等都可以)
	switch test(7) {
	case 7:
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	default:
		fmt.Println("default")
	}

	// 2.case 后的各个表达式的值的数据类型，必须和 switch 的表达式数据类型一致
	switch n2 {
	//错误示范
	//case n1:    ->n1为int32
	//	fmt.Println("A")
	case 'b':
		fmt.Println("A")
	case 77:
		fmt.Println("B-2")
	default:
		fmt.Println("default-2")
	}

	//3.case 后面可以带多个表达式，使用逗号间隔。比如 case 表达式 1, 表达式 2 ...
	switch 7 {
	// 随意达成一个
	case 10, 7, 6, 5, 4, 0:
		fmt.Println("实例-3")
	}

	// 4.case后面的表达式如果是常量值(字面量)，则要求不能重复
	switch 7 {
	// 随意达成一个
	case 10:
		fmt.Println("实例-4")
		// 错误示范
		//case 10:
		//	fmt.Println("实例-4")
	}

	//5.case 后面不需要带 break , 程序匹配到一个 case 后就会执行对应的代码块，然后退出 switch，如 果一个都匹配不到，则执行 default

	//6.switch 后也可以不带表达式，类似 if --else 分支来使用
	switch {
	case 1 > 2:
		fmt.Println("实例-6 1>2")
	case 1 > 0:
		fmt.Println("实例-6 1>0")
	}

	//  7.switch 后也可以直接声明/定义一个变量，！分号结束！，!!!不推荐!!!。
	switch n3 := 77; {
	case n3 > 88:
		fmt.Println("实例-7 n3>88")
	default:
		fmt.Println("实例-7 n3!>88")
	}

	// 8.switch 穿透-fallthrough ，如果在 case 语句块后增加 fallthrough(一次) ,则会继续执行下一个 case(不需要进行case的判断)，switch 穿透
	n4 := 77
	switch n4 {
	case 77:
		fmt.Println("实例-8 case 77")
		fallthrough
	case 66:
		fmt.Println("实例-8 fallthrough case 66")
	default:
		fmt.Println("实例-8 fallthrough default")
	}

	// 9.Type Switch:switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际指向的 变量类型
	var x interface{}
	var y = 10.0 // float
	x = y
	switch x.(type) {
	case float64:
		fmt.Println("x为float")
	}
}
// 88
