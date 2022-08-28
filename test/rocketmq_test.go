/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/18 9:42 PM
 */

package test

import (
	"shershon1991/go-standard-package/app/rocketmqdemo/consumer"
	"shershon1991/go-standard-package/app/rocketmqdemo/producer"
	"testing"
)

func TestProducer(t *testing.T) {
	producer.Simple()
}

func TestConsumer(t *testing.T) {
	consumer.Simple()
}
