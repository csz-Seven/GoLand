package main

import "fmt"

func bubbleSort(arr *[5]int) {
	//{66, 55, 77, 99, 11}

	//{55, 66, 77, 99, 11}
	//{55, 66, 77, 99, 11}
	//{55, 66, 77, 99, 11}
	//{55, 66, 77, 11, 99}

	fmt.Println(len(*arr))
	temp := 0
	//外层n-1
	for j := 0; j < len(*arr)-1; j++ {
		for i := 0; i < len(*arr)-1-j; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				temp = (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = temp
			}
		}
	}
}
func main() {
	arr := [5]int{66, 55, 77, 99, 11}
	bubbleSort(&arr)
	fmt.Println(arr)
}
