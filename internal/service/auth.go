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

// func (a *AuthService) ValidateRegistration(user anki.User) (db.ReportValidator, error) {
// 	report, flaw := a.db.ValidateRegistration(user)
// 	log.Println(flaw)
// 	if !flaw {
// 		return db.ReportValidator{}, errors.New("форма не прошла валидацию")
// 	}

// 	return report, errors.New("форма не прошла валидацию")

// 	// report, flaw := a.db.ValidateRegistration(user)
// 	// if flaw {
// 	// 	return &db.ReportValidator{}, errors.New("форма не прошла валидацию")
// 	// }

// 	// return report, nil
// 	// return a.db.ValidateRegistration(user.Email, user.Username, user.Password, user.RepeatPassword)
// }
