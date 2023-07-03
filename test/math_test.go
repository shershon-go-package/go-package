/**
 * @Description
 **/
package test

import (
	"fmt"
	"shershon1991/go-package/app/mathpkg"
	"testing"
)

// 取整
func TestRound(t *testing.T) {
	mathpkg.Round()
}

// 最大值和最小值
func TestCompare(t *testing.T) {
	mathpkg.Compare()
}

// 取模和取余
func TestR(t *testing.T) {
	mathpkg.R()
}

// 随机数
func TestRand(t *testing.T) {
	mathpkg.Rand()
}

func TestOther(t *testing.T) {
	fmt.Printf("%f\n", float32(1/3))          // 0.000000
	fmt.Printf("%f\n", float32(1)/float32(3)) // 0.333333
	fmt.Println(498066 / 100)                 // 0.333333
	fmt.Println(fmt.Sprintf("%.2f", float32(498012/100)))
	fmt.Println(fmt.Sprintf("%.2f", float32(498012)/100))
	fmt.Println(fmt.Sprintf("%.2f", float32(0)/100))
}
