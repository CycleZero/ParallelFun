package jwt

import (
	"github.com/golang-jwt/jwt/v4"

	"strconv"
	"time"
)

var secret = []byte("secret")

type Claims struct {
	jwt.RegisteredClaims
	UserID    string
	UserName  string
	UserEmail string
}

func GenerateToken(user *biz.User) string {
	claim := Claims{
		UserID:    strconv.Itoa(int(user.ID)),
		UserName:  user.Name,
		UserEmail: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(user.CreatedAt.Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(user.CreatedAt),
			Issuer:    "parallelfun",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString(secret)
	if err != nil {
		return ""
	}
	return signedString
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func VerifyToken(tokenString string) bool {
	_, err := ParseToken(tokenString)
	if err != nil {
		return false
	}
	return true
}
