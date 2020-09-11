package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-demo/global"
	"go-demo/global/response"
	"go-demo/model/request"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("登录过期,请重试")
	TokenNotValidYet = errors.New("登录信息失效")
	TokenMalformed   = errors.New("登录凭证异常")
	TokenInvalid     = errors.New("登录凭证异常")
)

func NewJwt() *JWT {
	return &JWT{[]byte(global.GVA_CONFIG.Jwt.SigningKey)}
}

func JwtAuth() gin.HandlerFunc  {
	return func(c * gin.Context) {
		token := c.GetHeader("token")
		if token == ""  {
			response.LoginError("未登录", c)
		}
		newJwt := NewJwt()
		userClaims, err := newJwt.ParseToken(token)
		if err != nil {
			response.Result(3, "", err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", userClaims)
		c.Next()
	}
}


func CreateToken(claims request.UserClaims) (string, error) {
	newJwt := NewJwt()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(newJwt.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*request.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.UserClaims{}, func(token *jwt.Token) (i interface{}, err error) {
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

	if token != nil {
		if claims, ok := token.Claims.(*request.UserClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}
}


