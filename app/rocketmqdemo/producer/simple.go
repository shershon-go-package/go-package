package producer

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
)

func Simple() {
	// 初始化生产者
	p, err := rocketmq.NewProducer(
		producer.WithGroupName("testGroup"),
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2),
	)
	if err != nil {
		fmt.Printf("create err:%s ", err)
		os.Exit(1)
	}
	err = p.Start()
	if err != nil {
		fmt.Printf("start err:%s ", err)
		os.Exit(2)
	}
	for i := 0; i < 10; i++ {
		msg := &primitive.Message{
			Topic: "Test",
			Body:  []byte("Hello Rocketmq " + strconv.Itoa(i)),
		}
		msg.WithTag("学生")
		msg.WithKeys([]string{"小学"})
		sync, err := p.SendSync(context.Background(), msg)
		if err != nil {
			fmt.Printf("send err: %s \n", err)
		} else {
			fmt.Printf("send succcess res=%s\n ", sync)
		}
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("close err:%s ", err)
	}
}
