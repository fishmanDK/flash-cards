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

type Category interface {
	// GetDataAboutUser(id string) (db.ResultInfoAboutUser, error)
	GetCategories(userId string) (db.UserCategories, error)
	GetCategoryById(userId, categoryName string) (db.Category, error)
	CreateCategory(userId, categoryName string) (string, error)
	UpdateTitle(userId, oldName, newName string) error
	DeleteCategory(userId, categoryName string) error
}

type Question interface {
	GetAllQuestions(userId, categoryName string) (db.ResUserQuestions, error)
	GetQuestion(userId, questonName string) (db.Question, error)
	CreateQuestion(userId, categoryName string, input anki.CreateQuestion) (string, error)
	UpdateQustion(userId, categoryName, questionName string, input anki.UpdateQuestion) error
	DeleteQuestion(userId, categoryName, questionName string) error
}

type Service struct {
	Autorization
	Validate
	Category
	Question
}

func NewService(db *db.Repository) *Service    {
	return &Service{
		Autorization: NewAuthService(db.Autorization),
		Validate:     NewAuthService(db.Autorization),
		Category:     NewCategoryService(db.CategoryMethonds),
		Question:     NewQuestionService(db.QuestionMethonds),
	}
}
