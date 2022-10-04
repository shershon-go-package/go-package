package test

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"testing"
	"time"
)

type Stu struct {
	Id         int
	Name       string
	Age        int
	Like       []string
	Home       string
	Province   string
	City       string
	Area       string
	SchoolName string
	ClassName  string
	Phone      string
	Desc       map[string]interface{}
}

var jsonT = jsoniter.ConfigCompatibleWithStandardLibrary

// 使用jsonIter
func BenchmarkUnMarshalIter(b *testing.B) {
	jsonStr := `{"Id":96,"Name":"张三","Age":30,"Like":["打游戏","看动漫"],"Home":"北京市昌平区回龙观东大街紫和园2号楼2单元1101","Province":"北京","City":"北京","Area":"昌平区","SchoolName":"北京大学","ClassName":"某某班","Phone":"172222222222","Desc":{"xx":"xx","性别":"男"}}`
	for i := 0; i < b.N; i++ {
		var d Stu
		_ = jsonT.Unmarshal([]byte(jsonStr), &d)
	}
}

// 使用encoding/json
func BenchmarkUnMarshalJson(b *testing.B) {
	jsonStr := `{"Id":96,"Name":"张三","Age":30,"Like":["打游戏","看动漫"],"Home":"北京市昌平区回龙观东大街紫和园2号楼2单元1101","Province":"北京","City":"北京","Area":"昌平区","SchoolName":"北京大学","ClassName":"某某班","Phone":"172222222222","Desc":{"xx":"xx","性别":"男"}}`
	for i := 0; i < b.N; i++ {
		var d Stu
		_ = json.Unmarshal([]byte(jsonStr), &d)
	}
}

func BenchmarkMarshalIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := Stu{
			Id:         1,
			Name:       "张三",
			Age:        30,
			Like:       []string{"打游戏", "看动漫"},
			Home:       "北京市昌平区回龙观东大街紫和园2号楼2单元1101",
			City:       "北京",
			Province:   "北京",
			Area:       "昌平区",
			SchoolName: "北京大学",
			ClassName:  "某某班",
			Phone:      "172222222222",
			Desc:       map[string]interface{}{"性别": "男", "xx": "xx"},
		}
		_, _ = jsonT.Marshal(d)
	}
}
func BenchmarkMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := Stu{
			Id:         1,
			Name:       "张三",
			Age:        30,
			Like:       []string{"打游戏", "看动漫"},
			Home:       "北京市昌平区回龙观东大街紫和园2号楼2单元1101",
			City:       "北京",
			Province:   "北京",
			Area:       "昌平区",
			SchoolName: "北京大学",
			ClassName:  "某某班",
			Phone:      "172222222222",
			Desc:       map[string]interface{}{"性别": "男", "xx": "xx"},
		}
		_, _ = json.Marshal(d)
	}
}

type Order struct {
	Id       int               `json:"id,omitempty"`
	OrderNum string            `json:"orderNum,omitempty"`
	Money    float64           `json:"money,omitempty"`
	PayTime  time.Time         `json:"payTime"`
	Extend   map[string]string `json:"extend"`
}

// --------------------------------------------------------------------------------

// 序列化
func TestMarshalToString(t *testing.T) {
	order := Order{
		Id:       10,
		OrderNum: "100200300",
		Money:    99.99,
		PayTime:  time.Now(),
		Extend:   map[string]string{"name": "张三"},
	}
	// 直接转成字符串
	jsonStr, _ := jsoniter.MarshalToString(order)
	fmt.Println("jsonStr:", jsonStr)
	//// 转成字节[]bythn
	//marshal, _ := jsoniter.Marshal(order)
	//fmt.Println("marshal:", string(marshal))
}

// 序列化字节
func TestMarshalToByte(t *testing.T) {
	order := Order{
		Id:       10,
		OrderNum: "100200300",
		Money:    99.99,
		PayTime:  time.Now(),
		Extend:   map[string]string{"name": "张三"},
	}
	// 转成字节[]byte
	marshal, _ := jsoniter.Marshal(order)
	fmt.Println("marshal:", marshal)
}

// 反序列化
func TestUnmarshalTmp(t *testing.T) {
	str := `{"id":10,"orderNum":"100200300","money":99.99,"payTime":"2021-12-28T23:44:36.258311+08:00","extend":{"name":"张三"}}`
	var order Order
	// 从字符串反序列化
	_ = jsoniter.UnmarshalFromString(str, &order)
	fmt.Println("order:", order)

	var order2 Order
	// 从[]byte反序列化
	_ = jsoniter.Unmarshal([]byte(str), &order2)
	fmt.Println("order2:", order2)
}

// 使用encoding/json反序列化
func TestTypeCovert(t *testing.T) {
	// money 为float64类型，故意设置成字符串
	str := `{"id":10,"orderNum":"100200300","money":"99.99","payTime":"2021-12-28T23:44:36.258311+08:00","extend":{"name":"张三"}}`
	var order Order
	// 使用标准库解析
	err := json.Unmarshal([]byte(str), &order)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println("order: ", order)
}

