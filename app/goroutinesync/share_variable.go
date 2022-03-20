/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/20 12:28 PM
 */

package goroutinesync

import (
	"fmt"
	"sync"
	"time"
)

var num = 0

// 互斥锁
var mutex = sync.Mutex{}

// 读写锁
var rwMutex = sync.RWMutex{}

func ShareVaribale() {
	for i := 0; i < 100; i++ {
		go incrNum()
	}
	time.Sleep(3)
	fmt.Println(num)
}

func incrNum() {
	mutex.Lock()
	num = num + 1
	mutex.Unlock()
}
