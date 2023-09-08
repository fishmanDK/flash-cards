package service

import (
	anki "github.com/fishmanDK/anki_telegram"
	"github.com/fishmanDK/anki_telegram/internal/db"
)

type Autorization interface {
	CreateUser(user anki.User) error
	Authentication(login, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Validate interface {
	ValidateRegistration(user anki.User) (*ReportValidator, error)
}

type Api interface {
	GetDataAboutUser(id string) (db.ResultInfoAboutUser, error)
	GetCategories(id string) (db.UserCategories, error)
}

type Service struct {
	Autorization
	Validate
	Api
}

func NewService(db *db.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(db.Autorization),
		Validate:     NewAuthService(db.Autorization),
		Api:          NewApiService(db.Api),
	}
}
