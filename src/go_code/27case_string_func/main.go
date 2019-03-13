package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1.func len(v Type) int
	// 返回v的长度，取决于具体
	// 数组：v中的元素个数
	// 数组指针：*v中元素个数
	// 切片、隐射、v中的元素个数;nil为0
	// 字符串：v中字节个数
	// 通道： 通道缓存中的队列(未读取)元素数量
	str1 := "str1字符串"
	fmt.Println("str1 len=", str1)

	// 2.字符串遍历，同时处理中文拆分问题 r:=[]rune(str)
	str2 := "str2字符串"
	str2Fn := []rune(str2)
	for i := 0; i < len(str2Fn); i++ {
		fmt.Printf("字符串=%c.\n", str2Fn[i])
	}

	// 3.字符串转整数 func Atoi(s string) (i int, err error)
	// Atoi是ParseInt(s, 10, 0)的简写
	// package strconv
	str3, err3 := strconv.Atoi("str3字符串")
	if err3 != nil {
		fmt.Println("转换错误", err3)
	} else {
		fmt.Println("转换成功", str3)
	}

	// 4.整数转字符串 func Itoa(i int) string
	// Itoa是FormatInt(i, 10) 的简写。
	// package strconv
	str4 := strconv.Itoa(4)
	fmt.Printf("str=%v,str=%T", str4, str4)

	// 5.字符串转[]btye
	str5 := []byte("str5字符串")
	fmt.Printf("str5=%v\n", str5)

	// 6.[]byte转字符串
	str6 := string([]byte{115, 116, 114, 53, 229, 173, 151, 231, 172, 166, 228, 184, 178})
	fmt.Printf("str6=%v\n", str6)

	// 7.10进制转换 2 8 16进制 func FormatInt(i int64, base int) string
	str7 := strconv.FormatInt(7, 2)
	fmt.Printf("str7=7 二进制=%v\n", str7)

	// 8.查找字符 是否存在指定的字符串当中 func Contains(s, substr string) bool
	// package strings
	fmt.Printf("str8字符串 是否包含8=%v\n", strings.Contains("str8字符串", "8"))

	// 9.统计A字符串在B中出现的个数 func Count(s, sep string) int
	// package strings
	fmt.Printf("str999存在%v个9\n", strings.Count("str999", "9"))

	// 10.判断两个utf-8编码字符串（将unicode大写、小写、标题三种格式字符视为相同）是否相同
	// func EqualFold(s, t string) bool
	// package strings
	fmt.Printf("str10 EqualFold STR10 = %v\n", strings.EqualFold("str10", "STR10"))

	// 11.返回子串在 字符串 第一次出现的index值，如果没有返回-1
	// func Index(s, sep string) int
	// package strings
	fmt.Printf("11在str11首次出现的index值=%v\n", strings.Index("str11", "11"))

	// 12.返回子串在 字符串 最后一次出现的index值，如果没有返回-1
	fmt.Printf("12在12str12最后一次出现的index值=%v\n", strings.LastIndex("12str12", "12"))

	// 13.替换指定的A子串 为 B子串
	//  func Replace(s, old, new string, n int) string
	// n 为替换的次数
	fmt.Printf("str替换成string=%v\n", strings.Replace("13str13str", "str", "string", -1))

	// 14. 按照 子串 进行字符串分割
	// func Split(s, sep string) []string
	fmt.Printf("'str,14,字符串' = %v\n", strings.Split("str,14,字符串", ","))

	// 15.字符串大小写转换
	str15 := "str STRING"
	fmt.Printf("大写=%v 小写=%v\n", strings.ToUpper(str15), strings.ToLower(str15))

	// 16. 去除 字符串 两边的空格 func TrimSpace(s string) string
	fmt.Printf(" str16 =%q\n", strings.TrimSpace(" str16 "))

	// 17. 去除 字符串 两边指定的子串 func Trim(s string, cutset string) string
	fmt.Printf("'! str17! 'Trim=>%v\n", strings.Trim("! str17! ", " !"))

	// 18.去除 字符串 左边/右边 指定子串
	// func TrimLeft(s string, cutset string) string
	// func TrimRight(s string, cutset string) string
	fmt.Printf("'!str18!' left=>%v right=>%v\n", strings.TrimLeft("!str18!", "!"), strings.TrimRight("!str18!", "!"))

	// 19.判断字符串 是不是该子串 开头/末尾
	// func HasPrefix(s, prefix string) bool
	// func HasSuffix(s, suffix string) bool
	str19 := "str19字符串"
	fmt.Printf("str19字符串 prefix==str %v,suffix==字符串 %v,", strings.HasPrefix(str19, "str"), strings.HasSuffix(str19, "字符串"))
}
