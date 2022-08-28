/**
 * @Author Shershon
 * @Description RSA数字签名
 * @Date 2021/7/2 3:01 下午
 **/
package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

// 对数据进行数字签名
func GetRSASign(data, privateKeyPath string) (string, error) {
	// 读取私钥
	privateKey, err := ReadRSAPKCS1PrivateKey(privateKeyPath)
	if err != nil {
		return "", err
	}
	// 计算Sha1散列值
	hash := sha256.New()
	hash.Write([]byte(data))
	sum := hash.Sum(nil)
	// 从1.5版本规定，使用RSASSA-PKCS1-V1_5-SIGN 方案计算签名
	signPKCS1v15, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, sum)
	// 结果转成base64
	toString := base64.StdEncoding.EncodeToString(signPKCS1v15)
	return toString, err
}

// 验证签名
func VerifyRsaSign(data, publicKeyPath, base64Sign string) (bool, error) {
	// 反解base64
	sign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return false, err
	}
	// 获取公钥
	publicKey, err := ReadRSAPublicKey(publicKeyPath)
	if err != nil {
		return false, err
	}
	// 计算Sha1散列值
	hash := sha256.New()
	hash.Write([]byte(data))
	bytes := hash.Sum(nil)
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, bytes, sign)
	return err == nil, err
}
