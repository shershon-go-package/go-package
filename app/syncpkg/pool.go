/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/03/03 15:48
 */

package syncpkg

import (
	"fmt"
	"sync"
)

func TestPool() {
	// 创建对象池
	pool := sync.Pool{New: func() interface{} {
		return make([]string, 5)
	}}
	// 首次获取
	fmt.Printf("不设置直接获取: %v\n", pool.Get())
	// 设置后获取
	pool.Put([]string{"Hello", "World"})
	// 设置后获取
	fmt.Printf("设置后,第一次获取: %v\n", pool.Get())
	fmt.Printf("设置后,第二次获取: %v\n", pool.Get())
}
