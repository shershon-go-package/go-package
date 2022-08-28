/**
 * @Author Shershon
 * @Description os.exec使用
 * @Date 2021/6/25 4:03 下午
 **/
package test

import (
	"fmt"
	"os/exec"
	"testing"
)

// 在环境变量PATH中搜索可执行文件
func TestLookPath(t *testing.T) {
	path, err := exec.LookPath("go")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(path)
}

// 使用Run()执行命令
func TestRunExec(t *testing.T) {
	// 执行: go version
	cmd := exec.Command("sleep", "3s")
	// 具体执行
	err := cmd.Run()
	if err != nil {
		t.Errorf("执行失败:%v\n", err)
	}
	fmt.Printf("cmd.Path: %v \n", cmd.Path)
	fmt.Printf("cmd.Args: %v \n", cmd.Args)
}

// 使用Start执行命令
func TestStart(t *testing.T) {
	// 执行: go version
	cmd := exec.Command("sleep", "3s")
	// Start开始执行c包含的命令，但并不会等待该命令完成即返回
	err := cmd.Start()
	if err != nil {
		t.Errorf("执行失败:%v\n", err)
	}
	//Wait会阻塞直到该命令执行完成
	err = cmd.Wait()
	fmt.Printf("执行完成: %v \n", err)
	fmt.Printf("cmd.Path: %v \n", cmd.Path)
	fmt.Printf("cmd.Args: %v \n", cmd.Args)
}

// 执行命令并获取输出结果
func TestOutput(t *testing.T) {
	// 执行: go version
	output, _ := exec.Command("go", "version").Output()
	fmt.Printf("结果: %s \n", output)
	// 执行: du -sh .
	output2, _ := exec.Command("du", "-sh", ".").Output()
	fmt.Printf("结果: %s \n", output2)
}

// 执行命令并返回标准输出和错误输出合并的切片
func TestCombinedOutput(t *testing.T) {
	// 执行: go version-1 故意写错
	output, _ := exec.Command("go", "version-1").CombinedOutput()
	fmt.Printf("结果: %s \n", output)
}
