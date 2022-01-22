package test
import (
	"testing"
)

// 通过测试函数
func TestPass(t *testing.T) {
	t.Log("这个是通过测试函数")
}

// 不通过测试函数
func TestFail(t *testing.T) {
	t.Error("运行测试失败！")
}