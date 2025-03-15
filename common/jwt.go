package common

import (
	"gin_demo/model"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

var jwtKey = []byte("hzx secret key")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// ReleaseToken 生成 JWT 令牌
func ReleaseToken(user model.User) (string, error) {
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(7 * 24 * time.Hour)), // 过期时间 7 天)
			IssuedAt:  jwt.At(time.Now()),
			Issuer:    "gin_demo",
			Subject:   "user token",
		},
	}

	// 使用 HMAC SHA-256 签名生成 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用 `SignedString` 进行签名
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
