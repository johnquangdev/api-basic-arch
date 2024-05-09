package token

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Gmail string `json:"gmail"`
	jwt.StandardClaims
}

// generate jwt
func GenerateJwt(ctx context.Context, tokenTimeLife time.Duration, secretkey string, name string, gmail string) (string, error) {
	if secretkey == "" {
		return "", fmt.Errorf("Secret key not found")
	}
	claims := &JwtCustomClaims{
		Name:  name,
		Gmail: gmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * tokenTimeLife).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
