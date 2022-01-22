/**
 * @Author Mr.LiuQH
 * @Description aes共用方法
 * @Date 2021/6/29 7:37 下午
 **/
package crypto

import "bytes"

// 补码
func PKCS7Padding(originByte []byte, blockSize int) []byte {
	// 计算补码长度
	padding := blockSize - len(originByte)%blockSize
	// 生成补码
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 追加补码
	return append(originByte, padText...)
}

// 解码
func PKCS7UNPadding(originDataByte []byte) []byte {
	length := len(originDataByte)
	unpadding := int(originDataByte[length-1])
	return originDataByte[:(length - unpadding)]
}
