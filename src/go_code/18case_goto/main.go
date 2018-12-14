package main

import "fmt"

func main() {
	fmt.Println("1")
	goto label1
	// 此处会被跳过
	fmt.Println("2")

label1:
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
}
