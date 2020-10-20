package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func GenRsaKey(bits int, filePath string) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:    "RSA PRIVATE KEY",
		Bytes:   derStream,
	}
	// 创建文件
	err = ioutil.WriteFile(filePath + "/private.pem", pem.EncodeToMemory(block), 0644)
	if err != nil {
		return err
	}
	//file, err := os.Create(filePath + "/private.pem")
	//defer file.Close()
	//if err != nil {
	//	return err
	//}
	//err = pem.Encode(file, block)
	//if err != nil {
	//	return err
	//}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:    "RSA PUBLIC KEY",
		Bytes:   derPkix,
	}
	err = ioutil.WriteFile(filePath + "/public.pem", pem.EncodeToMemory(block), 0644)
	if err != nil {
		return err
	}
	//file, err = os.Create(filePath + "/public.pem")
	//if err != nil {
	//	return err
	//}
	//err = pem.Encode(file, block)
	//if err != nil {
	//	return err
	//}

	return nil
}
