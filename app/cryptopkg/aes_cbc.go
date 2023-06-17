/**
 * @Author Shershon
 * @Description 加密:AES, 模式:CBC(密码分组链模式), 填充:Pkcs7, 密文编码:Base64
 * @Date 2021/6/29 5:35 下午
 **/
package cryptopkg

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AES加密
func AesEncryptByCBC(str, key string) string {
	// 判断key长度
	keyLenMap := map[int]struct{}{16: {}, 24: {}, 32: {}}
	if _, ok := keyLenMap[len(key)]; !ok {
		panic(any("key长度必须是 16、24、32 其中一个"))
	}
	// 待加密字符串转成byte
	originDataByte := []byte(str)
	// 秘钥转成[]byte
	keyByte := []byte(key)
	// 创建一个cipher.Block接口。参数key为密钥，长度只能是16、24、32字节
	block, _ := aes.NewCipher(keyByte)
	// 获取秘钥长度
	blockSize := block.BlockSize()
	// 补码填充
	originDataByte = PKCS7Padding(originDataByte, blockSize)
	// 选用加密模式
	blockMode := cipher.NewCBCEncrypter(block, keyByte[:blockSize])
	// 创建数组，存储加密结果
	encrypted := make([]byte, len(originDataByte))
	// 加密
	blockMode.CryptBlocks(encrypted, originDataByte)
	// []byte转成base64
	return base64.StdEncoding.EncodeToString(encrypted)
}

// 解密
func AesDecryptByCBC(encrypted, key string) string {
	// 判断key长度
	keyLenMap := map[int]struct{}{16: {}, 24: {}, 32: {}}
	if _, ok := keyLenMap[len(key)]; !ok {
		panic(any("key长度必须是 16、24、32 其中一个"))
	}
	// encrypted密文反解base64
	decodeString, _ := base64.StdEncoding.DecodeString(encrypted)
	// key 转[]byte
	keyByte := []byte(key)
	// 创建一个cipher.Block接口。参数key为密钥，长度只能是16、24、32字节
	block, _ := aes.NewCipher(keyByte)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 选择加密模式
	blockMode := cipher.NewCBCDecrypter(block, keyByte[:blockSize])
	// 创建数组，存储解密结果
	decodeResult := make([]byte, blockSize)
	// 解密
	blockMode.CryptBlocks(decodeResult, decodeString)
	// 解码
	padding := PKCS7UNPadding(decodeResult)
	return string(padding)
}
