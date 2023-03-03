/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/03/03 12:58
 */

package syncpkg

import (
	"fmt"
	"strconv"
	"sync"
)

// 声明全局变量，文件内容
var fileContent string

// 声明全局读写互斥锁
var rwMutex sync.RWMutex

// 声明全局等待组
var waitg sync.WaitGroup

func NoAsStructFieldRWMutex() {
	// 设置计数器
	waitg.Add(5)
	for i := 1; i <= 5; i++ {
		name := "同学" + strconv.Itoa(i)
		if i%2 == 0 {
			go readFile(name)
		} else {
			go writeFile(name, strconv.Itoa(i))
		}
	}
	// 等待所有计数器执行完成
	waitg.Wait()
	fmt.Println("运行结束!")
}

// 读文件
func readFile(name string) {
	// 释放读锁
	defer rwMutex.RUnlock()
	// 获取读锁
	rwMutex.RLock()
	// 打印读取内容
	fmt.Printf("%s 获取读锁，读取内容为: %s \n", name, fileContent)
	// 计数器减一
	waitg.Done()
}

// 写文件
func writeFile(name, s string) {
	// 释放写锁
	defer rwMutex.Unlock()
	// 获取写锁
	rwMutex.Lock()
	// 写入内容
	fileContent = fileContent + " " + s
	fmt.Printf("%s 获取写锁，写入内容: %s。 文件内容变成: %s \n", name, s, fileContent)
	// 计数器减一
	waitg.Done()
}

// 定义一个文件结构体
type fileResource struct {
	content string
	rwLock  sync.RWMutex
	wg      sync.WaitGroup
}

func AsStructFieldRWMutex() {
	// 声明结构体
	var file fileResource
	// 设置计数器
	file.wg.Add(5)
	for i := 1; i <= 5; i++ {
		name := "同学" + strconv.Itoa(i)
		if i%2 == 0 {
			go file.readFile(name)
		} else {
			go file.writeFile(name, strconv.Itoa(i))
		}
	}
	// 等待所有计数器执行完成
	file.wg.Wait()
	fmt.Println("运行结束!")
}

// 读文件
func (f *fileResource) readFile(name string) {
	defer f.rwLock.RUnlock()
	f.rwLock.RLock()
	fmt.Printf("%s 获取读锁，读取内容为: %s \n", name, f.content)
	f.wg.Done()
}

// 写文件
func (f *fileResource) writeFile(name, s string) {
	defer f.rwLock.Unlock()
	f.rwLock.Lock()
	f.content = f.content + " " + s
	fmt.Printf("%s 获取写锁，写入内容: %s。 文件内容变成: %s \n", name, s, f.content)
	f.wg.Done()
}
