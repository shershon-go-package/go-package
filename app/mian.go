package main

import (
	"fmt"
)

type Test struct {

}

func (t Test) Echo(str string)  {
	fmt.Println(str,"echo。。。。")
}

func (t *Test) Hello(str string)  {
	fmt.Println(str,"hello。。。。")
}

func main() {
	var t1  = make([]int,1)
	var t2 = new(Test)
	println("t1: \n",t1)
	println("t11: \n",[]int{1,2,3,4})
	print("t2\n",t2)


}
