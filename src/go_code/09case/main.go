package main

import "fmt"

//基本数据类型的转换 需要显示转换 （go不会自动转换）
func main() {
	var i int32 = 100;
	var n1 float32 = float32(i)
	var n2 int64 = int64(i) // 无论高->低 还是低->高 都需要显示转换
	fmt.Printf("i=%v n1=%v\n", i, n1)
	fmt.Printf("n2=%v", n2)

	// 被转换的是 变量存储的数据（值）,变量本身的数据类型没有变化.
}
