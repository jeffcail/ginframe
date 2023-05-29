package _jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"sync"
	"time"
)

var (
	EXPIRE            time.Duration
	SECRET            string
	lazyClimsInstance *JwtClaims
	once              sync.Once
)

type JwtClaims struct {
	ID       interface{} `json:"id"`
	Username string      `json:"username"`
	jwt.StandardClaims
}

func NewJwtInstance(expire time.Duration, secret string) *JwtClaims {
	if lazyClimsInstance == nil {
		once.Do(func() {
			EXPIRE = expire
			SECRET = secret
			lazyClimsInstance = &JwtClaims{}
		})
	}
	return lazyClimsInstance
}

// GenerateToken 签发Token
func (j *JwtClaims) GenerateToken(claims *JwtClaims) (string, error) {
	c := JwtClaims{
		ID:       claims.ID,
		Username: claims.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(EXPIRE).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(SECRET))
}

// ParseToken 解析TOKEN
func (j *JwtClaims) ParseToken(token string) (*JwtClaims, error) {
	t, err := jwt.ParseWithClaims(token, &JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET), nil
		})
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*JwtClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
