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

var m = sync.Map{}

func Map() {
	for i := 0; i < 100; i++ {
		go func() {
			m.Store("1", 1)
		}()
	}
	time.Sleep(time.Second * 2)
	// 遍历map
	m.Range(func(key, value interface{}) bool {
		// 返回false结束遍历
		return true
	})
	m.LoadOrStore("1", 3)
	//m.Delete("1")
	load, _ := m.Load("1")
	fmt.Println(load)
}
