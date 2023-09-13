package service

import (
	anki "github.com/fishmanDK/anki_telegram"
	"github.com/fishmanDK/anki_telegram/internal/db"
)

type QuestionService struct {
	db db.QuestionMethonds
}

func NewQuestionService(db db.QuestionMethonds) *QuestionService {
	return &QuestionService{
		db: db,
	}
}

func (a *QuestionService) GetAllQuestions(userId, categoryName string) (db.ResUserQuestions, error) {
	return  a.db.GetQuestions(userId, categoryName)
}

func (a *QuestionService) GetQuestion(userId, questonName string) (db.Question, error){
	return a.db.GetQuestion(userId, questonName)
}

func (a *QuestionService) CreateQuestion(userId, categoryName string, input anki.CreateQuestion) (string, error) {
	return a.db.CreateQuestion(userId, categoryName, input)
}

func (a *QuestionService) UpdateQustion(userId, categoryName, questionName string, input anki.UpdateQuestion) error {
	return a.db.UpdateQustion(userId, categoryName, questionName, input)
}

func (a *QuestionService) DeleteQuestion(userId, categoryName, questionName string) error {
	return a.db.DeleteQuestion(userId, categoryName, questionName)
}
