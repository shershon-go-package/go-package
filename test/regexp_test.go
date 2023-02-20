/**
 * @Author Shershon
 * @Description 正则表达式使用
 * @Date 2021/7/5 2:54 下午
 **/
package test

import (
	"fmt"
	"regexp"
	"testing"
)

// 查找字节b中是否存在pattern字符串
func TestMatch(t *testing.T) {
	matched, _ := regexp.Match("hello", []byte("hello word"))
	fmt.Printf("matched: %t\n", matched)
	matched2, _ := regexp.Match("go", []byte("hello word"))
	fmt.Printf("matched2: %t\n", matched2)
}

// 查找字节b中是否存在pattern字符串
func TestMatchString(t *testing.T) {
	matchString, _ := regexp.MatchString("hello", "hello,word")
	fmt.Printf("matchString: %t\n", matchString)
	matchString2, _ := regexp.MatchString("go", "hello,word")
	fmt.Printf("matchString2: %t\n", matchString2)
}

// 根据正则表达式创建对应的对象
func TestRegexp(t *testing.T) {
	compile, err := regexp.Compile("go*")
	if err != nil {
		t.Error(err)
	}
	// 查找
	find := compile.Find([]byte("hello go"))
	fmt.Printf("find: %s\n", find)
	// 替换
	all := compile.ReplaceAll([]byte("hello go"), []byte("php"))
	fmt.Printf("ReplaceAll: %s\n", all)

}
