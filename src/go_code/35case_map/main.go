package main

import "fmt"

func main() {
	// 1.map :key:value数据结构
	// 基本语法: var map 变量 map[keyType]valueType
	var map1 map[string]string
	// 所有的map在使用之前都需要make分配数据空间
	map1 = make(map[string]string, 10)
	map1["k1"] = "蔡徐坤"
	map1["k2"] = "唱"
	map1["k3"] = "跳"
	map1["k4"] = "Rap"
	map1["k5"] = "篮球"
	fmt.Println(map1)

	// 2.备注
	// a. map使用之前必须map
	// a. map中的key值如果重复出现，则前者key会被后者key覆盖
	// a. map中key：value结构为无序的

	// 3.使用方式
	// a.1.已介绍
	// b.
	map2 := make(map[string]string, 7)
	map2["no.1"] = "个人练习生"
	map2["no.2"] = "在美国校队打球"
	fmt.Println(map2)

	// c.
	map3 := map[string]string{
		"no.1": "蔡",
		"no.2": "徐",
		"no.3": "坤",
		"no.4": "真",
		"no.5": "蔡",
	}
	fmt.Println(map3)

	// 4.map删除  delete(map,"key")内置函数,key存在删除，如果不存在key不报错.
	delete(map3, "no.4")
	delete(map3, "no.5")
	fmt.Println(map3)

	// 5.判断是否存在key
	map3No4, exist := map3["no.4"]
	if exist {
		fmt.Println("存在", map3No4)
	} else {
		fmt.Println("map3No4不存在")
	}

	// 6.map遍历
	// 1.for(！无法实现遍历，因为map为无序的！)
	// 2.for-range
	array6 := map[string]string{
		"遍历1": "1值",
		"遍历2": "2值",
		"遍历3": "3值",
	}
	for key, value := range array6 {
		fmt.Printf("key=%v value=%v \n", key, value)
	}

	// 7.map长度 使用 func len内置函数
	fmt.Printf("array6的长度=%v \n", len(array6))

	// 8.map切片
	// 切片的数据类型为map，则map的个数就为动态变化的
	// 切片和map都需要map！
	mapSlice1 := make([]map[string]string, 2)

	mapSlice1[0] = make(map[string]string, 2)
	mapSlice1[0]["no.1"] = "鸡"
	mapSlice1[0]["no.2"] = "你"

	mapSlice1[1] = make(map[string]string, 2)
	mapSlice1[1]["no.1"] = "太"
	mapSlice1[1]["no.2"] = "美"

	fmt.Println(mapSlice1)

	// 增加切片
	mapSlice1 = append(mapSlice1, map[string]string{
		"no.1": "Boom",
		"no.2": "Boom!",
	})
	fmt.Println(mapSlice1)


	// 9.使用细节
	// 9.1map为引用类型传递
	// 9.2map为的长度为动态变化
}
