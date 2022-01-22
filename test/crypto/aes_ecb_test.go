/**
 * @Author Mr.LiuQH
 * @Description AES加密模式ECB测试使用
 * @Date 2021/6/29 5:44 下午
 **/
package crypto

import (
	"52lu/go-study-example/app/crypto"
	"fmt"
	"strings"
	"testing"
)
// 加密
func TestECBEncrypt(t *testing.T) {
	key := strings.Repeat("a", 16)
	data := "hello word"
	s := crypto.AesEncryptByECB(data, key)
	fmt.Printf("加密密钥: %v \n", key)
	fmt.Printf("加密数据: %v \n", data)
	fmt.Printf("加密结果: %v \n", s)
}
// 解密
func TestECBDecrypt(t *testing.T) {
	key := strings.Repeat("a", 16)
	data := "mMAsLF/fPBfUrP0mPqZm1w=="
	s := crypto.AesDecryptByECB(data, key)
	fmt.Printf("解密密钥: %v \n", key)
	fmt.Printf("解密数据: %v \n", data)
	fmt.Printf("解密结果: %v \n", s)
}
