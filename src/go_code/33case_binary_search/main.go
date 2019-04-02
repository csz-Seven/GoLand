package main

import "fmt"

func binarySearch(arr [7]int, findValue int, firstIndex int, lastIndex int) {
	// 二分查找前置条件：有序序列

	//fmt.Println(arr)
	// firstIndex > lastIndex   为 未找到查询结束
	if firstIndex > lastIndex {
		fmt.Println("未能找到对应的Index")
	}

	middleIndex := (firstIndex + lastIndex) / 2
	// arr[middle] > findValue  为 [firstIndex,middleIndex)
	if arr[middleIndex] > findValue {
		binarySearch(arr, findValue, firstIndex, middleIndex-1)
	} else if arr[middleIndex] < findValue {
		binarySearch(arr, findValue, middleIndex+1, lastIndex)
	} else {
		fmt.Println("已找到Index=", middleIndex)
	}
	// arr[middle] == findValue 为 查询成立.
}
func main() {
	dictionary := [7]int{0, 1, 2, 3, 4, 5, 6}
	binarySearch(dictionary, 5, 0, len(dictionary)-1)
}
