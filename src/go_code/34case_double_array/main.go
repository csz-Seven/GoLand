package main

import "fmt"

func main() {
	// 二维数组

	// 1.使用方式
	// a-先声明/定义，在赋值.
	var array1 [2][3]int
	array1[1][1] = 10
	fmt.Println(array1)
	// b-声明并赋值(初始化)
	var array2 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//var array2 [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(array2)

	// 2.二维数组遍历
	// a-for嵌套循环
	// b-for-range
	var array3 = [2][2]int{{1, 2}, {3, 4}}
	for index, value := range array3 {
		for indexChild, valueChild := range value {
			fmt.Printf("array3[%v][%v]=%v \n", index, indexChild, valueChild)
		}
	}
}
