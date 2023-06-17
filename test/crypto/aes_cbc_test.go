/**
 * @Author Shershon
 * @Description 加密:AES, 模式:CBC(密码分组链模式), 填充:Pkcs7, 密文编码:Base64
 * @Date 2021/6/29 5:40 下午
 **/
package crypto

import (
	"fmt"
	"shershon1991/go-tools/app/cryptopkg"
	"strings"
	"testing"
)

// AES加密
func TestAesEncryptByCBC(t *testing.T) {
	key := strings.Repeat("a", 16)
	fmt.Printf("key: %v 长度: %d \n", key, len(key))
	text := "abc"
	fmt.Printf("待加密文案: %v \n", text)
	encrypt := cryptopkg.AesEncryptByCBC(text, key)
	fmt.Printf("加密结果: %v \n", encrypt)
}

// AES解密
func TestAesDecryptByCBC(t *testing.T) {
	key := strings.Repeat("a", 16)
	fmt.Printf("key: %v 长度: %d \n", key, len(key))
	text := "rMX6r9x+PnTOhfgDH4jjXg=="
	fmt.Printf("待解密文案: %v \n", text)
	decrypt := cryptopkg.AesDecryptByCBC(text, key)
	fmt.Printf("解密结果: %v \n", decrypt)
}
