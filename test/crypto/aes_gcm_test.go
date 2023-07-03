/**
 * @Author Shershon
 * @Description 加密:AES, 模式:GCM, 密文编码:Base64
 * @Date 2021/6/30 4:59 下午
 **/
package crypto

import (
	"fmt"
	"shershon1991/go-package/app/cryptopkg"
	"strings"
	"testing"
)

func TestAesGCM(t *testing.T) {
	key := strings.Repeat("a", 16)
	data := "hello word!"
	// 加密
	gcm := cryptopkg.AesEncryptByGCM(data, key)
	fmt.Printf("密钥key: %s \n", key)
	fmt.Printf("加密数据: %s \n", data)
	fmt.Printf("加密结果: %s \n", gcm)
	// 解密
	byGCM := cryptopkg.AesDecryptByGCM(gcm, key)
	fmt.Printf("解密结果: %s \n", byGCM)
}
