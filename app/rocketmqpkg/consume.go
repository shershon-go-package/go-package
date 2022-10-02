/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/10/02 14:03
 */

package rocketmqpkg

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"os"
	"time"
)

// 消费普通消息
func ConsumeSimple() {
	newPushConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("test"),
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
	)

	defer func(newPushConsumer rocketmq.PushConsumer) {
		err := newPushConsumer.Shutdown()
		if err != nil {
			fmt.Printf("关闭consumer失败 err:%s \n", err)
			os.Exit(1)
		}
	}(newPushConsumer)

	if err != nil {
		fmt.Printf("生成consumer失败 err:%s \n", err)
		os.Exit(1)
	}

	err = newPushConsumer.Subscribe("SimpleTopic", consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, msg := range msgs {
				nowStr := time.Now().Format("2006-01-02 15:04:05")
				fmt.Printf("%s 读取到一条消息,消息内容: %s \n", nowStr, string(msg.Body))
			}
			return consumer.ConsumeSuccess, nil
		},
	)

	if err != nil {
		fmt.Printf("读取消息失败 err:%s \n", err)
		os.Exit(1)
	}

	err = newPushConsumer.Start()
	if err != nil {
		fmt.Printf("启动consumer失败 err:%s \n", err)
		os.Exit(1)
	}

	time.Sleep(time.Hour)
}

// 消费延时消息
func ConsumeDelay() {
	newPushConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("test"),
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
	)

	defer func(newPushConsumer rocketmq.PushConsumer) {
		err := newPushConsumer.Shutdown()
		if err != nil {
			fmt.Printf("关闭consumer失败 err:%s \n", err)
			os.Exit(1)
		}
	}(newPushConsumer)

	if err != nil {
		fmt.Printf("生成consumer失败 err:%s \n", err)
		os.Exit(1)
	}

	err = newPushConsumer.Subscribe("DelayTopic", consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, msg := range msgs {
				nowStr := time.Now().Format("2006-01-02 15:04:05")
				fmt.Printf("%s 读取到一条消息,消息内容: %s \n", nowStr, string(msg.Body))
			}
			return consumer.ConsumeSuccess, nil
		},
	)

	if err != nil {
		fmt.Printf("读取消息失败 err:%s \n", err)
		os.Exit(1)
	}

	err = newPushConsumer.Start()
	if err != nil {
		fmt.Printf("启动consumer失败 err:%s \n", err)
		os.Exit(1)
	}

	time.Sleep(time.Hour)
}

// 消费事务消息
func ConsumeTransaction() {
	newPushConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("test"),
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
	)

	defer func(newPushConsumer rocketmq.PushConsumer) {
		err := newPushConsumer.Shutdown()
		if err != nil {
			fmt.Printf("关闭consumer失败 err:%s \n", err)
			os.Exit(1)
		}
	}(newPushConsumer)

	if err != nil {
		fmt.Printf("生成consumer失败 err:%s \n", err)
		os.Exit(1)
	}

	err = newPushConsumer.Subscribe("TransactionTopic", consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, msg := range msgs {
				nowStr := time.Now().Format("2006-01-02 15:04:05")
				fmt.Printf("%s 读取到一条消息,消息内容: %s \n", nowStr, string(msg.Body))
			}
			return consumer.ConsumeSuccess, nil
		},
	)

	if err != nil {
		fmt.Printf("读取消息失败 err:%s \n", err)
		os.Exit(1)
	}

	err = newPushConsumer.Start()
	if err != nil {
		fmt.Printf("启动consumer失败 err:%s \n", err)
		os.Exit(1)
	}

	time.Sleep(time.Hour)
}
