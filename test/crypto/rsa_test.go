/**
 * @Author Shershon
 * @Description RSA测试使用
 * @Date 2021/7/1 4:40 下午
 **/
package crypto

import (
	"fmt"
	"shershon1991/go-standard-package/app/cryptopkg"
	"testing"
)

// 测试生成密钥对
func TestGenerateKey(t *testing.T) {
	key, err := cryptopkg.GenerateRSAPKCS1Key(1024, "../../tmp")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", key)
}

// 读取密钥
func TestReadKey(t *testing.T) {
	// pkcs1格式-私钥
	privatePKCS1KeyPath := "../../tmp/private_ssl.pem"
	privatePKCS1Key, err := cryptopkg.ReadRSAPKCS1PrivateKey(privatePKCS1KeyPath)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("PKCS1私钥: %#v\n", privatePKCS1Key)

	// pkcs8格式-公钥
	publicPKCS8KeyPath := "../../tmp/public_ssl.pem"
	publicPKCS8Key, err := cryptopkg.ReadRSAPublicKey(publicPKCS8KeyPath)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("PKCS8公钥: %#v\n", publicPKCS8Key)
}

// 加密测试
func TestRsaEncrypt(t *testing.T) {
	publicKeyPath := "../../tmp/public_ssl.pem"
	data := "123456"
	encrypt, err := cryptopkg.RSAEncrypt(data, publicKeyPath)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("加密结果:%v \n", encrypt)
}

// 解密测试
func TestRsaDecrypt(t *testing.T) {
	privateKeyPath := "../../tmp/private_ssl.pem"
	data := "pUYa4set6XkBshfio5g2hzPx1tA67sxEvJBpJiuK3McJ9cPJAXzuRkWIy4s6cDQOhrPUaNXhr3M3WLHH19/eaqcNZz1yOFZwgGKmkWtdmygtLB/wrDant9uRfXrvzlV9iMq+cUlqsrwuCa0wcGEBNHRhIJOQSTs+SxaRTeoRCbU="
	encrypt, err := cryptopkg.RSADecrypt(data, privateKeyPath)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("解密结果:%v \n", encrypt)
}

// 数据加签
func TestAddSign(t *testing.T) {
	privateKeyPath := "../../tmp/private_ssl.pem"
	data := "123456"
	sign, err := cryptopkg.GetRSASign(data, privateKeyPath)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("数据签名: %v \n", sign)
}

// 数据签名验证
func TestVaSign(t *testing.T) {
	publicKeyPath := "../../tmp/public_ssl.pem"
	data := "123456"
	sign := "QnGqGbIqoHjJG1l+JiaOKWBdX+h00lnKCoO2rTYKIro9hoaDj7nqmu+Mxsuo+2jumicvCNBZNOpMzYryjZf0x7Q4ycLBtqtCWuFRasiInUO7Avy19LRTjdMf2xw9968vilB/xEAQ53JXIDUVvCsMxTfpHI9oRiWEGXWNkhfkjkQ="
	verifyRsaSign, err := cryptopkg.VerifyRsaSign(data, publicKeyPath, sign)
	if err != nil {
		fmt.Printf("验签失败: %v \n", err)
	}
	fmt.Printf("验签结果: %v \n", verifyRsaSign)
}
