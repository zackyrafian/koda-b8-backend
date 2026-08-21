package libs

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(id int64, role string) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))
	claims := MyCustomClaims{
		UserID: id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  ss, err := token.SignedString(key)
  if err != nil {
    return "", err
  }
  return ss, nil
}

func VerifyToken(token string) (*MyCustomClaims, error) {
	x, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(x *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if !x.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := x.Claims.(*MyCustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}