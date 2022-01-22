/**
 * @Author Mr.LiuQH
 * @Description json 反序列化
 * @Date 2021/7/8 2:49 下午
 **/
package test

import (
	"encoding/json"
	"fmt"
	"testing"
)



func TestGenerateJson(t *testing.T) {
	s := GoStudy{
		Name: "Go语言高级编程",
		Desc: "这是一本Go学习书籍",
		BookClass:BookClass{
			Name: "IT行业书籍",
			Company: "xxx出版社",
		},
	}
	bytes, _ := json.Marshal(s)
	fmt.Printf("json: %s \n",bytes)
}

type BookClass struct {
	// 设置成标签都为 name
	Name  string  `json:"name"`
	Company string `json:"company"`
}
type GoStudy struct {
	//设置成标签都为 name
	Name string `json:"name"`
	Desc string `json:"desc"`
	// 匿名字段无 json标签
	BookClass
}
func TestUnMarshal(t *testing.T) {
	jsonStr := `{"desc":"这是一本Go学习书籍","company":"xxx出版社","name":"IT行业书籍"}`
	var goStudy GoStudy
	err := json.Unmarshal([]byte(jsonStr), &goStudy)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("反序列化结果: %+v\n",goStudy)
}

