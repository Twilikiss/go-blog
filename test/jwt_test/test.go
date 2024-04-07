/*
测试验证jwt的相关操作
*/
package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	SECRET_KEY = "243223ffslsfsldfl412fdsfsdf" //私钥
)

// CustomClaims 自定义Claims
type CustomClaims struct {
	UserId int64
	jwt.RegisteredClaims
}

func main() {
	//生成token
	maxAge := 60 * 60 * 24
	customClaims := &CustomClaims{
		UserId: 11, //用户id
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        time.Now().String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(maxAge) * time.Second)), // 过期时间，必须设置
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Elysia", // 非必须，也可以填充用户名，
		},
	}
	//采用HMAC SHA256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("token: %v\n", tokenString)

	//解析token
	ret, err := ParseToken(tokenString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("userinfo: %v\n", ret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
