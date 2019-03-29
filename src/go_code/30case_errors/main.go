package main

import (
	"errors"
)

// 与其他语言不同 不存在 try catch 这类处理
// 替代的是 defer panic recover
// panic 抛出异常错误
// defer -> recover 用于捕获错误，然后进一步处理.
// errors.New(返回error类型)自定义错误 + panic

func readConf(name string) (error) {
	if name == "config" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func main() {
	// defer + recover
	//defer func() {
	//	err := recover() // 接收错误信息
	//	if err != nil {
	//		fmt.Println("recover捕获错误err=>", err)
	//	}
	//}()
	//
	//num1 := 7
	//num2 := 0
	//fmt.Println(num1 / num2)

	//errors.New + panic 自定义错误 并 抛出
	err2 := readConf("777")
	if err2 != nil {
		panic(err2)
	}
}
