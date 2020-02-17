package model

import "github.com/dgrijalva/jwt-go"

//JwtClaims for user credentials
type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
