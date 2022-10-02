/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/09/29 17:27
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
