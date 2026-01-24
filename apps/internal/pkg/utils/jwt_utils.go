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

// JwtHandler 定义一个结构体来管理 JWT 逻辑
type JwtHandler struct {
	secretKey      []byte        // 密钥
	expireDuration time.Duration // 过期时间
	issuer         string        // 签发者
}

func NewJwtHandler(secret string, expire time.Duration, issuer string) *JwtHandler {
	return &JwtHandler{
		secretKey:      []byte(secret),
		expireDuration: expire,
		issuer:         issuer,
	}
}

// GeneratorJwt 生成 jwt 令牌
func (j *JwtHandler) GeneratorJwt(userID string, username string) (string, error) {
	// 1.构建 CustomClaims
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expireDuration)),
		},
	}

	// 进行base加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 进行签名加密
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", errmsg.NewInternalErr("token加密失败", err)
	}
	return tokenString, nil
}

// ParseJwt 解析token
func (j *JwtHandler) ParseJwt(tokenString string) (CustomClaims, error) {
	var claims CustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (any, error) {
		// 类型断言
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errmsg.New(errmsg.CodeInternalErr, "签名算法不一致", nil)
		}
		return j.secretKey, nil
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
