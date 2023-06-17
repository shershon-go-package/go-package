/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/05/17 14:31
 */

package funkpkg

import "fmt"
import "github.com/thoas/go-funk"

func Exist() {
	// 判断任意类型
	fmt.Println("str->", funk.Contains([]string{"a", "b"}, "a"))
	// int 类型
	fmt.Println("int->", funk.ContainsInt([]int{100, 200}, 300))
}

func IndexOf() {
	strArr := []string{"php", "go", "java", "c++"}
	// 具体类型
	fmt.Println("c++: ", funk.IndexOfString(strArr, "c++"))
	// 验证第一次出现位置
	fmt.Println("java: ", funk.IndexOfString(strArr, "java"))
	// 任意类型
	fmt.Println("go: ", funk.IndexOf(strArr, "go"))
	// 不存在时返回-1
	fmt.Println("c: ", funk.IndexOf(strArr, "c"))
}

func LastIndexOf() {
	strArr := []string{"php", "go", "java", "c++", "python", "c++", "go"}
	// 具体类型
	fmt.Println("c++: ", funk.LastIndexOfString(strArr, "c++"))
	// 验证第一次出现位置
	fmt.Println("java: ", funk.LastIndexOfString(strArr, "java"))
	// 任意类型
	fmt.Println("go: ", funk.LastIndexOf(strArr, "go"))
	// 不存在时返回-1
	fmt.Println("c: ", funk.LastIndexOf(strArr, "c"))
}

func Every() {
	strArr := []string{"go", "java", "c", "python"}
	fmt.Println("都存在: ", funk.Every(strArr, "go", "java"))
	fmt.Println("有一个存在: ", funk.Every(strArr, "php", "java"))
	fmt.Println("都不存在: ", funk.Every(strArr, "php", "c++"))
}

func Some() {
	strArr := []string{"go", "java", "c", "python"}
	fmt.Println("都存在: ", funk.Some(strArr, "go", "java"))
	fmt.Println("有一个存在: ", funk.Some(strArr, "php", "java"))
	fmt.Println("都不存在: ", funk.Some(strArr, "php", "c++"))
}

func LastOrFirst() {
	number := []int{10, 12, 23, 30}
	// 获取第一个元素
	fmt.Println("Head: ", funk.Head(number))
	// 获取最后一个元素
	fmt.Println("Last: ", funk.Last(number))
}

type Student struct {
	Name string
	Age  int
}

func Fill() {
	// 初始化切片
	var data = make([]int, 3)
	fill, _ := funk.Fill(data, 100)
	fmt.Printf("fill: %v \n", fill)
	// 将所有值设置成2
	input := []int{1, 2, 3}
	result, _ := funk.Fill(input, 2)
	fmt.Printf("result: %v \n", result)

	var studentData = make([]Student, 2)
	studentInfo, _ := funk.Fill(studentData, Student{"张三", 30})
	fmt.Printf("studentInfo: %v \n", studentInfo)
}

type cus struct {
	Name string
	Age  int
	Home string
}

func Join() {
	a := []int64{1, 3, 5, 7}
	b := []int64{3, 7}
	// 任意类型切片交集
	join := funk.Join(a, b, funk.InnerJoin)
	fmt.Println("join: ", join)
	// 指定类型取交集
	joinInt64 := funk.JoinInt64(a, b, funk.InnerJoinInt64)
	fmt.Println("joinInt64: ", joinInt64)
	// 自定义结构体交集
	sliceA := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "南京"},
	}
	sliceB := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "上海"},
	}
	result := funk.Join(sliceA, sliceB, funk.InnerJoin)
	fmt.Println("result: ", result)
}

func DiffSlice() {
	a := []int64{1, 3, 5, 7}
	b := []int64{3, 7, 10}
	// 任意类型切片取差集
	join := funk.Join(a, b, funk.OuterJoin)
	fmt.Println("join: ", join)
	// 指定类型取差集
	joinInt64 := funk.JoinInt64(a, b, funk.OuterJoinInt64)
	fmt.Println("joinInt64: ", joinInt64)
	// 自定义结构体差集
	sliceA := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "南京"},
	}
	sliceB := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "上海"},
	}
	result := funk.Join(sliceA, sliceB, funk.OuterJoin)
	fmt.Println("result: ", result)
}
