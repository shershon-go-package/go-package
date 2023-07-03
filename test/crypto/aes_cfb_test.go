/**
 * @Author Shershon
 * @Description 加密:AES, 模式:CFB, 填充:Pkcs7, 偏移量:默认为秘钥, 密文编码:Base64
 * @Date 2021/6/29 6:04 下午
 **/
package crypto

import (
	"fmt"
	"shershon1991/go-package/app/cryptopkg"
	"strings"
	"testing"
)

// 测试AES-OFB加密
func TestAesEncryptByCFB(t *testing.T) {
	key := strings.Repeat("a", 16)
	data := "hello go !"
	_, base643 := cryptopkg.AesEncryptByCFB(data, key)
	fmt.Printf("加密key: %v \n", key)
	fmt.Printf("加密key长度: %v \n", len(key))
	fmt.Printf("加密数据: %v \n", data)
	fmt.Printf("加密结果(CFB): %v \n", base643)
}

// 测试AES-CTR解密
func TestAesDecryptByCFB(t *testing.T) {
	key := strings.Repeat("a", 16)
	data := "Oe2qKyQC69LyY+8UWLgbVQ=="
	res1 := cryptopkg.AesDecryptByCFB(data, key)
	fmt.Printf("解密key: %v \n", key)
	fmt.Printf("解密数据: %v \n", data)
	fmt.Printf("解密结果(CFB): %v \n", res1)
}
