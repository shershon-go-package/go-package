/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/18 9:42 PM
 */

package test

import (
	"shershon1991/go-standard-package/app/rocketmqpkg/consumer"
	"shershon1991/go-standard-package/app/rocketmqpkg/producer"
	"testing"
)

// 发送普通消息
func TestSimpleProducer(t *testing.T) {
	producer.Simple()
}

// 消费普通消息
func TestSimpleConsumer(t *testing.T) {
	consumer.Simple()
}

// 发送延时消息
func TestDelayProducer(t *testing.T) {
	producer.Delay()
}

// 消费延时消息
func TestDelayConsumer(t *testing.T) {
	consumer.Delay()
}

// 发送事务消息
func TestTransactionProducer(t *testing.T) {
	producer.Transaction()
}

// 消费事务消息
func TestTransactionConsumer(t *testing.T) {
	consumer.Transaction()
}
