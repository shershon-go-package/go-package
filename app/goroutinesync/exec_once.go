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

var num2 = 0
var once = sync.Once{}

func ExecOnce() {
	for i := 0; i < 100; i++ {
		go once.Do(incrNum2)
	}
	time.Sleep(3)
	fmt.Println(num2)
}

func incrNum2() {
	num2 = num2 + 1
}
