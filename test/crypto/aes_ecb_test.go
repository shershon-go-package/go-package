/**
 * @Author Shershon
 * @Description 加密:AES, 模式:ECB(电码本模式), 填充:Pkcs7, 密文编码:Base64
 * @Date 2021/6/29 5:44 下午
 **/
package crypto

import (
	"fmt"
	"shershon1991/go-package/app/cryptopkg"
	"strings"
	"testing"
)

// 加密
func TestECBEncrypt(t *testing.T) {
	key := strings.Repeat("a", 16)
	data := "hello word"
	s := cryptopkg.AesEncryptByECB(data, key)
	fmt.Printf("加密密钥: %v \n", key)
	fmt.Printf("加密数据: %v \n", data)
	fmt.Printf("加密结果: %v \n", s)
}

// 解密
func TestECBDecrypt(t *testing.T) {
	key := strings.Repeat("a", 16)
	data := "mMAsLF/fPBfUrP0mPqZm1w=="
	s := cryptopkg.AesDecryptByECB(data, key)
	fmt.Printf("解密密钥: %v \n", key)
	fmt.Printf("解密数据: %v \n", data)
	fmt.Printf("解密结果: %v \n", s)
}
