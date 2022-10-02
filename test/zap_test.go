/**
 * @Author Shershon
 * @Description zap 使用
 * @Date 2021/6/18 10:52 上午
 **/
package test

import (
	"shershon1991/go-standard-package/app/zappkg"
	"testing"
)

// 创建日志记录器
func TestCreateLogger(t *testing.T) {
	zappkg.CreateLogger()
}

// 使用默认记录日志
func TestRecordLogWithDefault(t *testing.T) {
	zappkg.RecordLogWithDefault()
}

// 使用Sugar记录器
func TestRecordLogWithSuage(t *testing.T) {
	zappkg.TestRecordLogWithSugar()
}

// 定制Logger
func TestCustomLogger(t *testing.T) {
	//zappkg.Log2File()
	zappkg.Log2FileAndConsole()
}

// 文件切割和日志归档
func TestCutAndArchive(t *testing.T) {
	zappkg.CutAndArchive()
}
