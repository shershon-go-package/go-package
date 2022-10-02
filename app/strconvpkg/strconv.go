package strconvpkg

import (
	"fmt"
	"strconv"
)

func ParseFloat() {
	f, err := strconv.ParseFloat("9223372036854775807", 64)
	fmt.Printf("f: %+v err:%+v\n", f, err)
}

func ParseInt() {
	a := "100"
	res, _ := strconv.ParseInt(a, 10, 64)
	fmt.Printf("t:%T, v:%v \n", res, res)
}
