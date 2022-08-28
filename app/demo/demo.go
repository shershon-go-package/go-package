package demo

import "fmt"

func init() {
	fmt.Println("init执行")
}

const name = "张三"

var age = 10

func Run() {
	fmt.Println("name=", name)
	fmt.Println("age=", age)
}
