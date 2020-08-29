package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

type Claims struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

type Token struct {
	Token string `json:"token"`
}

// Claims 的默认字段
// iss：JWT Token 的签发者
// sub：主题
// exp：JWT Token 过期时间
// aud：接收 JWT Token 的一方
// iat：JWT Token 签发时间
// nbf：JWT Token 生效时间
// jti：JWT Token ID
func Sign(ctx *gin.Context, c Claims, secret string) (tokenString string, err error) {
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}

	// the token content
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),             // token 的生效使时间
		"iat":      time.Now().Unix(),             // token 的签发时间
		"exp":      time.Now().Add(3 * time.Hour), // token 的过期时间
	})

	// sign the token with the specified secret
	tokenString, err = token.SignedString([]byte(secret))

	return
}

// 解析token
func Parse(tokenString, secret string) (*Claims, error) {
	c := &Claims{}

	// parse token
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return c, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.ID = uint64(claims["id"].(float64))
		c.Username = claims["username"].(string)
		return c, nil
	} else {
		return c, err
	}
}

// secretFunc 验证密码格式
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// 确认 alg 是预期的算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// 从请求头解析token
func ParseRequest(c *gin.Context) (*Claims, error) {
	header := c.Request.Header.Get("Authorization")

	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Claims{}, ErrMissingHeader
	}

	var t string
	fmt.Sscanf(header, "Bearer %s", &t)

	return Parse(t, secret)

}
