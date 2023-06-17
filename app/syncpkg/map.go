/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/03/03 16:17
 */

package syncpkg

import (
	"fmt"
	"sync"
)

func TestMap1() {
	// 定义map 类型
	var syncMap sync.Map
	// 新增
	syncMap.Store("name", "张三")
	load, _ := syncMap.Load("name")
	fmt.Printf("Store新增 -> name:%v\n", load)
	// 找到则不更新，返回旧值
	store, loaded := syncMap.LoadOrStore("name", "李四")
	fmt.Printf("找到则返回旧值 -> name:%v loaded:%v \n", store, loaded)
	// 找不到则新增
	age, loaded := syncMap.LoadOrStore("age", 31)
	fmt.Printf("找不到则新增 -> age:%v loaded:%v \n", age, loaded)
}

func TestMap2() {
	// 定义map 类型
	var syncMap sync.Map
	// 新增
	syncMap.Store("name", "张三")
	syncMap.Store("age", 20)

	// 找到情况
	andDelete, loaded := syncMap.LoadAndDelete("name")
	fmt.Printf("找到-> val:%v loaded:%v \n", andDelete, loaded)
	search, ok := syncMap.Load("name")
	fmt.Printf("删除name后查找-> search:%v ok:%v \n", search, ok)

	// 找不到情况
	andDelete2, loaded := syncMap.LoadAndDelete("name2")
	fmt.Printf("找不到-> val:%v loaded:%v \n", andDelete2, loaded)

	syncMap.Delete("age")
	searchAge, ok := syncMap.Load("age")
	fmt.Printf("删除age后查找-> searchAge:%v ok:%v \n", searchAge, ok)
}

func TestMap3() {
	// 定义map 类型
	var syncMap sync.Map
	syncMap.Store("name", "张三")
	syncMap.Store("age", 31)
	syncMap.Store("home", "武汉洪山区")
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v value: %v \n", key, value)
		return true
	})
}
