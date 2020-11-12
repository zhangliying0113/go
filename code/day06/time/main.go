package main

import (
	"fmt"
	"time"
)

func f1() {
	now := time.Now()
	fmt.Println(now)           // 2020-11-09 10:53:49.3068111 +0800 CST m=+0.002018401 时区 CST：中国标准时间
	fmt.Println(now.Year())    // 2020
	fmt.Println(now.YearDay()) // 314
	fmt.Println(now.Month())   // November
	fmt.Println(now.Day())     // 9
	fmt.Println(now.Date())    // 2020 November 9
	fmt.Println(now.Hour())    // 10 当前时间的小时
	fmt.Println(now.Minute())  // 53 当前时间的分钟
	fmt.Println(now.Second())  // 当前时间的秒

	// 时间戳
	fmt.Println(now.Unix())     // 1604890429 时间戳毫秒
	fmt.Println(now.UnixNano()) // 1604890429306811100 时间戳纳秒

	// 根据时间戳推算时间
	ret := time.Unix(1604886600, 0)
	fmt.Println(ret)        // 2020-11-09 09:50:00 +0800 CST
	fmt.Println(ret.Year()) // 2020
	fmt.Println(ret.Day())  // 9

	// 时间间隔
	fmt.Println(time.Second) // 1s
	// 两个时间相加
	fmt.Println(now.Add(24 * time.Hour)) // 2020-11-10 10:53:49.3068111 +0800 CST m=+86400.002018401
	// 两个时间相减
	nextDay, err := time.Parse("2006-01-02 15:04:05", "2020-11-10 10:05:00")
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}
	fmt.Println(nextDay) // 2020-11-10 10:05:00 +0000 UTC 时间标准输出
	now = now.UTC()
	fmt.Println(now) // 2020-11-09 02:53:49.3068111 +0000 UTC
	d := nextDay.Sub(now)
	fmt.Println(d) // 31h11m10.6931889s   time.Parse 得到的时区为 UTC

	// 定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t) // 1 秒执行一次
	// }

	// 格式化时间 把语言中时间对象转化成字符串类型的时间
	fmt.Println(now.Format("2006-01-02"))              // 2020-11-09
	fmt.Println(now.Format("2006-01-02 15:04:05"))     // 2020-11-09 02:53:49
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))  // 2020-11-09 02:53:49 AM
	fmt.Println(now.Format("2006-01-02 15:04:05:000")) // 2020-11-09 03:01:42:000

	// 按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2020-11-09")
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)        // 2020-11-09 00:00:00 +0000 UTC
	fmt.Println(timeObj.Unix()) // 1604880000

	// sleep 5秒
	n := 5
	fmt.Println("start sleep")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("have sleep 5 sec")
}

// 解决 UTC 时区问题
func f2() {
	now := time.Now() // 本地的时间
	fmt.Println(now)  // 2020-11-09 10:51:20.1251041 +0800 CST m=+0.002000501
	// 明天的这个时间
	// 按照指定格式取解析一个字符串格式的时间
	// time.Parse("2006-01-02 15:04:05", "2019-08-04 14:41:50")
	// 按照东八区的时区和格式取解析一个字符串格式的时间
	// 根据字符串加载时区
	local, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Printf("load local failed,err:%v\n", err)
		return
	}

	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-11-10 10:48:00", local)
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	// 时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}

func main() {
	f1()
	// f2()
}
