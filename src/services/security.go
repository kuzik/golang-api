package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SecurityService struct {
	Secret string
}

type Claims struct {
	Username string `json:"username"`
	UserId   int    `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func (s SecurityService) GenerateToken(username string, userId int, expireTime time.Time) (string, error) {

	claims := Claims{
		username,
		userId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	var jwtSecret []byte
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func (s SecurityService) ParseToken(token string) (*Claims, error) {
	var jwtSecret []byte
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func (s SecurityService) EncodeSha(value string) string {
	m := hmac.New(sha256.New, []byte(s.Secret))
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
