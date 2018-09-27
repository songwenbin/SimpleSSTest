package loginserver

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
)

type SafeRsa struct {
	privateKey []byte
	publicKey  []byte
}

func NewSafeRsa(public, private []byte) *SafeRsa {
	return &SafeRsa{
		publicKey:  public,
		privateKey: private,
	}
}

func (sr *SafeRsa) Get() string {
	return string(sr.publicKey)
}

func (sr *SafeRsa) RsaEncrypt(data []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(sr.publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

func (sr *SafeRsa) RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//获取私钥
	block, _ := pem.Decode(sr.privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func SafeEncoder(sr *SafeRsa, clearContent string) string {
	fmt.Println(clearContent)
	// ECDSA_Encrypt_SEP256_K1
	enc, err := sr.RsaEncrypt([]byte(clearContent))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Encrypted:", string(enc))
	// Base64URLEncoded
	encodeString := base64.StdEncoding.EncodeToString(enc)
	fmt.Println(encodeString)

	return encodeString
}

func SafeDecoder(sr *SafeRsa, params string) string {
	// Base64URLEncoded
	decodeBytes, err := base64.StdEncoding.DecodeString(params)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decodeBytes))
	// ECDSA_Encrypt_SEP256_K1
	decstr, _ := sr.RsaDecrypt(decodeBytes)
	fmt.Println("Decrypted:", string(decstr))

	return string(decstr)
}
