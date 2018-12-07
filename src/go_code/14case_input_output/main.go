package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool

	//fmt.Println("姓名")
	//fmt.Scanln(&name)
	//
	//fmt.Println("年龄")
	//fmt.Scanln(&age)
	//
	//fmt.Println("薪水")
	//fmt.Scanln(&sal)
	//
	//fmt.Println("是否通过")
	//fmt.Scanln(&isPass)

	fmt.Println("姓名 年龄 薪水 是否通过")
	fmt.Scanf("%s %d %f %t",&name,&age,&sal,&isPass)

	fmt.Printf("姓名：%v\n年龄：%v\n薪水：%v\n是否通过：%v\n", name, age, sal, isPass)
}
// 68
