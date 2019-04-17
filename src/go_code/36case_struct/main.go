package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// struct 结构体

	// GoLand 面向对象
	// 1.面向对象OOP为一门思想，各编程语言对思想的实现各有不同.
	// 2.Go与传统的面向对象OOP语言有区别(如：java),支持面向对象编程特效.
	// 3.Go中没有类(class)，取而代之的是结构体(struct)来实现面向对象OOP特性.
	// 4.GO中没有传统语言的继承、重载、构造函数、析构函数、this等.
	// 5.Go有实现OOP 继承、封装、多态的特性.

	// 1.struct 快速演示
	// 参考对象：people
	// 结构体 声明结构体 如下：
	// type 结构体名称 struct{
	// 属性名1 type
	// 属性名2 type
	// }
	type People struct {
		Name  string
		Age   int
		Race  string
		Hobby string
	}
	// 结构体变量 如下：
	var people1 People
	people1.Name = "蔡徐坤"
	people1.Age = 20
	people1.Race = "yellow"
	people1.Hobby = "鸡你太美"
	fmt.Println(people1)

	// 2.结构体与 结构体变量的关系：结构体为自定义数据类型，泛指一类；结构体变量为产生的实体对象，更具体.

	// 3.创建结构体变量后，没有进行属性赋值，各类型属性为预设一个零值
	// 入 bool-false int-0 string-"" clice、map-nil(需要make分配空间)等
	type Preinstall struct {
		Name   string
		Age    int
		Scores [5]float64
		ptr    *int
		slice  []int
		Map    map[string]string
	}
	var preinstall1 Preinstall
	fmt.Println(preinstall1)

	// 4.结构体为值类型，不同结构体变量独立互不影响.
	// 值拷贝people1
	people2 := people1
	people2.Name = "caixukun"
	fmt.Println(people2)

	// 5.创建结构体变量的方式：
	// 结构体继续使用People
	// 5.1 直接声明 var people Person

	// 5.2 people := Person{"蔡徐坤",20,"yellow","鸡你太美"}  顺序参照结构体属性顺序

	// 5.3
	var people3 = new(People)
	//var people3 *People = new(People)
	// 未简化的写法
	(*people3).Age = 20
	// Go简化后的语法
	people3.Name = "某名"
	fmt.Println(*people3)

	// 5.4
	var people4 = &People{}
	//var people3 *People = new(People)
	// 未简化的写法
	(*people4).Age = 24
	// Go简化后的语法 内部调用(*people4).Age
	people4.Name = "某名4"
	fmt.Println(*people4)

	// 6.结构体使用注意事项和细节
	// 6.1结构体所有属性(字段)在内存中是 连续的关系: N+1个属性=N内存地址 + type所占的字节

	// 6.2结构体是独立的，与其他类型进行转换时需要保证 属性名、属性个数、属性类型 全部相同!(重点)
	type A struct {
		Name int
	}
	type B struct {
		Name int
	}
	var a A
	var b B
	//需要强制换换 而不是a = b
	a = A(b)
	fmt.Println(a)

	// 6.3struct 有一个 tag的用法，可以通过反射机制获取，常见使用场景 JSON序列化和反序列化.
	type TagCase struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var tagCase1 TagCase
	tagCase1.Name = "韩梅梅"
	tagCase1.Age = 60

	fmt.Println()
	// 引入encoding/json包
	// json.Marshal返回两个值 byte err,根据需要自行取值.
	jsonTagCase1, _ := json.Marshal(tagCase1)
	fmt.Println("JSON:", jsonTagCase1)
	fmt.Println("JSON->bype进行string()转换 和 tag的使用:", string(jsonTagCase1))
}
