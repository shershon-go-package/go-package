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

// 发送普通消息
func Simple() {
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

	res, err := newProducer.SendSync(context.Background(), primitive.NewMessage("SimpleTopic", []byte("一条简单消息")))
	if err != nil {
		fmt.Printf("消息发送失败 err:%s ", err)
		os.Exit(1)
	}
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s: 消息: %s发送成功 \n", nowStr, res.String())
}
