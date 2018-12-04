package main

import (
	"fmt"
	"unsafe"
)

// bool
func main() {
	//	占用一字节
	//  只允许true false
	var b = false
	fmt.Println(b, unsafe.Sizeof(b))
}
