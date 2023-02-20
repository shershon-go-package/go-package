/**
 * @Description
 **/
package excelize

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	f := float64(rand.Intn(10)) + rand.Float64()
	fmt.Printf("T:%T v:%.2f\n", f, f)

}

// 表格写入(单个赋值)
func TestWriteExcelBySetCellValue(t *testing.T) {
	excel := excelize.NewFile()
	// 创建一个工作表
	sheet := excel.NewSheet("学校")
	// 单个赋值
	_ = excel.SetCellValue("学校", "A2", "北京大学")
	_ = excel.SetCellValue("学校", "A3", "南京大学")
	// 设置sheet1的内容(默认创建)
	excel.SetCellValue("Sheet1", "A1", "张三")
	excel.SetCellValue("Sheet1", "A2", "小明")
	// 设置默认工作表
	excel.SetActiveSheet(sheet)
	// 保存表格
	if err := excel.SaveAs("../../tmp/test.xlsx"); err != nil {
		t.Error(err)
		return
	}
	fmt.Println("执行完成")
}

// 表格写入(按行写入)
func TestWriteByLine(t *testing.T) {
	excel := excelize.NewFile()
	// 写入标题
	titleSlice := []interface{}{"序号", "姓名", "年龄", "性别"}
	_ = excel.SetSheetRow("Sheet1", "A1", &titleSlice)
	data := []interface{}{
		[]interface{}{1, "张三", 19, "男"},
		[]interface{}{2, "小丽", 18, "女"},
		[]interface{}{3, "小明", 20, "男"},
	}
	// 遍历写入数据
	for key, datum := range data {
		axis := fmt.Sprintf("A%d", key+2)
		// 利用断言，转换类型
		tmp, _ := datum.([]interface{})
		_ = excel.SetSheetRow("Sheet1", axis, &tmp)
	}
	// 保存表格
	if err := excel.SaveAs("../../tmp/line.xlsx"); err != nil {
		t.Error(err)
		return
	}
	fmt.Println("执行完成")
}

// 表格写入(按行流式写入)
func TestWriteByStream(t *testing.T) {
	excel := excelize.NewFile()
	// 获取流式写入器
	streamWriter, err := excel.NewStreamWriter("Sheet1")
	if err != nil {
		t.Error("获取流式写入器失败: " + err.Error())
		return
	}
	// 按行写入
	if err := streamWriter.SetRow("A1", []interface{}{"序号", "商品码", "价格"}); err != nil {
		t.Error("获取流式写入器失败: " + err.Error())
		return
	}
	// 制作数据
	// 设置随机因子
	rand.Seed(time.Now().Unix())
	for i := 2; i < 500000; i++ {
		tmp := []interface{}{
			i,
			fmt.Sprintf("P-%d", rand.Intn(100000000)),
			fmt.Sprintf("%.2f", float64(rand.Intn(10))+rand.Float64()),
		}
		_ = streamWriter.SetRow("A"+strconv.Itoa(i), tmp)
	}
	// 调用 Flush 函数来结束流式写入过程
	if err = streamWriter.Flush(); err != nil {
		t.Error("结束流式写入失败: " + err.Error())
		return
	}
	// 保存表格
	if err := excel.SaveAs("../../tmp/stream.xlsx"); err != nil {
		t.Error(err)
		return
	}
	fmt.Println("执行完成")
}
