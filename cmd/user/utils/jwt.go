package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 设置jwt密钥
const (
	JwtKey = "CI6MTYxODQ3ODIwNiwiaWF0I"
)

// jwt声明结构体（即jwt的字段信息）
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//颁发token
func GetToken(user uint) string {
	expireTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		UserId: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(), // 颁发时间
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", // 签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		fmt.Println(err, "err")
	}
	return tokenString
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(JwtKey), nil
	})
	return token, Claims, err
}
