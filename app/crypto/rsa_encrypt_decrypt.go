/**
 * @Author Mr.LiuQH
 * @Description RSA加解密
 * @Date 2021/7/2 3:33 下午
 **/
package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
)

// 加密(使用公钥加密)
func RSAEncrypt(data, publicKeyPath string) (string, error) {
	// 获取公钥
	rsaPublicKey, err := ReadRSAPublicKey(publicKeyPath)
	if err != nil {
		return "", err
	}
	// 加密
	encryptPKCS1v15, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(data))
	if err != nil {
		return "",err
	}
	// 把加密结果转成Base64
	encryptString := base64.StdEncoding.EncodeToString(encryptPKCS1v15)
	return encryptString, err
}

// 解密(使用私钥解密)
func RSADecrypt(base64data,privateKeyPath string) (string,error) {
	// data反解base64
	decodeString, err := base64.StdEncoding.DecodeString(base64data)
	if err != nil {
		return "", err
	}
	// 读取密钥
	rsaPrivateKey, err := ReadRSAPKCS1PrivateKey(privateKeyPath)
	if err != nil {
		return "", err
	}
	// 解密
	decryptPKCS1v15, err := rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, decodeString)
	return string(decryptPKCS1v15),err
}