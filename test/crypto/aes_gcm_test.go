/**
 * @Author Mr.LiuQH
 * @Description AES-GCM 测试使用
 * @Date 2021/6/30 4:59 下午
 **/
package crypto

import (
	"52lu/go-study-example/app/crypto"
	"fmt"
	"strings"
	"testing"
)

func TestAesGCM(t *testing.T) {
	key := strings.Repeat("a",16)
	data := "hello word!"
	// 加密
	gcm := crypto.AesEncryptByGCM(data, key)
	fmt.Printf("密钥key: %s \n",key)
	fmt.Printf("加密数据: %s \n",data)
	fmt.Printf("加密结果: %s \n",gcm)
	// 解密
	byGCM := crypto.AesDecryptByGCM(gcm, key)
	fmt.Printf("解密结果: %s \n",byGCM)
}