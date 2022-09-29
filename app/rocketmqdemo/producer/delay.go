/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/09/29 17:25
 */

package producer

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"time"
)

// 发送延时消息
func Delay() {
	// 初始化生产者
	newProducer, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(1),
	)

	defer func(newProducer rocketmq.Producer) {
		err := newProducer.Shutdown()
		if err != nil {
			fmt.Printf("关闭producer失败 err:%s ", err)
			os.Exit(1)
		}
	}(newProducer)
	if err != nil {
		fmt.Printf("生成producer失败 err:%s ", err)
		os.Exit(1)
	}

	err = newProducer.Start()
	if err != nil {
		fmt.Printf("启动producer失败 err:%s ", err)
		os.Exit(1)
	}

	message := primitive.NewMessage("DelayTopic", []byte("一条延时消息"))
	// WithDelayTimeLevel 设置要消耗的消息延迟时间。参考延迟等级定义：1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
	// 延迟等级从1开始，例如设置param level=1，则延迟时间为1s。
	// 这里使用的是延时30s发送
	message.WithDelayTimeLevel(4)

	res, err := newProducer.SendSync(context.Background(), message)
	if err != nil {
		fmt.Printf("消息发送失败 err:%s ", err)
		os.Exit(1)
	}
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s: 消息: %s发送成功 \n", nowStr, res.String())
}
