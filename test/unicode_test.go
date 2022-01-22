/**
 * @Author Mr.LiuQH
 * @Description unicode 包测试使用
 * @Date 2021/6/17 5:02 下午
 **/
package test

import (
	"fmt"
	"testing"
	"unicode"
)

// 判断方法使用
func TestJudge(t *testing.T) {
	str1 := '刘'
	str2 := 'l'
	str3 := 'W'
	str4 := '!'
	// 是否为空格
	//rune1 := ' '
	//rune2 := 'h'
	//fmt.Printf("[%c] 是空格? %t\n", rune1, unicode.IsSpace(rune1))
	//fmt.Printf("[%c] 是空格? %t\n", rune2, unicode.IsSpace(rune2))
	//
	//// 是否为十进制
	//d1 := '1'
	//fmt.Printf("[%c] 是十进制数? %t\n", d1, unicode.IsDigit(d1))
	//d2 := '张'
	//fmt.Printf("[%c] 是十进制数? %t\n", d2, unicode.IsDigit(d2))
	//
	//// 是否为数字
	//fmt.Printf("[%c] 是数字? %t\n", d1, unicode.IsNumber(d1))
	//fmt.Printf("[%c] 是数字? %t\n", d2, unicode.IsNumber(d2))
	//fmt.Printf("[%c] 是数字? %t\n", 'L', unicode.IsNumber('L'))
	//fmt.Printf("[%c] 是数字? %t\n", str1, unicode.IsNumber(str1))

	// 是否为字母字符
	//fmt.Printf("[%c] 是字母? %t\n",str1,unicode.IsLetter(str1))
	//fmt.Printf("[%c] 是字母? %t\n",str2,unicode.IsLetter(str2))
	//fmt.Printf("[%c] 是字母? %t\n",str3,unicode.IsLetter(str3))
	//fmt.Printf("[%c] 是字母? %t\n",str4,unicode.IsLetter(str4))

	// 是否为标点符号
	//fmt.Printf("[%c] 是标点符号? %t\n", str1, unicode.IsPunct(str1))
	//fmt.Printf("[%c] 是标点符号? %t\n", str2, unicode.IsPunct(str2))
	//fmt.Printf("[%c] 是图形字符? %t\n",str3,unicode.IsGraphic(str3))
	//fmt.Printf("[%c] 是标点符号? %t\n", str4, unicode.IsPunct(str4))

	// 是否为小写字母
	//fmt.Printf("[%c] 是小写字母? %t\n", str1, unicode.IsLower(str1))
	//fmt.Printf("[%c] 是小写字母? %t\n", str2, unicode.IsLower(str2))
	//fmt.Printf("[%c] 是小写字母? %t\n", str3, unicode.IsLower(str3))
	//fmt.Printf("[%c] 是小写字母? %t\n", str4, unicode.IsLower(str4))
	// 是否为大写字母
	//fmt.Printf("[%c] 是大写字母? %t\n", str1, unicode.IsUpper(str1))
	//fmt.Printf("[%c] 是大写字母? %t\n", str2, unicode.IsUpper(str2))
	//fmt.Printf("[%c] 是大写字母? %t\n", str3, unicode.IsUpper(str3))
	//fmt.Printf("[%c] 是大写字母? %t\n", str4, unicode.IsUpper(str4))

	// 是否为汉字
	fmt.Printf("[%c] 是汉字? %t\n", str1, unicode.Is(unicode.Scripts["Han"], str1))
	fmt.Printf("[%c] 是汉字? %t\n", str2, unicode.Is(unicode.Scripts["Han"], str2))
	fmt.Printf("[%c] 是汉字? %t\n", str3, unicode.Is(unicode.Scripts["Han"], str3))
	fmt.Printf("[%c] 是汉字? %t\n", str4, unicode.Is(unicode.Scripts["Han"], str4))
}

func TestToType(t *testing.T) {
	str1 := 'W'
	str2 := 'a'
	fmt.Printf("[%c] 转成小写: %c \n", str1, unicode.ToLower(str1))
	fmt.Printf("[%c] 转成大写: %c \n", str2, unicode.ToUpper(str2))
}
