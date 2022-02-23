package packageuse

import (
	"fmt"
	"strconv"
	"time"
)

/******************************** 获取时间 start *********************************/
// 获取当前时间
func GetTime1() {
	currentTime := time.Now()
	fmt.Printf("类型: %T 值: %+v\n", currentTime, currentTime)
	unix := time.Now().Unix()
	fmt.Printf("当前时间戳(单位秒): %v \n", unix)
	nano := time.Now().UnixNano()
	fmt.Printf("当前时间戳(单位纳秒): %v \n", nano)
	format := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("当前时间(Y-m-d H:i:s): %v \n", format)
	format2 := time.Now().Format("20060102150405")
	fmt.Printf("当前时间(YmdHis): %v \n", format2)
}

// 获取当前年、月、日、时、分、秒、星期几
func GetTime2() {
	format := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("当前时间(Y-m-d H:i:s): %v \n", format)
	fmt.Printf("当前年: %v \n", time.Now().Year())
	fmt.Printf("当前月: %v \n", time.Now().Month())
	fmt.Printf("当前日: %v \n", time.Now().Day())
	fmt.Printf("当前小时: %v \n", time.Now().Hour())
	fmt.Printf("当前分钟: %v \n", time.Now().Minute())
	fmt.Printf("当前秒: %v \n", time.Now().Second())
	fmt.Printf("当前星期几: %v \n", time.Now().Weekday())
}

// 获取时分秒
func GetTime3() {
	now := time.Now()
	hour, min, sec := now.Clock()
	fmt.Printf("时间: %v, 时: %v, 分: %v, 秒: %v\n", now, hour, min, sec)
}

// 创建指定时间
func GetTime4() {
	date := time.Date(2020, 8, 23, 13, 53, 53, 0, time.Local)
	fmt.Printf("类型: %T, 值: %+v\n", date, date)
}

// 获取当前时间是，今年的第几周
func GetTime5() {
	year, week := time.Now().ISOWeek()
	fmt.Printf("year: %v, week: %v\n", year, week)
}

// 获取今天是今年的第几天
func GetTime6() {
	num := time.Now().YearDay()
	fmt.Printf("获取今天是今年的第%v天\n", num)
}

/******************************** 获取时间 end *********************************/

/******************************** 字符串转时间 start *********************************/
// 字符串转时间类型
func Str2Date1() {
	str := "1616319808"
	unix, _ := strconv.ParseInt(str, 10, 64)
	format := time.Unix(unix, 0)
	fmt.Printf("字符串时间戳 --> 类型: %T, 值: %+v\n", format, format)

	// 字符串时间
	strDate := "2019-04-10 12:54:03"
	// 注意layout格式需要和字符串时间格式一致
	location, _ := time.Parse("2006-01-02 15:04:05", strDate)
	fmt.Printf("字符串时间 --> 类型: %T 值: %v \n", location, location)
}

func Str2Date2() {
	// 将字符时间: 2020-09-12 14:34:10 转成 20200912143410
	strDate := "2020-09-12 14:34:10"
	// 1.先转成时间类型
	location, _ := time.ParseInLocation("2006-01-02 15:04:05", strDate, time.Local)
	// 2.再格式化
	format := location.Format("20060102150405")
	fmt.Printf("类型:%T, 值: %+v\n", format, format)
}

/******************************** 字符串转时间 end *********************************/

/******************************** 时间比较 start *********************************/
func CompareDate() {
	strDate1 := "2022-02-22 21:25:00"
	strDate2 := "2022-02-22 21:25:00"
	// 1.先转成时间类型
	time1, _ := time.ParseInLocation("2006-01-2 15:04:05", strDate1, time.Local)
	time2, _ := time.ParseInLocation("2006-01-2 15:04:05", strDate2, time.Local)
	// strDate2 > strDate1 ?
	before := time1.Before(time2)
	fmt.Printf("strDate2 > strDate1 ? %t\n", before)
	// strDate2 > strDate1 ?
	equal := time1.Equal(time2)
	fmt.Printf("strDate2 = strDate1 ? %t\n", equal)
}

/******************************** 时间比较 end *********************************/

/******************************** 时间计算 start *********************************/
// 时间相加
func CalculateDate1() {
	now := time.Now()
	fmt.Printf("现在的时间: %v\n", now)
	// 十分钟前
	duration, _ := time.ParseDuration("-10m")
	fmt.Printf("十分钟前: %v\n", now.Add(duration))
	// 一小时前
	duration2, _ := time.ParseDuration("-1h")
	fmt.Printf("一小时前: %v\n", now.Add(duration2))
	// 一天后、一月后、一年后
	fmt.Printf("一天后: %v\n", now.AddDate(0, 0, 1)) // 一天后
	fmt.Printf("一月后: %v\n", now.AddDate(0, 1, 0)) // 一月后
	fmt.Printf("一年后: %v\n", now.AddDate(1, 0, 0)) // 一年后
}

// 时间相减
func CalculateDate2() {
	day1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-02-22 12:00:00", time.Local)
	day2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-02-22 13:00:00", time.Local)
	fmt.Printf("相差几秒: %v\n", day2.Sub(day1).Seconds())
	fmt.Printf("相差多少分钟: %v\n", day2.Sub(day1).Minutes())
	fmt.Printf("相差多少小时: %v\n", day2.Sub(day1).Hours())
}

/******************************** 时间计算 end *********************************/

/******************************** 定时器 start *********************************/
//NewTicker
func Ticker1() {
	// 创建定时器，间隔设置每秒
	ticker := time.NewTicker(time.Second)
	// 启动一个协程，打印定时器里面的时间
	go func(ticker *time.Ticker) {
		for i := 0; i < 3; i++ {
			fmt.Println(<-ticker.C)
		}
		// 关闭定时器
		ticker.Stop()
	}(ticker)
	// 手动阻塞
	time.Sleep(3 * time.Second)
	fmt.Println("end")
}

// Tick
func Ticker2() {
	// 创建定时器，间隔设置为每秒
	chTime := time.Tick(time.Second)
	// 启动一个协程，打印定时器里面的时间
	go func(ch <-chan time.Time) {
		for i := 0; i < 3; i++ {
			fmt.Println(<-ch)
		}
	}(chTime)
	// 手动阻塞
	time.Sleep(4 * time.Second)
	fmt.Printf("end")
}

/******************************** 定时器 end *********************************/

/******************************** 延迟执行 start *********************************/
func DelayExec1() {
	fmt.Printf("开始时间: %v\n", time.Now())
	timer := time.NewTimer(3 * time.Second)
	// 此处会阻塞，知道timer.C中有数据写入
	fmt.Printf("timer通道里的时间： %v\n", <-timer.C)
}

func DelayExec2() {
	// 创建一个计时器，返回的是chan
	ch := time.After(5 * time.Second)
	fmt.Printf("开始时间 %v \n", time.Now())
	// 此处会阻塞5秒
	out := <-ch
	fmt.Printf("变量out->  类型: %T 值:%v  \n", out, out)
	fmt.Printf("结束时间 %v \n", time.Now())
}

/******************************** 延迟执行 end *********************************/
