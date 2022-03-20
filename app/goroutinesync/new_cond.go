/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/20 12:28 PM
 */

package goroutinesync

import (
	"sync"
)

var mutex4 = &sync.Mutex{}
var cond = sync.NewCond(mutex4)

func NewCond() {
	for i := 0; i < 100; i++ {
		go incrNum4()
	}
	// 发送命令给一个随机获得锁的协程
	cond.Signal()
	// 发送命令给所有获得锁的协程
	cond.Broadcast()
}

func incrNum4() {
	// 获取锁，标识当前协程可以处理命令
	cond.L.Lock()
	// 可添加退出执行命令队列的条件
	for true {
		// 等待命令
		cond.Wait()
		// do something
	}
	// 释放锁，标识当前协程处理完毕，退出协程
	cond.L.Unlock()
}
