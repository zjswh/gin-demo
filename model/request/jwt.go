package request

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	Id int
	Username string
	Phone string
	Avatar string
	jwt.StandardClaims
}
