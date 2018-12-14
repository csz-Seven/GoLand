package main

import "fmt"

func main() {
	// break 跳出for循环 或 switch
	// 在for使用环境中 break可以指定label标签(名字随意)，从而跳出某层for循环
	for i := 0; i < 7; i++ {
		for j := 0; j < 10; j++ {
			if j == 1 {
				break
			}
			fmt.Println("j=", j)
		}
	}

	fmt.Println("------------------------------------------")

label1: //此处为label 名字自定义
	for i := 0; i < 7; i++ {
		for j := 0; j < 10; j++ {
			if j == 1 {
				break label1 //break可以通过label指定跳出for层
			}
			fmt.Println("j=", j)
		}
	}
}
