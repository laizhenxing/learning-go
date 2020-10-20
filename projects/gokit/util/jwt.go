package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	Name string `json:"username"`
	jwt.StandardClaims
}

var JwtSecret = []byte("jwtSecret01")

func GenerateToken(name string) (string, error) {
	uc := UserClaim{
		Name:           name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 60).Unix(),	// 设置60s过期
		},
	}

	// 读取私钥文件
	privateKeyBytes, err := ioutil.ReadFile("pems/private.pem")
	if err != nil {
		return "", err
	}
	// 使用jwt解析私钥
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)

	if err != nil {
		return "", err
	}
	tokenOjb := jwt.NewWithClaims(jwt.SigningMethodRS256, uc)
	token, err := tokenOjb.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(tokenStr string) error {
	publicKeyBytes, err := ioutil.ReadFile("pems/public.pem")
	if err != nil {
		return err
	}
	// 使用 jwt 解析公钥
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return err
	}
	token, err  := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return err
	}
	if token.Valid {
		fmt.Println(token.Claims)
	}
	return nil
}

func ParseTokenWithClaims(tokenStr string) (interface{}, error) {
	publicKeyBytes, err := ioutil.ReadFile("pems/public.pem")
	if err != nil {
		return nil, err
	}
	// 使用 jwt 解析公钥
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无法解析")
}
