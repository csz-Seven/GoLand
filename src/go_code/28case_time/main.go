package main

import (
	"fmt"
	"time"
)

// packge time 提供时间显示和测量用的函数。日历为公历
func main() {
	//1. time.Time为一个类型，标示时间
	time1 := time.Now()
	fmt.Printf("当前时间为%v 类型为%T\n", time1, time1)

	// 2.获取详细的 时间单位
	fmt.Printf("年=%v\n", time1.Year())
	// 这种Time.Month() 为英文月份 ，如果需要数字表示月份 可以使用int转换
	fmt.Printf("En-月=%v\n", time1.Month())
	fmt.Printf("Num-月=%v\n", int(time1.Month()))
	fmt.Printf("日=%v\n", time1.Day())

	fmt.Printf("时=%v\n", time1.Hour())
	fmt.Printf("分=%v\n", time1.Minute())
	fmt.Printf("秒=%v\n", time1.Second())

	// 3.格式化日期
	// 3.1为使用Printf 和 Sprintf拼接想要的时间格式
	fmt.Printf("当前日期:%d-%d-%d %d:%d:%d\n",
		time1.Year(),
		time1.Month(),
		time1.Day(),
		time1.Hour(),
		time1.Minute(),
		time1.Second())

	timeFormat := fmt.Sprintf("当前日期:%d-%d-%d %d:%d:%d\n",
		time1.Year(),
		time1.Month(),
		time1.Day(),
		time1.Hour(),
		time1.Minute(),
		time1.Second())
	fmt.Println(timeFormat)
	// 3.2使用time.Format() 进行格式化  备注：！！！2006-01-02 15:04:05！！！为固定格式 Go诞生的寓意
	fmt.Printf(time1.Format("2006/01/02 15:04:05\n"))
	fmt.Printf(time1.Format("2006-01-02 15=04=05\n"))
	fmt.Printf(time1.Format("15:04:05\n"))

	// 4.time 时间常量 备注：常用的时间段。没有定义一天或超过一天的单元，以避免夏时制的时区切换的混乱。
	/**
	const (
    Nanosecond  Duration = 1  纳秒
    Microsecond          = 1000 * Nanosecond 微秒
    Millisecond          = 1000 * Microsecond 毫秒
    Second               = 1000 * Millisecond 秒
    Minute               = 60 * Second 分
    Hour                 = 60 * Minute 时
	)
	*/

	// 5.func Sleep(d Duration)
	// Sleep阻塞当前go程至少d代表的时间段。d<=0时，Sleep会立刻返回。
	for i := 0; i <= 7; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}

	// 6.从时间点January 1, 1970 UTC到时间点t所经过的时间
	// Unix时间戳(秒)
	// UnixNano(纳秒) 如果纳秒为单位的unix时间超出了int64能表示的范围，结果是未定义的。注意这就意味着Time零值调用UnixNano方法的话，结果是未定义的。
	fmt.Printf("Unix时间戳(秒)=%v \n Unix时间戳(纳秒)=%v", time1.Unix(), time1.UnixNano())
}
