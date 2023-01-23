package common

import (
	"time"

	"github.com/RaymondCode/simple-demo/repository"
	"github.com/dgrijalva/jwt-go"
)

/* common/jwt.go */
// jwt加密密钥
var jwtKey = []byte("1234jkhj1k2h*")

const ToeknPrefix = "douyin-"

type Claims struct {
	UserId uint64
	jwt.StandardClaims
}

// ReleaseToken 生成token
func ReleaseToken(user repository.User) (string, error) {
	// token的有效期
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		// 自定义字段
		UserId: user.Id,
		// 标准字段
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expirationTime.Unix(),
			// 发放时间
			IssuedAt: time.Now().Unix(),
		},
	}
	// 使用jwt密钥生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	// 返回token
	return ToeknPrefix + tokenString, nil
}

// resolve token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
