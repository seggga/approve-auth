package tokens

import "github.com/golang-jwt/jwt/v4"

// Claims ...
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
