/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/20 12:28 PM
 */

package goroutinesync

import (
	"fmt"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func WaitGroup() {
	for i := 0; i < 100; i++ {
		go incrNum3()
	}
	// 等待其他协程处理完毕（共享变量为0）
	waitGroup.Wait()
	fmt.Println("done")
}

func incrNum3() {
	// 添加需要等待的协程数量（共享变量+1）
	waitGroup.Add(1)
	// do something
	// 标记当前协程处理完毕（共享变量-1）
	waitGroup.Done()
}
