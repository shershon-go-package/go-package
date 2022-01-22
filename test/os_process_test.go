/**
 * @Author Mr.LiuQH
 * @Description os包中的process
 * @Date 2021/6/25 10:42 上午
 **/
package test

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestFindProcess(t *testing.T) {
	// 获取当前进程信息
	process, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("当前线程信息%+v \n", process)
	// 让线程立刻退出
	//_ = process.Kill()
	// 阻塞进程直到退出，返回一个描述ProcessState描述进程的状态和可能的错误
	// Wait方法会释放绑定到进程p的所有资源。

}
// 创建进程
func TestCreateProcess(t *testing.T) {
	// 创建一个新进程执行； ls /
	args := []string{"/"}
	process, err := os.StartProcess("/bin/ls", args, &os.ProcAttr{})
	if err != nil {
		t.Error("创建新进程失败: " + err.Error())
		return
	}
	fmt.Printf("当前新线程信息: %+v \n", process)

	// 2秒后向进程发送信号
	time.AfterFunc(2 * time.Second, func() {
		fmt.Println("发送进程退出信号...")
		_ = process.Signal(os.Kill)
	})
	// 手动阻塞看是否执行：发送信号
	time.Sleep(3 * time.Second)

	// 等待进程退出,返回ProcessState类型
	processState, _ := process.Wait()
	// 返回一个已退出进程的id
	fmt.Printf("当前进程Id: %v \n", processState.Pid())
	// 报告进程是否已退出
	fmt.Printf("进程是否已退出: %v \n", processState.Exited())
	// 报告进程是否成功退出，如在Unix里以状态码0退出。
	fmt.Printf("进程是否成功退出: %v \n", processState.Success())
	// 返回已退出进程及其子进程耗费的系统CPU时间。
	fmt.Printf("进程及子进程耗费系统CPU时间: %v \n", processState.SystemTime())
	// 返回已退出进程及其子进程耗费的用户CPU时间。
	fmt.Printf("进程及子进程耗费用户CPU时间: %v \n", processState.UserTime())
	fmt.Printf("进程状态: %s \n",processState.String())
}