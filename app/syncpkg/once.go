/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/03/03 15:39
 */

package syncpkg

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func TestOnce() {
	echo := func() {
		t := time.Now().Unix()
		fmt.Printf("当前时间戳 %v \n", strconv.FormatInt(t, 10))
	}
	var once sync.Once
	// 虽然遍历调用，但是只会执行一次
	for i := 1; i <= 10; i++ {
		go func() {
			once.Do(echo)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("运行结束！")
}
