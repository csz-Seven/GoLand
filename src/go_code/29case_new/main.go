package main

import "fmt"

func main() {
	new1 := 1
	fmt.Printf("new1=>类型%T-值%v-地址%v\n", new1, new1, &new1)

	new2 := new(int) // *int 相当于指针
	*new2 = 2
	fmt.Printf("new2=>类型%T-地址%v-值%v(地址)指向的值%v\n", new2, &new2, &new2, *new2)
}
