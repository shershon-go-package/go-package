/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/03/03 10:48
 */

package syncpkg

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 声明全局等待组
var wg sync.WaitGroup

// 声明全局锁
var mutex sync.Mutex

// 声明全局余票
var ticket int = 10

func NoAsStructFieldMutex() {
	// 设置等待组计数器
	wg.Add(3)
	// 窗口卖票
	go saleTicket("窗口A", &wg)
	go saleTicket("窗口B", &wg)
	go saleTicket("窗口C", &wg)
	wg.Wait()
	fmt.Println("运行结束!")
}

// 卖票流程
func saleTicket(windowName string, wg *sync.WaitGroup) {
	// 卖票流程结束后关闭
	defer wg.Done()
	for {
		// 加锁
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(10 * time.Millisecond)
			ticket--
			fmt.Printf("%s 卖出一张票，余票: %d \n", windowName, ticket)
		} else {
			fmt.Printf("%s 票已卖完! \n", windowName)
			mutex.Unlock()
			break
		}
		// 解锁
		mutex.Unlock()
	}
}

// 声明一个票池
type ticketPool struct {
	over int
	lock sync.Mutex
	wg   sync.WaitGroup
}

func AsStructFieldMutex() {
	// 创建一个票池
	ticketP := ticketPool{over: 10}
	fmt.Printf("T:%T v: %v \n", ticketP, ticketP)
	// 设置窗口数量
	windowNum := 3
	ticketP.wg.Add(windowNum)
	// 定义3个窗口售票
	for i := 1; i <= windowNum; i++ {
		go ticketP.sellTicket("窗口" + strconv.Itoa(i))
	}
	ticketP.wg.Wait()
	fmt.Println("运行结束!")
}

func (t *ticketPool) sellTicket(windowName string) {
	// 等待组减一
	defer t.wg.Done()
	for {
		// 加锁
		t.lock.Lock()
		if t.over > 0 {
			time.Sleep(10 * time.Millisecond)
			t.over--
			fmt.Printf("%s 卖出一张票，余票: %d \n", windowName, t.over)
		} else {
			t.lock.Unlock()
			fmt.Printf("%s 票已卖完！\n", windowName)
			break
		}
		// 正常售票流程解锁
		t.lock.Unlock()
	}
}
