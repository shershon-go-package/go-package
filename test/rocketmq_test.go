/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/18 9:42 PM
 */

package test

import (
	"shershon1991/go-standard-package/app/rocketmqpkg"
	"shershon1991/go-standard-package/app/rocketmqpkg/aliyunmq"
	"testing"
)

// 发送普通消息
func TestSimpleProducer(t *testing.T) {
	rocketmqpkg.Simple()
}

// 消费普通消息
func TestSimpleConsumer(t *testing.T) {
	rocketmqpkg.ConsumeSimple()
}

// 发送延时消息
func TestDelayProducer(t *testing.T) {
	rocketmqpkg.Delay()
}

// 消费延时消息
func TestDelayConsumer(t *testing.T) {
	rocketmqpkg.ConsumeDelay()
}

// 发送事务消息
func TestTransactionProducer(t *testing.T) {
	rocketmqpkg.Transaction()
}

// 消费事务消息
func TestTransactionConsumer(t *testing.T) {
	rocketmqpkg.ConsumeTransaction()
}

// aliyun-发送普通消息
func TestAliSimpleProducer(t *testing.T) {
	aliyunmq.Simple()
}

// aliyun-消费普通消息
func TestAliSimpleConsumer(t *testing.T) {
	aliyunmq.ConsumeSimple()
}
