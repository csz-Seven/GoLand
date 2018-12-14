package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 1.
	for i := 0; i < 10; i++ {
		fmt.Println("7", i)
	}

	//使用方式2
	i := 0
	for i < 10 {
		fmt.Println("7~", i)
		i++
	}

	//使用方式3
	ii := 0
	for {
		if ii < 10 {
			fmt.Println("7*", ii)
		} else {
			break
		}
		ii++
	}

	// 字符串遍历1(不适用于中文等字符,原因是for是按1字节逐个遍历。)
	var str string = "seven"
	for i := 0; i < len(str); i++ {
		//fmt.Println("\n", str[i])  此处输出的是as码
		fmt.Printf("%c \n", str[i])
	}
	// 字符串遍历2 for-range(适用于中文等字符)
	var strC string = "七七七"
	for index, val := range strC {
		fmt.Printf("index=%d  val=%c\n", index, val)
	}

	// 随机生成
	for{
		//rand.Seed(time.Now().Unix())
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println(n)
		if n == 99 {
			break
		}
	}
	// 100
}
