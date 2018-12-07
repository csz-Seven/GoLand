package main

import (
	"fmt"
)

// 指针
func main() {
	var i int = 10
	fmt.Println("i地址为=", &i)

	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr指向的值=%v\n", *ptr)
}
