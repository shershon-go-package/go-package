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

type TestListener struct{}

// ExecuteLocalTransaction 执行本地事务
// primitive.CommitMessageState : 提交
// primitive.RollbackMessageState : 回滚
// primitive.UnknowState : 触发回查函数 CheckLocalTransaction
func (t TestListener) ExecuteLocalTransaction(message *primitive.Message) primitive.LocalTransactionState {
	fmt.Println("执行本地事务")
	return primitive.UnknowState
}

// CheckLocalTransaction 回查函数
// primitive.CommitMessageState : 提交
// primitive.RollbackMessageState : 回滚
// primitive.UnknowState : 触发会查函数 CheckLocalTransaction
func (t TestListener) CheckLocalTransaction(ext *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("执行回查函数")
	return primitive.CommitMessageState
}

// 发送事务消息
func Transaction() {
	newTransactionProducer, err := rocketmq.NewTransactionProducer(
		&TestListener{},
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
	)
	defer func(newProducer rocketmq.TransactionProducer) {
		err := newProducer.Shutdown()
		if err != nil {
			fmt.Printf("关闭producer失败 err:%s ", err)
			os.Exit(1)
		}
	}(newTransactionProducer)
	if err != nil {
		fmt.Printf("生成producer失败 err:%s ", err)
		os.Exit(1)
	}
	if err = newTransactionProducer.Start(); err != nil {
		fmt.Printf("启动producer失败 err:%s ", err)
		os.Exit(1)
	}
	res, err := newTransactionProducer.SendMessageInTransaction(context.Background(), primitive.NewMessage("TransactionTopic", []byte("这是一条事务消息2")))
	if err != nil {
		fmt.Printf("消息发送失败 err:%s ", err)
		os.Exit(1)
	}
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s: 消息: %s发送成功 \n", nowStr, res.String())
	time.Sleep(time.Hour)
}
