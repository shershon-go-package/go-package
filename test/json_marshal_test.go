/**
 * @Author Mr.LiuQH
 * @Description encoding/json 测试使用
 * @Date 2021/7/7 10:41 上午
 **/
package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 学校
type School struct {
	Name    string `json:"schoolName"`
	Address string `json:"schoolAddress"`
}

// 学生
type StudentA struct {
	Name string `json:"name"`
	// 匿名字段,而且没有json标签
	School
}

// 序列化-匿名字段 (字段标签不冲突)
func TestAnonymousTagDifferent(t *testing.T) {
	var student = StudentA{
		Name: "小明",
		School: School{
			Name:    "北京大学",
			Address: "北京海淀区",
		},
	}
	jsonByte, _ := json.Marshal(student)
	fmt.Printf("json: %s \n", jsonByte)
}

// 班级
type Class struct {
	Name string `json:"name"` // 标签名和学生名一样
	Desc string `json:"desc"`
}

// 学生
type StudentB struct {
	Name string `json:"name"` // 标签名和班级名一样
	// 匿名字段,而且没有json标签
	Class
}

// 序列化-匿名字段 (字段标签冲突)
func TestAnonymousTagSame(t *testing.T) {
	var student = StudentB{
		Name: "小明",
		Class: Class{
			Name: "高二(1)班",
			Desc: "优秀班级",
		},
	}
	jsonByte, _ := json.Marshal(student)
	fmt.Printf("json: %s \n", jsonByte)
}

type Bird struct {
	Name string `json:"name"`
	// 匿名字段，有json标签
	Category `json:"category"`
}
type Category struct {
	Name string `json:"categoryName"`
}
// 序列化-匿名字段,有json标签
func TestAnonymousWithTag(t *testing.T) {
	b := Bird{
		Name: "喜鹊",
		Category:Category{Name: "鸟类"},
	}
	jsonByte, _ := json.Marshal(b)
	fmt.Printf("json: %s \n", jsonByte)
}



