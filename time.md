# time
Go 语言使用 Location 来表示地区相关的时区，一个 Location 可能表示多个时区。
Time struct 是一个时间点，基于时区，Time 零值代表时间点 January 1, year 1, 00:00:00.000000000 UTC
## 时间转换
获取时区实例

指定格式字符串（必须`以2006-01-02 15:04:05`这个时间点准的格式）

ParseInLocation
## 定时器
定时器是进程规划自己在未来某一时刻接获通知的一种机制。
两种定时器：Timer（到达指定时间触发且只触发一次）和 Ticker（间隔特定时间触发）。
## code
```
package main

import (
	"time"
	"fmt"
)

// Go 语言使用 Location 来表示地区相关的时区，一个 Location 可能表示多个时区。
// Time struct 是一个时间点，基于时区
// Time 零值代表时间点 January 1, year 1, 00:00:00.000000000 UTC

func main () {

	fmt.Println(time.Now())
	// 函数 LoadLocation 可以根据名称获取特定时区的实例
	fmt.Println(time.LoadLocation(""))
	fmt.Println(time.LoadLocation("UTC"))
	fmt.Println(time.LoadLocation("Asia/Shanghai"))
	loc, _ := time.LoadLocation("Local")
	fmt.Println(loc)

	// time
	var sec, nsec int64
	//sec = 10
	fmt.Println(sec, nsec)
	t := time.Unix(sec, nsec)
	fmt.Println(t)
	fmt.Println(t.Unix())

	fmt.Println(t.IsZero())
	fmt.Println(t.UnixNano())

	pt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-12-26 22:00:00", loc)
	fmt.Println(pt)
	fmt.Println(pt.Unix())
	fmt.Println(pt.In(loc).Format("2006-01-02 15:04:05"))
	fmt.Println(pt.In(loc).Format(time.RFC3339))

	fmt.Println(t.String())

	// http://blog.csdn.net/qq_26981997/article/details/53454606
	pt, _ = time.ParseInLocation(t.String(), "2017-12-26 22:00:00", loc)
	fmt.Println(pt)

	local, _ := time.LoadLocation("America/Los_Angeles")
	timeFormat := "2006-01-02 15:04:05"

	// 通过unix标准时间的秒，纳秒设置时间
	time1 := time.Unix(1480390585, 0)

	// 洛杉矶时间
	time2, _ := time.ParseInLocation(timeFormat, "2016-11-28 19:36:25", local)
	fmt.Println(time1.In(local).Format(timeFormat))
	fmt.Println(time2.In(local).Format(timeFormat))

	// 运行时，该服务器必须设置为中国时区，否则最好是采用"Asia/Chongqing"之类具体的参数。
	chinaLocal, _ := time.LoadLocation("Local")
	fmt.Println(time2.In(chinaLocal).Format(timeFormat))

	// output:
	// 2016-11-28 19:36:25
	// 2016-11-28 19:36:25
	// 2016-11-29 11:36:25
	// 中国比洛杉矶快16小时
}

```
