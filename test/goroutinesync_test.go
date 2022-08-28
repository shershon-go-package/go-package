/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/20 12:33 PM
 */

package test

import (
	"shershon1991/go-standard-package/app/goroutinesync"
	"testing"
)

// 共享变量
func TestShareVariable(t *testing.T) {
	goroutinesync.ShareVaribale()
}

// 仅执行一次
func TestExecOnce(t *testing.T) {
	goroutinesync.ExecOnce()
}

// 等待其他协程处理
func TestWaitGroup(t *testing.T) {
	goroutinesync.WaitGroup()
}

// 消息通知
func TestNewCond(t *testing.T) {
	goroutinesync.NewCond()
}

// 多协程 map
func TestMap(t *testing.T) {
	goroutinesync.Map()
}

// 多协程对象池
func TestPool(t *testing.T) {
	goroutinesync.Pool()
}
