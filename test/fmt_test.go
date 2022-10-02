/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/02/23 9:40 PM
 */

package test

import (
	"shershon1991/go-standard-package/app/fmtpkg"
	"testing"
)

// 输出到控制台
func TestOutput2Console1(t *testing.T) {
	//fmt_pkg.Output2Console1()
	fmtpkg.Output2Console2()
	//fmt_pkg.Output2Console3()
}

// 输出到文件
func TestOutput2File(t *testing.T) {
	//fmt_pkg.Output2File1()
	fmtpkg.Output2File2()
}

// 生成字符串
func TestGenerateStr(t *testing.T) {
	fmtpkg.GenerateStr()
}

// 生成错误类型
func TestGenerateError(t *testing.T) {
	fmtpkg.GenerateError()
}
