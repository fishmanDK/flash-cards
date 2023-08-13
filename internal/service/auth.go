package service

import (
	"errors"
	"html/template"
	"log"

	anki "github.com/fishmanDK/anki_telegram"
	"github.com/fishmanDK/anki_telegram/internal/db"
	"github.com/gin-gonic/gin"
)

type TemplateData struct {
    Report db.ReportValidator
}

type AuthService struct {
	db db.Autorization
}

func NewAuthService(db db.Autorization) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (a *AuthService) CreateUser(user anki.User) error {
	tmpl := template.Must(template.ParseFiles("/Users/denissvecnikov/golang/anki/internal/template/registration.html"))

	if report, flaw := a.db.ValidateRegistration(user.Email, user.Username, user.Password, user.RepeatPassword); flaw {
		log.Println(report)
		data := TemplateData{
			Report: *report,
		}
		log.Println(data.Report.Email)
		tmpl.Execute(gin.DefaultWriter, data)

		return errors.New("форма не прошла валидацию")
	}
	
	
	return a.db.CreateUser(user)
}
