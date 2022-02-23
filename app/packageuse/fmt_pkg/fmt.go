package fmt_pkg

import (
	"fmt"
	"os"
)

/******************** 输出到控制台 start ******************/
// Printf
type People struct {
	name string
	age  int
}

func Output2Console1() {
	// 打印字符串
	str := "hello world"
	fmt.Printf("%%s -> %s\n", str)
	fmt.Printf("%%v -> %v\n", str)
	// 打印布尔类型
	b := true
	fmt.Printf("%%t -> %t\n", b)
	fmt.Printf("%%v -> %v\n", b)
	// 打印浮点型
	f := 1.234567890
	fmt.Printf("默认宽度，默认精度:%f\n", f)
	fmt.Printf("宽度9，默认精度:%9f\n", f)
	fmt.Printf("默认宽度，精度2:%.2f\n", f)
	fmt.Printf("宽度9，精度2:%9.2f\n", f)
	fmt.Printf("宽度9，精度0:%9.f\n", f)
	// 打印整数
	a := 54321
	fmt.Printf("十进制:%d\n", a)
	fmt.Printf("二进制:%b\n", a)
	fmt.Printf("八进制:%o\n", a)
	fmt.Printf("十六进制:%x\n", a)
	fmt.Printf("十六进制，字母大写:%X\n", a)
	// 打印指针
	str2 := "hello"
	fmt.Printf("打印指针:%p\n", &str2)
	// 打印变量类型
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arr类型:%T\n", arr)
	map1 := map[string]int{
		"张三": 20,
		"李四": 25,
	}
	fmt.Printf("map1类型:%T\n", map1)
	fruit := struct {
		name  string
		price float64
	}{"香蕉", 4.99}
	fmt.Printf("fruit类型:%T\n", fruit)
	user := People{"李四", 30}
	fmt.Printf("user类型:%T\n", user)
	// 打印结构体
	people := struct {
		name, home, school string
	}{"张三", "北京", "北京大学"}
	fmt.Printf("%v\n", people)
	fmt.Printf("%+v\n", people)
	fmt.Printf("%#v\n", people)
	// 打印Unicode
	fmt.Printf("%U\n", 'A')
	fmt.Printf("%c\n", 65)
}

// Print
func Output2Console2() {
	// 打印相邻的字符串
	fmt.Print("hello", "world", "!\n")
	fmt.Print("hello", "world", 2, 3, "!\n")
}

// Println
func Output2Console3() {
	fmt.Println("hello", "world", 2, 3, "!")
	fmt.Println("hello", "world", "!")
}

/******************** 输出到控制台 end ******************/

/******************** 输出到文件 start ******************/
func Output2File1() {
	_, _ = fmt.Fprintln(os.Stdout, "hello", "go")
	_, _ = fmt.Fprintln(os.Stdout, "hello", "php")
}
func Output2File2() {
	// 打开文件
	file, err := os.OpenFile("../tmp/test_output2file.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 写入文件
	_, err = fmt.Fprintln(file, "hello", "world")
	if err != nil {
		fmt.Println(err.Error())
	}
}

/******************** 输出到文件 end ******************/

/******************** 生成字符串 start ******************/
func GenerateStr() {
	sprintf := fmt.Sprintf("水果:%s, 价格:%.3f", "香蕉", 1.99)
	fmt.Printf("Sprintf:%s\n", sprintf)
	sprint := fmt.Sprint("I", "LOVE", "YOU")
	fmt.Printf("Sprint:%s\n", sprint)
	sprintln := fmt.Sprintln("I", "LOVE", "YOU")
	fmt.Printf("Sprintln:%s\n", sprintln)
}

/******************** 生成字符串 end ******************/

/******************** 生成错误类型 start ******************/
func GenerateError() {
	// 返回错误类型
	err := fmt.Errorf("错误信息:%s", "参数不全")
	fmt.Printf("类型:%T, 值:%v\n", err, err)
}

/******************** 生成错误类型 end ******************/

/******************** 接收控制台收入 start ******************/
func ReceiveFromConsole1() {
	var (
		name   string
		age    int
		isBody bool
	)
	_, err := fmt.Scan(&name, &age, &isBody)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("name:%s, age:%d, isBody:%t", name, age, isBody)
}

/******************** 接收控制台收入 end ******************/
