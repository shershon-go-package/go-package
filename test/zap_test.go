/**
 * @Author Shershon
 * @Description zap 使用
 * @Date 2021/6/18 10:52 上午
 **/
package test

import (
	"shershon1991/go-standard-package/app/packageuse"
	"testing"
)

// 创建日志记录器
func TestCreateLogger(t *testing.T) {
	packageuse.CreateLogger()
}

// 使用默认记录日志
func TestRecordLogWithDefault(t *testing.T) {
	packageuse.RecordLogWithDefault()
}

// 使用Sugar记录器
func TestRecordLogWithSuage(t *testing.T) {
	packageuse.TestRecordLogWithSugar()
}

// 定制Logger
func TestCustomLogger(t *testing.T) {
	//packageuse.Log2File()
	packageuse.Log2FileAndConsole()
}

// 文件切割和日志归档
func TestCutAndArchive(t *testing.T) {
	packageuse.CutAndArchive()
}
