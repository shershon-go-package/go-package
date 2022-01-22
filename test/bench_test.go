package test

import (
	"fmt"
	"testing"
	"time"
)

// 测试函数Sprintf性能
func BenchmarkCompute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("成绩:%d",80)
	}
}

// 并行测试
func BenchmarkParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = fmt.Sprintf("成绩:%d",80)
		}
	})
}

// 重置时间使用
func BenchmarkTime(b *testing.B) {
	// 准备工作
	time.Sleep(time.Second * 3)
	// 重置时间
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("hello:%v","word")
	}
}