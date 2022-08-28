package consumer

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"os"
	"time"
)

func Simple() {
	// 创建消费者
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})))
	if err != nil {
		fmt.Printf("create consumer err:%s \n", err)
		os.Exit(2)
	}
	err = c.Subscribe("Test", consumer.MessageSelector{},
		func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, msg := range ext {
				fmt.Printf("Subscribe result:%s \n", msg)
			}
			return consumer.ConsumeSuccess, nil
		})
	//
	//err := c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context,
	//	msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	//	for i := range msgs {
	//		fmt.Printf("subscribe callback: %v \n", msgs[i])
	//	}
	//
	//	return consumer.ConsumeSuccess, nil
	//})

	if err != nil {
		fmt.Printf("Subscribe  err:%s \n", err)
		os.Exit(1)
	}
	err = c.Start()
	if err != nil {
		fmt.Printf("Start  err:%s \n", err)
		os.Exit(1)
	}
	time.Sleep(time.Hour)
	_ = c.Shutdown()
}
