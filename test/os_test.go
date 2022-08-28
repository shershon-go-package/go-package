/**
 * @Author Shershon
 * @Description os包
 * @Date 2021/6/24 5:32 下午
 **/
package test

import (
	"fmt"
	"os"
	"testing"
)

// 系统相关
func TestSystemInfo(t *testing.T) {
	hostname, _ := os.Hostname()
	// 主机名: huideMacBook-Pro.local
	fmt.Printf("主机名:%v \n", hostname)
	// 调用者所在进程的进程ID: 42862
	fmt.Printf("调用者所在进程的进程ID: %v \n", os.Getpid())
	// 调用者所在进程的进程的父进程ID: 42861
	fmt.Printf("调用者所在进程的进程的父进程ID: %v \n", os.Getppid())
	getwd, _ := os.Getwd()
	fmt.Printf("当前目录路径: %v \n", getwd)

}

func TestExit(t *testing.T) {
	fmt.Println("调用前打印...")
	// 调用退出程序：code范围应在 0 <= x <= 125
	os.Exit(0)
	// 后面代码不会执行
	fmt.Println("调用后，这里不会输出")
}

// 环境变量相关
func TestEnv(t *testing.T) {
	// 所有环境变量
	fmt.Printf("所有环境变量:%+v \n", os.Environ())
	// 设置环境变量
	_ = os.Setenv("my-name", "张三")
	// 获取环境变量
	fmt.Printf("获取环境变量: %v \n", os.Getenv("my-name"))
	// 清空所有环境变量
	os.Clearenv()
	fmt.Printf("清空环境变量后:%+v \n", os.Environ())
}
