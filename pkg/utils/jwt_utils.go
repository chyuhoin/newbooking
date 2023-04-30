package utils

import (
	"github.com/dgrijalva/jwt-go"
	"newbooking/pkg/entity"
	"time"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// CustomClaims 载荷，可添加自己需要的一些信息
type CustomClaims struct {
	Id       string `json:"userId"`
	UserName string `json:"userName"`
	RoleId   string `json:"roleId"`
	jwt.StandardClaims
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(SignKey),
	}
}

// CreateToken 通过claims生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

// GenerateToken 给定用户和过期时间，生成一个jwt的token
func GenerateToken(user *entity.User, expiredTimeByMinute int64) (*string, error) {
	j := &JWT{
		[]byte(SignKey),
	}
	claims := CustomClaims{
		user.Id,
		user.Username,
		user.Role,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),                   // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + expiredTimeByMinute*60), // 过期时间 一小时
			Issuer:    "zhy",                                             //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		return nil, err
	}

	return &token, nil
}
