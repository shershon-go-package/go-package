package packageuse

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 取整
func Round() {
	// 向上取整
	fmt.Println("50.345,向上取整 -> ", math.Ceil(50.345))
	fmt.Println("50.745,向上取整 -> ", math.Ceil(50.745))
	// 向下取整
	fmt.Println("50.345,向下取整 -> ", math.Floor(50.345))
	fmt.Println("50.745,向下取整 -> ", math.Floor(50.745))
	// 四舍五入
	fmt.Println("50.345,四舍五入 -> ", math.Floor(50.345+0.5))
	fmt.Println("50.745,四舍五入 -> ", math.Floor(50.745+0.5))
}

// 最大值和最小值
func Compare() {
	a := 12.4555
	b := 12.8234
	fmt.Printf("%.4f和%.4f,最大值是:%.4f\n", a, b, math.Max(a, b))
	fmt.Printf("%f和%f,最小值是:%f\n", a, b, math.Min(a, b))
}

// 取模和取余
func R() {
	a := 20.0
	b := -3.0
	fmt.Printf("%.2f对%.2f 取模: %.2f\n", a, b, math.Mod(a, b))
	fmt.Printf("%.2f对%f 取余: %.2f\n", a, b, math.Remainder(a, b))
}

// 随机数
func Rand() {
	// 设置随机因子(需要设置成非固定值)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++ {
		fmt.Println("随机整数[0.100)", rand.Intn(100))
		fmt.Println("随机浮点数[0.0,1.0)", rand.Float64())
	}
}
