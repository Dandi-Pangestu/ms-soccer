package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"ms-soccer/service/shared/domains"
	log "ms-soccer/service/shared/log/app"
)

type CustomClaims struct {
	jwt.StandardClaims
}

type Authable interface {
	Decode(tokenStr string) (*CustomClaims, error)
	Encode(userID string) (string, error)
}

type jwtService struct {
	SecretKey string
}

func NewJWTService(secretKey string) Authable {
	return &jwtService{
		SecretKey: secretKey,
	}
}

func (s *jwtService) Decode(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(s.SecretKey), nil
	})

	if err != nil {
		log.Error(&logrus.Fields{"error": err.Error()}, "Error while parse jwt")
		return nil, domains.ErrUnauthorized
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Error(nil, "Error while claims jwt")
		return nil, domains.ErrUnauthorized
	}
}

func (s *jwtService) Encode(userID string) (string, error) {
	now := time.Now().UTC()

	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(24 * time.Hour * 7).Unix(),
			IssuedAt:  now.Unix(),
			Subject:   userID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		log.Error(&logrus.Fields{"error": err.Error()}, "Error while encode jwt")
		return "", err
	}

	return tokenStr, nil
}
