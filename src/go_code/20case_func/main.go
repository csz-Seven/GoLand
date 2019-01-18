package main

import "fmt"

func test6(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test6 n1= ", *n1)
}
func test8(n1 int) int {
	return n1
}

func test11(n1 int, n2 int) (r1 int, r2 int) {
	r1 = n1 + n2
	r2 = n1 - n2
	return
}

func test13(n1 int, args ... int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// 函数细节
func main() {
	// 1.函数形参和返回值 都可以是多个
	// 2.形参列表和返回值列表的数据类型可以说 值类型和引用类型

	// 3.函数的命名遵循标识符命名规范，首字母不能是数字，首字母大写该函数可以被本包文件和其
	//它包文件使用

	// 4.函数变量的局部性，函数外不生效 作用域

	// !5.基本数据类型和数组默认都是值类型传递，在函数内部修改，不会影响到外部.

	// !6.如果希望函数内部的修改 与外部变量同步，可以传入变量的地址，函数内以指针的方式操作变量.
	num6 := 20
	test6(&num6)
	fmt.Println("main() num6= ", num6)

	// 7.GO函数不支持 传统函数的重载
	// test07(n1 int){}   test07(n1 int,n2 int){}

	// 8.函数也是一种数据类型，可以赋值给变量，通过该变量可以对函数进行调用
	case8 := test8
	fmt.Printf("case8的类型%T ， test8的类型是%T", case8, test8)

	case8Res := case8(10)
	fmt.Println("case8Res=", case8Res)

	// 9.函数是数据类型，因此在Go中，函数可以作为形参，并调用

	// 10.为了简化数据类型定义，Go支持自定义数据类型.
	// 基本语法： type 自定义名 数据类型  ----->类型的别名
	type sevenInt int
	//虽然此时 sevenInt 和 int 都是int类型,但是go认为sevenInt和int是两个类型
	var case10Int int
	var case10SevenInt sevenInt
	case10SevenInt = 77
	//	此处需要进行类型转换 go认为sevenInt和int是不一样的类型
	case10Int = int(case10SevenInt)
	//case10Int = case10SevenInt
	fmt.Println("case10Int=", case10Int, "case10SevenInt=", case10SevenInt)

	// 11.函数的返回值支持命名
	// test11

	// 12.使用_ 标识符，可以忽略不需要的返回值.

	// 13.Go支持可变参数
	sum13 := test13(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("sum13=", sum13)
}
