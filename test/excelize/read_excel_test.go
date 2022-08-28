/**
 * @Description https://xuri.me/excelize/zh-hans/base/installation.html#read
 **/
package excelize

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
)

// 读取表格
func TestRead(t *testing.T) {
	// 打开表格文件
	openFile, err := excelize.OpenFile("../../tmp/line.xlsx")
	if err != nil {
		t.Error("打开表格文件失败: " + err.Error())
		return
	}
	// 读取指定工作表所有数据
	rows, err := openFile.GetRows("Sheet1")
	if err != nil {
		t.Error("读取失败: " + err.Error())
		return
	}
	for _, row := range rows {
		fmt.Printf("%+v\n", row)
	}
	fmt.Println("执行完成!")
}
