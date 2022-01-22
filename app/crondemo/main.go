package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)
// 定义类型
type city struct {
	Name string
}
// 实现cron.Job接口
func (s city) Run()  {
	fmt.Printf("%s时间:%v\n",s.Name,time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	// 开启秒级别支持,默认支持最小粒度是分钟
	c := cron.New(cron.WithSeconds())
	// 添加定时任务
	_, _ = c.AddFunc("@every 1m", func() {
		fmt.Printf("时间:%v\n",time.Now().Format("2006-01-02 15:04:05"))
	})
	// 启动
	c.Start()
	// 防止程序直接退出
	for {}
}

