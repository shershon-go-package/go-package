/**
 * @Description TODO
 **/
package test

import (
	"fmt"
	"shershon1991/go-study-example/app/packageuse"
	"testing"
)

// 取整
func TestRound(t *testing.T) {
	packageuse.Round()
}

// 最大值和最小值
func TestCompare(t *testing.T) {
	packageuse.Compare()
}

// 取模和取余
func TestR(t *testing.T) {
	packageuse.R()
}

// 随机数
func TestRand(t *testing.T) {
	packageuse.Rand()
}

func TestOther(t *testing.T) {
	fmt.Printf("%f\n", float32(1/3))          // 0.000000
	fmt.Printf("%f\n", float32(1)/float32(3)) // 0.333333
	fmt.Println(498066 / 100)                 // 0.333333
	fmt.Println(fmt.Sprintf("%.2f", float32(498012/100)))
	fmt.Println(fmt.Sprintf("%.2f", float32(498012)/100))
	fmt.Println(fmt.Sprintf("%.2f", float32(0)/100))
}
