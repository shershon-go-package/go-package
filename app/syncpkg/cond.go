/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/03/03 15:26
 */

package syncpkg

import (
	"fmt"
	"sync"
	"time"
)

func TestCond() {
	// 声明互斥锁
	var mutex sync.Mutex
	// 声明条件变量
	cond := sync.NewCond(&mutex)
	for i := 1; i <= 10; i++ {
		go func(i int) {
			defer cond.L.Unlock()
			cond.L.Lock()
			// 等待通知,阻塞当前协程
			cond.Wait()
			// 等待通知后打印输出
			fmt.Printf("输出:%d ! \n", i)
		}(i)
	}

	// 单个通知
	time.Sleep(time.Second)
	fmt.Println("单个通知A！")
	cond.Signal()
	time.Sleep(time.Second)
	fmt.Println("单个通知B！")
	cond.Signal()

	// 广播通知
	time.Sleep(time.Second)
	fmt.Println("广播通知！并睡眠1秒，等待其他协程输出!")
	cond.Broadcast()

	// 等待其他协程处理完
	time.Sleep(time.Second)
	fmt.Println("运行结束！")
}
