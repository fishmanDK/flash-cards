package service

import (
	anki "github.com/fishmanDK/anki_telegram"
	"github.com/fishmanDK/anki_telegram/internal/db"
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
