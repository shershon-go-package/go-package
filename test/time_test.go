package test

import (
	"fmt"
	"shershon1991/go-package/app/timepkg"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	strDate := "2022-02-22 10:10:10"
	if _, err := time.Parse("2006-01-02", strDate); err == nil {
		fmt.Printf("2006-01-02\n")
		return
	} else if _, err := time.Parse("2006-01-02 15:04:05", strDate); err == nil {
		fmt.Printf("2006-01-02 15:04:05\n")
		return
	} else if _, err := time.Parse("2006-01-02 15:04:05.000", strDate); err == nil {
		fmt.Printf("2006-01-02 15:04:05.000\n")
		return
	}
}

// 获取时间
func TestGetTime(t *testing.T) {
	//packageuse.GetTime1()
	//packageuse.GetTime2()
	//packageuse.GetTime3()
	//packageuse.GetTime4()
	//packageuse.GetTime5()
	timepkg.GetTime6()
}

// 字符串转时间
func TestStr2Date(t *testing.T) {
	//packageuse.Str2Date1()
	timepkg.Str2Date2()
}

// 时间比较
func TestCompareDate(t *testing.T) {
	timepkg.CompareDate()
}

// 时间计算
func TestCalculate(t *testing.T) {
	//packageuse.CalculateDate1()
	timepkg.CalculateDate2()
}

// 定时器
func TestTicker(t *testing.T) {
	//packageuse.Ticker1()
	timepkg.Ticker2()
}

// 延迟执行
func TestDelayExec(t *testing.T) {
	//packageuse.DelayExec1()
	timepkg.DelayExec2()
}
