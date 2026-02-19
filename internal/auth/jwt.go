package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type CustomClaims struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	AccountID uint   `json:"account_id"`
	Role      string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(user *User) (string, error) {
	claims := CustomClaims{
		UserID:    user.ID,
		Username:  user.Username,
		AccountID: user.AccountID,
		Role:      user.Account.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
