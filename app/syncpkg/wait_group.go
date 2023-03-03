/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/20 12:28 PM
 */

package syncpkg

import (
	"fmt"
	"sync"
	"time"
)

func NoWaitGroup() {
	// 创建通道
	intChan := make(chan int)

	// 计算1-50的和
	go func(intChan chan int) {
		sum := 0
		for i := 1; i <= 50; i++ {
			sum += i
		}
		intChan <- sum
	}(intChan)
	// 计算51-100的和
	go func(intChan chan int) {
		sum := 0
		for i := 51; i <= 100; i++ {
			sum += i
		}
		intChan <- sum
	}(intChan)
	// 另外创建个channl聚合结果
	go func(intChan chan int) {
		sum1 := <-intChan
		sum2 := <-intChan
		fmt.Printf("sum1 = %d, sum2 = %d \nsum1 + sum2 = %d \n", sum1, sum2, sum1+sum2)
	}(intChan)

	// 注意,需求手动sleep
	time.Sleep(time.Second)
	fmt.Println("运行结束！")
}

func UseWaitGroup() {
	var wg sync.WaitGroup
	// 设置，需要等待3个协程执行完成
	wg.Add(3)
	// 创建通道
	intChan := make(chan int)
	// 计算1-50的和
	go func(intChan chan int, wg *sync.WaitGroup) {
		sum := 0
		for i := 1; i <= 50; i++ {
			sum += i
		}
		intChan <- sum
		// 计数器减一
		wg.Done()
	}(intChan, &wg)
	// 计算51-100的和
	go func(intChan chan int, wg *sync.WaitGroup) {
		sum := 0
		for i := 51; i <= 100; i++ {
			sum += i
		}
		intChan <- sum
		// 计数器减一
		wg.Done()
	}(intChan, &wg)
	// 另外创建个channl聚合结果
	go func(intChan chan int, wg *sync.WaitGroup) {
		sum1 := <-intChan
		sum2 := <-intChan
		fmt.Printf("sum1 = %d, sum2 = %d \nsum1 + sum2 = %d \n", sum1, sum2, sum1+sum2)
		// 计数器减一
		wg.Done()
	}(intChan, &wg)
	// 阻塞，直到等待组的计数器等于0
	wg.Wait()
	fmt.Println("运行结束!")
}
