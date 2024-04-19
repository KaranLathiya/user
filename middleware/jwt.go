package middleware

import (
	"user/config"

	error_handling "user/error"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
}

func VerifyJWTToken(token string,audience string, subject string) error {
	jwtKey := []byte(config.ConfigVal.JWTKey)
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return error_handling.JWTErrSignatureInvalid
		}
		return error_handling.InternalServerError
	}	

	if !tkn.Valid {
		return error_handling.JWTTokenInvalid
	}
	if claims.Audience != audience && claims.Subject == subject {
		return error_handling.JWTTokenInvalidDetails
	}
	return nil
}
