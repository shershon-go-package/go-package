/**
 * @Author Shershon
 * @Description 生成rsa私钥和密钥对
 * @Date 2021/7/1 11:10 上午
 **/
package cryptopkg

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// 保存私钥公钥目录
var savKeyPath string

// 定义结构体保存私钥和公钥
type KeySaveInfo struct {
	PrivateKeyPem string
	PublicKeyPem  string
}

// 生成密钥对(格式是: PKCS1)
func GenerateRSAPKCS1Key(size int, savePath string) (KeySaveInfo, error) {
	// 设置保存路径
	savKeyPath = savePath
	// 声明保存结果
	var keySaveInfo KeySaveInfo
	// 使用随机数据生成器random,生成一对具有指定字位数的RSA密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return keySaveInfo, err
	}
	// 判断bits位数长度
	if bitLen := privateKey.N.BitLen(); bitLen != size {
		return keySaveInfo, fmt.Errorf("size too short size=%d ,bit=%d", size, bitLen)
	}
	// 保存私钥
	err = savePrivateKeyPKCS1(privateKey, &keySaveInfo)
	if err != nil {
		return keySaveInfo, err
	}
	// 保存公钥
	err = savePublicKeyPKCS1(privateKey, &keySaveInfo)
	if err != nil {
		return keySaveInfo, err
	}
	return keySaveInfo, nil
}

// 私钥以pem格式保存到文件
func savePrivateKeyPKCS1(privateKey *rsa.PrivateKey, keySaveInfo *KeySaveInfo) error {
	// MarshalPKCS1PrivateKey 将 RSA 私钥转换为 PKCS1
	x509PriKey := x509.MarshalPKCS1PrivateKey(privateKey)
	// 创建私钥文件，后缀为pem
	privateFile := fmt.Sprintf("%s/%s", savKeyPath, "private_ssl.pem")
	keySaveInfo.PrivateKeyPem = privateFile
	fileHandle, err := os.Create(privateFile)
	if err != nil {
		return err
	}
	defer fileHandle.Close()
	// 创建pem.Block结构体
	pemBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509PriKey}
	// 保存到文件
	return pem.Encode(fileHandle, &pemBlock)
}

// 公钥以pem格式保存到文件
func savePublicKeyPKCS1(privateKey *rsa.PrivateKey, keySaveInfo *KeySaveInfo) error {
	// MarshalPKCS1PrivateKey 将 RSA 私钥转换为 PKCS1
	x509PublicKey := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	// 创建公钥文件
	publicKeyFile := fmt.Sprintf("%s/%s", savKeyPath, "public_ssl.pem")
	keySaveInfo.PublicKeyPem = publicKeyFile
	fileHandle, err := os.Create(publicKeyFile)
	if err != nil {
		return err
	}
	defer fileHandle.Close()
	// 创建pen.Block
	pemBlock := pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509PublicKey}
	// 保存到文件
	return pem.Encode(fileHandle, &pemBlock)
}
