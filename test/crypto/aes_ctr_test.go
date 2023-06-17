/**
 * @Author Shershon
 * @Description 加密:AES, 模式:CTR(计算器模式), 填充:Pkcs7, 密文编码:Base64
 * @Date 2021/6/29 6:04 下午
 **/
package crypto

import (
	"fmt"
	"shershon1991/go-tools/app/cryptopkg"
	"strings"
	"testing"
)

// 测试AES-CTR加密
func TestAesEncryptByCTR(t *testing.T) {
	key := strings.Repeat("a", 16)
	data := "hello word"
	hex, base64 := cryptopkg.AesEncryptByCTR(data, key)
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
	res := cryptopkg.AesDecryptByCTR(data, key)
	fmt.Printf("解密key: %v \n", key)
	fmt.Printf("解密数据: %v \n", data)
	fmt.Printf("解密结果: %v \n", res)
}
