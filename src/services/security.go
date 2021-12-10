package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SecurityService struct {
}

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserId   int    `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func (s SecurityService) GenerateToken(username string, password string, userId int, secret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		s.EncodeSha(username, secret),
		s.EncodeSha(password, secret),
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

func (s SecurityService) EncodeSha(value string, secret string) string {
	m := hmac.New(sha256.New, []byte(secret))
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
