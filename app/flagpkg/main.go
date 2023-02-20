/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/02/20 11:00
 */

package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	// 接收字符串
	str = flag.String("name", "默认名", "用户姓名")
	// 接收整型
	age = flag.Int("age", 18, "用户年龄")
	// 接收布尔型
	smoking = flag.Bool("smoking", false, "是否吸烟")
)

/*func main1() {
	// 解析
	flag.Parse()
	fmt.Printf("str: %+v \n", *str)
	fmt.Printf("age: %+v \n", *age)
	fmt.Printf("smoking: %+v \n", *smoking)
}
go run main.go -name=张三 -age=100 -smoking=true
*/

var (
	name2    string
	age2     int
	smoking2 bool
	weight2  float64
)

/*func main2() {
	// 接收字符串
	flag.StringVar(&name2, "name2", "默认名", "用户姓名")
	// 接收整型
	flag.IntVar(&age2, "age2", 18, "用户年龄")
	// 接收布尔类型
	flag.BoolVar(&smoking2, "smoking2", false, "是否吸烟")
	// 接收浮点型
	flag.Float64Var(&weight2, "weight2", 80.0, "体重")

	flag.Parse()
	fmt.Printf("name2: %+v \n", name2)
	fmt.Printf("age2: %+v \n", age2)
	fmt.Printf("smoking2: %+v \n", smoking2)
	fmt.Printf("weight2: %+v \n", weight2)
}
go run main.go -name2=张三 -age2=22 -smoking2 true -weight2=88.9
*/

// 自定义变量，并实现flag.Value接口
type likes []string

func (l *likes) String() string {
	return fmt.Sprintf("%v", *l)
}
func (l *likes) Set(s string) error {
	split := strings.Split(s, ",")
	*l = split
	return nil
}

/*func main3() {
	var likeList likes
	// 接收自定义类型
	flag.Var(&likeList, "likes", "接收自定义类型")
	flag.Parse()
	fmt.Println(likeList)
}
go run main.go -likes=篮球,足球,游戏
*/

/*func main4() {
	// 注意Parse是在Args之前调用
	flag.Parse()
	// 一次接收所有的参数
	args := flag.Args()
	fmt.Println(args)
}
go run main.go 张三 18 男
*/

/*func main() {
	// 注意Parse是在Arg之前调用
	flag.Parse()
	// 获取指定索引位置参数
	p0 := flag.Arg(0)
	p1 := flag.Arg(1)
	p2 := flag.Arg(2)
	fmt.Printf("索引=0,v=%v \n", p0)
	fmt.Printf("索引=1,v=%v \n", p1)
	fmt.Printf("索引=2,v=%v \n", p2)
}
go run main.go 张三 18 男*/
