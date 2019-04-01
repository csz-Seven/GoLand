package main

import "fmt"

func main() {
	// slice切片
	// 使用和数组类似
	// 切片是 数组的一个引用，切片为引用类型.
	// 切片长度是动态变化的，可以把它比作 动态变化长度的数组,其数据结构是struct结构体.

	// 1.基本使用
	intArr := [...]int{1, 2, 3, 4, 5, 6, 7}
	// 声明\定义切片
	slice1 := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice1元素=", slice1)
	fmt.Println("slice1元素个数=", len(slice1))
	fmt.Println("slice1的动态容量=", cap(slice1))
	// struct结构体
	//type slice struct {
	//	ptr *[2]int
	//  len int
	//  cap int
	//}

	// 2.使用方式

	// 2.1 定义一个切片，然后让切片去引用已经创建好的 数组.
	// 1.基本使用的案例 就是这种方式.

	// 2.2 通过make来直接创建切片，省去手动引用创建数组的步骤.
	// make([]type,len,[cap])
	// type类型 len长度 cap容量(可选，前置条件cap>=len)
	// 各元素在未赋值的情况下，存在默认值.
	// make其实就是 自动创建了数组 并且引用，但是这种数组是对外不可见的，只能通过slice访问.
	slice2 := make([]float64, 7, 10)
	slice2[1] = 1
	slice2[2] = 2
	slice2[6] = 6
	fmt.Println("slice2元素=", slice2)
	fmt.Println("slice2元素个数=", len(slice2))
	fmt.Println("slice2的动态容量=", cap(slice2))

	// 2.3 定义切片，并且直接指定具体的数组,类似于2.2中的make创建.
	slice3 := []string{"1", "2", "3"}
	fmt.Println("slice2元素=", slice3)
	fmt.Println("slice2元素个数=", len(slice3))
	fmt.Println("slice2的动态容量=", cap(slice3))

	// 3.切片遍历方式
	// 3.1for 循环常规遍历
	// 3.2for-range
	fmt.Printf("\nfor-range循环\n")
	for index, value := range slice1 {
		fmt.Printf("index=%v,value=%v\n", index, value)
	}

	// 4.使用注意事项和细节

	// 1= 切片初始化 var slice = arr[startIndex:endIndex] 范围：[startIndex,endIndex) 不包括endIndex

	// 2= 切片虽然是动态增长的 ，但是在初始化过程中，是不可以越界的,范围在[0,len(arr)）之间
	//    简写： var slice = arr[0,endI] ---> var slice = arr[:endI]
	//    简写： var slice = arr[startI,len(arr)] ---> var slice = arr[startI:]
	//    简写： var slice = arr[0,len(arr)] ---> var slice = arr[:]

	// 3= cap是内置函数，用于统计切片的容量，即最大可以存放个数.

	// 4= 切片可以连续切片 A->B->C
	fmt.Printf("\nA、B连续切片\n")
	sliceA := []string{"连", "续", "切", "片"}
	sliceB := sliceA[2:]
	fmt.Println("连续切片A=", sliceA)
	fmt.Println("连续切片B=", sliceB)
	sliceB[0] = "切切"
	fmt.Println("sliceB[0]= '切切' 连续切片A=", sliceA)
	fmt.Println("sliceB[0]= '切切' 连续切片B=", sliceB)

	// 5= append内置函数，为切片动态增加元素.
	appendCase := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("appendCase=", appendCase)
	appendCase = append(appendCase, 7, 8, 9, 10)
	fmt.Println("appendCase追加=", appendCase)
	// append原理分析
	// 本质是切片对引用数组的扩容操作，
	// 因为数组本身长度是不可变的，所以在Go底层在扩容中创建一个新的数组arr，
	// 先是将原先slice的值拷贝新数组arr中，再将追加的拷贝进去，
	// 最后slice重新引用新的arr.

	// 6= copy内置函数，不存在引用类型的影响 A copy B
	copyA := make([]int, 7)
	copyB := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("copyA=", copyA)
	fmt.Println("copyB=", copyB)
	copy(copyA, copyB)
	fmt.Println("copyA copy()后=", copyA)
	// 在copy之后 修改B或A的值都不会产生影响.

	// 5.string 和 slice
	// string 底层是byte数组，因此string是可以进行切片slice处理的
	// string是不可变的，所以不能通过 数组下标方式str[xIndex]修改
	// 如果需要改变字符串 1.重新给字符串赋值 2.通过string->[]byte 、 string->rune,最后转换为string.
	str := "csz.seven@gmail.com"
	strSlice := str[4:]
	println("strSlice切片引入str byte数组", strSlice)
	// 错误示范
	//strSlice[0] = "s"

	// 正确示范(非中文字符)
	strByte := []byte(str)
	strByte[0] = 'g'
	str = string(strByte)
	println("str 通过[]byte() 修改=", str)

	// 强类型语言 ->注意类型
	strRune := []rune(str)
	strRune[0] = '广'
	str = string(strRune)
	println("str 通过[]rune() 修改=", str)
}
