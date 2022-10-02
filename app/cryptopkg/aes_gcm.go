/**
 * @Author Shershon
 * @Description AES加密模式GCM-电码本模式
 * @Date 2021/6/30 4:24 下午
 **/
package cryptopkg

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// 加密(GCM 不需要补码)
func AesEncryptByGCM(data, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(fmt.Sprintf("NewCipher error:%s", err))
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(fmt.Sprintf("NewGCM error:%s", err))
	}
	// 生成随机因子(这里固定取密钥指定位数)
	//nonce := make([]byte, gcm.NonceSize())
	//if _,err := io.ReadFull(rand.Reader,nonce); err != nil {
	//	panic(fmt.Sprintf("make rand nonce error:%s", err))
	//}
	nonceStr := key[:gcm.NonceSize()]
	nonce := []byte(nonceStr)
	fmt.Printf("nonceStr = %v \n", nonceStr)
	seal := gcm.Seal(nonce, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(seal)
}

// 解密(GCM 不需要解码)
func AesDecryptByGCM(data, key string) string {
	// 反解base64
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(fmt.Sprintf("base64 DecodeString error:%s", err))
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(fmt.Sprintf("NewCipher error:%s", err))
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(fmt.Sprintf("NewGCM error:%s", err))
	}
	nonceSize := gcm.NonceSize()
	if len(dataByte) < nonceSize {
		panic("dataByte to short")
	}
	nonce, ciphertext := dataByte[:nonceSize], dataByte[nonceSize:]
	open, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(fmt.Sprintf("gcm Open error:%s", err))
	}
	return string(open)
}
