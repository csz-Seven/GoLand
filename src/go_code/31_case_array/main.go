package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 数组-可以存放多个 ！同一类型数据！。数组的数据类型为：值类型，与其他语言不一样，其他语言大多为引用类型.
	// 1.定义数组
	var array1 [7]float64 // 此时定义完成 各下标就已经存在零值了。
	array1[0] = 7.0
	array1[1] = 7.1
	array1[2] = 7.2
	array1[3] = 7.3
	array1[4] = 7.4
	array1[5] = 7.5
	array1[6] = 7.6
	// 错误示范 数组明确 同一类型的存放，该值为int 7
	//array1[6] = 7
	total1 := 0.0
	arraylen1 := len(array1)
	for i := 0; i < arraylen1; i++ {
		total1 += array1[i]
	}
	fmt.Printf("array1=%v \ntotal1=%v\n", array1, total1)

	// 2.
	// 数组的地址可以通过数组名获取 &array2
	// 数组的地址 其实等于 数组第一个元素下标的地址
	// 数组各个元素之间的地址 可以通过不同的数据类型 所占用的大小 递加或减 得出(16进制运算) 如int64-8 int32-4
	var array2 [3]int //此时int所占的是8个字符
	fmt.Printf("初始化array2=%v\n", array2)
	array2[1] = 1
	array2[2] = 2
	fmt.Printf("赋值后array2=%v\n", array2)
	fmt.Printf("array2的地址=%p, array2[0]的地址=%p\n", &array2, &array2[0])

	//3.四种初始化数组的方式
	// 可以指定下标的方式初始化值
	var init1 [2]int = [2]int{1: 1, 0: 2}
	var init2 = [2]int{1, 2}
	var init3 = [...]int{0, 1, 2}
	init4 := [...]string{"a", "b", "c"}
	fmt.Printf("init1=%v\n", init1)
	fmt.Printf("init2=%v\n", init2)
	fmt.Printf("init3=%v\n", init3)
	fmt.Printf("init4=%v\n", init4)

	//4. 数组遍历
	// for常规遍历

	// for-range遍历
	// 语法:for index,value := range array4{}
	// 说明 index(tip:当前下标) value(tip:当前下标对应的值)为自定义的变量名，不需要时_可以通过下划线忽略.
	arrary4 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	for index, value := range arrary4 {
		fmt.Printf("arrary4[%v]=%v\n\n", index, value)
	}

	//5.数组 使用注意和细节
	// (1.数组为多个相同类型的集合，其初始化的长度后无法动态增加数组的长度(例如JavaScript的Array为动态长度),定义长度后就已经固定.
	// (2.var array []int是slice切片，需要注意不要写错.
	// (3.数组为值类型， 数组中的元素为任意相同类型.
	// (4.数组在定义时，即便没有赋值操作，也已经存在默认值(根据不同类型 生成 零值)
	//     数值类型 0
	// 	   字符类型 “”
	//     bool类型 false
	// (5.Go语言数组 默认情况下为 值类型，修改数组不会互相影响,如果想修改原数组的值 可以通过指针的方式进行修改.

	// 随机生成多个数字，并进行反转输出.
	var array6 [6]int
	len := len(array6)
	// Seed使用给定的seed来初始化生成器到一个确定的状态。
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		// Intn返回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic。
		array6[i] = rand.Intn(77) // [0,77)范围
	}
	fmt.Printf("未反转=>>> array6=%v\n", array6)

	for i := 0; i < len/2; i++ {
		temp := array6[len-(1+i)]
		array6[len-(1+i)] = array6[i]
		array6[i] = temp
	}
	fmt.Printf("已反转=>>> array6=%v", array6)
}
