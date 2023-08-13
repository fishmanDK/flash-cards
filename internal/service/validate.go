package service

import (
	"errors"
	"log"

	anki "github.com/fishmanDK/anki_telegram"
)

type ReportValidator struct {
	Email    string
	Username string
	Password string
}

func (a *AuthService) ValidateRegistration(user anki.User) (*ReportValidator, error) {
	var report ReportValidator
	// report := new(ReportValidator)

	flaw := false //это для того, чтобы отследить были ли ошибки

	if user.Password != user.RepeatPassword {
		report.Password = "different passwords"
		flaw = true
	}

	if response := a.db.SearchEmail(user.Email); response {
		report.Email = "mail already exists"
		flaw = true
	}

	if response := a.db.SearchUsername(user.Username); response {
		report.Username = "username already exists"
		flaw = true
	}

	if !flaw {
		return &ReportValidator{}, nil
	}
	log.Println(report)

	return &report, errors.New("форма не прошла валидацию")
}
