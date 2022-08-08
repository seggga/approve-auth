package tokens

import (
	"github.com/golang-jwt/jwt/v4"
)

// CheckToken ...
func CheckToken(tokenString, jwtSecret string) (bool, string) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return false, ""
	}

	return token.Valid, claims.Username
}
