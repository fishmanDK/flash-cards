package service

import (
	"errors"
	"time"

	anki "github.com/fishmanDK/anki_telegram"
	"github.com/fishmanDK/anki_telegram/internal/db"
	"github.com/golang-jwt/jwt"
)

const (
	TokenTtl   = time.Hour * 12
	signingKey = "2gtqr29ba434bf$sdh*rth!fi"
)

type AuthService struct {
	db db.Autorization
}

func NewAuthService(db db.Autorization) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (a *AuthService) CreateUser(user anki.User) error {
	person := anki.FinalUser{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	return a.db.CreateUser(person)
}

type Claims struct {
	jwt.StandardClaims
	UserId string `json: "user_id"`
}

func (a *AuthService) Authentication(login, password string) (string, error) {
	id, err := a.db.CheckUser(login, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTtl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("токен не был подписан с использованием алгоритма HMAC")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("token нельзя привести к типу *tokenClaims")
	}

	return claims.UserId, nil
}

