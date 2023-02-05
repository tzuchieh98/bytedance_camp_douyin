package util

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"github.com/linzijie1998/bytedance_camp_douyin/model"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	UserInfo model.UserInfo
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.DOUYIN_CONFIG.JWT.SigningKey),
	}
}

// 创建 token
func (j *JWT) CreateToken(userInfo model.UserInfo) (string, error) {
	ep, _ := ParseDuration(global.DOUYIN_CONFIG.JWT.ExpiresTime)
	claims := CustomClaims{
		UserInfo: userInfo,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,        // 签名生效时间
			ExpiresAt: time.Now().Add(ep).Unix(),       // 过期时间 7天  配置文件
			Issuer:    global.DOUYIN_CONFIG.JWT.Issuer, // 签名的发行者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		// if ve, ok := err.(*jwt.ValidationError); ok {
		// 	if ve.Errors&jwt.ValidationErrorMalformed != 0 {
		// 		return nil, TokenMalformed
		// 	} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
		// 		// Token is expired
		// 		return nil, TokenExpired
		// 	} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
		// 		return nil, TokenNotValidYet
		// 	} else {
		// 		return nil, TokenInvalid
		// 	}
		// }
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
