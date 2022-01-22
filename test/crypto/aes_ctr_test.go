/**
 * @Author Mr.LiuQH
 * @Description AES加密模式CTR-计算器模式-测试使用
 * @Date 2021/6/29 6:04 下午
 **/
package crypto

import (
	"52lu/go-study-example/app/crypto"
	"fmt"
	"strings"
	"testing"
)

// 测试AES-CTR加密
func TestAesEncryptByCTR(t *testing.T) {
	key := strings.Repeat("a", 16)
	data := "hello word"
	hex, base64 := crypto.AesEncryptByCTR(data, key)
	fmt.Printf("加密key: %v \n", key)
	fmt.Printf("加密key长度: %v \n", len(key))
	fmt.Printf("加密数据: %v \n", data)
	fmt.Printf("加密结果(hex): %v \n", hex)
	fmt.Printf("加密结果(base64): %v \n", base64)
}
// 测试AES-CTR解密
func TestAesDecryptByCTR(t *testing.T) {
	key := strings.Repeat("a", 16)
	data := "Oe2qKyQC+9KgJu8UWLgbVQ=="
	res := crypto.AesDecryptByCTR(data, key)
	fmt.Printf("解密key: %v \n", key)
	fmt.Printf("解密数据: %v \n", data)
	fmt.Printf("解密结果: %v \n", res)
}