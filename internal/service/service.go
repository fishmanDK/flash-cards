package service

import (
	anki "github.com/fishmanDK/anki_telegram"
	"github.com/fishmanDK/anki_telegram/internal/db"
)

type Autorization interface {
	CreateUser(user anki.User) error
	// ValidateRegistration(user anki.User) (*ReportValidator, bool)
}

type Service struct {
	Autorization
}

func NewService(db *db.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(db.Autorization),
	}
}

