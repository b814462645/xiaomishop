package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key") // 这是你的JWT签名密钥

// 自定义声明
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成JWT Token
func GenerateJWT(username string) (string, error) {
	// 设置Token过期时间为1小时
	expirationTime := time.Now().Add(1 * time.Hour)

	// 创建JWT声明，包括用户名和过期时间
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 使用HS256算法创建Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名Token并获取完整的编码后的字符串
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析并验证JWT Token
func ParseJWT(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	// 解析Token，传入密钥验证签名
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