// 使用json-iterator反序列化
func TestTypeCovertWithJsonIter(t *testing.T) {
	// money 为float64类型，故意设置成字符串
	str := `{"id":10,"orderNum":"100200300","money":"99.99","payTime":"2021-12-28T23:44:36.258311+08:00","extend":{"name":"张三"}}`
	var order Order
	// 定义
	var jsonNew = jsoniter.ConfigCompatibleWithStandardLibrary
	// 自适应类型
	extra.RegisterFuzzyDecoders()
	err := jsonNew.Unmarshal([]byte(str), &order)
	if err != nil {
		fmt.Println("jsonNew err:", err)
	}
	fmt.Println("order: ", order)
}

// 定义结构体，里面包含私有属性
type Demo struct {
	FirstName string `json:"firstName,omitempty"`
	lastName  string
}

// 解析私有属性
func TestDealPrivate(t *testing.T) {
	d := Demo{
		FirstName: "张",
		lastName:  "三丰",
	}
	// 开启解析私有属性，注：私有属性不能有json标签，否则不能解析
	extra.SupportPrivateFields()
	var jsonNew = jsoniter.ConfigCompatibleWithStandardLibrary
	res, err := jsonNew.MarshalToString(d)
	fmt.Println("序列化-err:", err)
	fmt.Println("序列化-res:", res)
	// 反序列化私有属性
	jsonStr := `{"firstName":"张","lastName":"三丰"}`
	var d2 Demo
	err = jsonNew.UnmarshalFromString(jsonStr, &d2)
	fmt.Println("反序列化-err:", err)
	fmt.Println("反序列化-d2:", d2)
}

type timeDemo struct {
	CreateTime time.Time `json:"createTime"`
}

// 解析时间
func TestDealTime(t *testing.T) {
	td := timeDemo{CreateTime: time.Now()}
	var jsonNew = jsoniter.ConfigCompatibleWithStandardLibrary
	// 转成以秒
	extra.RegisterTimeAsInt64Codec(time.Second)
	// 序列化
	res, err := jsonNew.MarshalToString(td)
	fmt.Println("时间序列化-err:", err)
	fmt.Println("时间序列化-res:", res)
	// 反序列化
	str := `{"createTime":1640791445}`
	var tds timeDemo
	err = jsonNew.UnmarshalFromString(str, &tds)
	fmt.Println("时间反序列化-err:", err)
	fmt.Println("时间反序列化-res:", tds)
}

// 直接读取json 字符串
func TestReadJsonString(t *testing.T) {
	str := `{
    "id":10,
    "extend":{
        "name":"张三"
    },
    "desc":[
        {
            "score":"100"
        },
        {
            "score":"90"
        }
    ]
}`
	fmt.Println("id:", jsoniter.Get([]byte(str), "id").ToInt())
	fmt.Println("extend.name:", jsoniter.Get([]byte(str), "extend", "name").ToString())
	fmt.Println("desc.0.score:", jsoniter.Get([]byte(str), "desc", 0, "score").ToInt())
	fmt.Println("desc.1.score:", jsoniter.Get([]byte(str), "desc", 1, "score").ToInt())
}

type PayInfo struct {
	OrderId  int
	PayMoney float64 `json:"payMoney"`
}

// 没有json标签的大驼峰属性，会转成下划线变量
func TestSetNamingStrategy(t *testing.T) {
	payInfo := PayInfo{
		OrderId:  100,
		PayMoney: 9.9,
	}
	// -------- 序列化 --------
	// 设置后，没有json标签的属性，会自动转成 xx_xx
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	res, _ := jsoniter.MarshalToString(payInfo)
	fmt.Println("序列化:", res)
	// -------- 反序列化 --------
	var p PayInfo
	_ = jsoniter.UnmarshalFromString(res, &p)
	fmt.Printf("反序列化:%+v \n", p)
}

func TestBinaryAsStringCodec(t *testing.T) {
	str := `{"OrderId":100,"payMoney":9.9}`
	binaryAsStringExtension := extra.BinaryAsStringExtension{}
	jsoniter.RegisterExtension(&binaryAsStringExtension)
	res, err := jsoniter.MarshalToString([]byte(str))
	fmt.Println("时间反序列化-err:", err)
	fmt.Println("时间反序列化-res:", res)
}

func init() {
	// 在init中设置一次即可
	extra.RegisterFuzzyDecoders()
}
func TestUseRegisterFuzzyDecodersWithGoroutine(t *testing.T) {
	payInfo := PayInfo{
		OrderId:  100,
		PayMoney: 9.9,
	}
	// 在多个协程中设置extra.RegisterFuzzyDecoders()
	for i := 0; i < 10; i++ {
		go func() {
			jsonNew := jsoniter.ConfigCompatibleWithStandardLibrary
			_, _ = jsonNew.MarshalToString(payInfo)
		}()
	}
	time.Sleep(time.Second * 1)
	fmt.Println("ok")
}
