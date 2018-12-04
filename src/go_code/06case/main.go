package main

import "fmt"

// go的字符串，由字节组成
// go没有专门的字符类型
func main() {
	var c1 byte = 'a'
	var c2 byte = '0'

	fmt.Println(c1, c2)
	fmt.Printf("%c %c\n", c1, c2)

	var c3 int = '测'
	fmt.Printf("%c\n", c3)

	var c4 int = 22269
	fmt.Printf("%c", c4)
}
