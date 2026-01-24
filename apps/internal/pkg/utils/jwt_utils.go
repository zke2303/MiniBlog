// Package utils 工具类
package utils

import (
	"time"

	"mini-blog/internal/pkg/errmsg"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims 自定义的 Claims
type CustomClaims struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GeneratorJwt 生成 jwt 令牌
func GeneratorJwt(userID string, username string) (string, error) {
	// 1.构建 CustomClaims
	claims := CustomClaims{
		UserID:   userID,
		Username: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "github.com/nanfeng/miniblog",
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
	}

	// 进行base加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 进行签名加密
	tokenString, err := token.SignedString([]byte("tentcoo@123"))
	if err != nil {
		return "", errmsg.NewInternalErr("token加密失败", err)
	}
	return tokenString, nil
}

// PasrseJwt 解析token
func PasrseJwt(tokenString string) (CustomClaims, error) {
	var claims CustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (any, error) {
		// 类型断言
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errmsg.New(errmsg.CodeInternalErr, "签名算法不一致", nil)
		}
		return []byte("tentcoo@123"), nil
	})
	// 校验是否解析成功
	if err != nil {
		return CustomClaims{}, errmsg.InvalidParamErr.Wrap(err)
	}

	// 判断解析的token是否有效
	if !token.Valid {
		return CustomClaims{}, errmsg.New(10002, "token无效", nil)
	}

	return claims, nil
}
