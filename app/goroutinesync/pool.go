/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/20 12:28 PM
 */

package goroutinesync

import (
	"sync"
)

var p = sync.Pool{
	// 当池子中没有对象了, 用于创建新对象
	New: func() interface{} {
		return "3"
	},
}

func Pool() {
	// 从池子中获取一个对象
	r := p.Get()
	// 用完后将对象返回池子中
	p.Put(r)
}
