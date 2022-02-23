package packageuse

import (
	"fmt"
	"strconv"
)

func ParseFloat() {
	f, err := strconv.ParseFloat("9223372036854775807", 64)
	fmt.Printf("f: %+v err:%+v\n", f, err)
}
