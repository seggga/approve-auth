package tokens

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// CreateTokenPair ...
func CreateTokenPair(login, secret string) (string, string, error) {
	mySigningKey := []byte(secret)

	// generate access token
	claims := &Claims{
		Username: login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
			Issuer:    "team9",
			Subject:   "auth-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", "", err
	}

	// generate refresh token
	claims = &Claims{
		Username: login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			Issuer:    "team9",
			Subject:   "auth-service",
		},
	}
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshString, err := refresh.SignedString(mySigningKey)
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshString, nil
}
