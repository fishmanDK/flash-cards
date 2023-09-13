package db

import (
	anki "github.com/fishmanDK/anki_telegram"
	"go.mongodb.org/mongo-driver/mongo"
)

type Autorization interface {
	CreateUser(user anki.FinalUser) error
	SearchEmail(email string) bool
	SearchUsername(username string) bool
	CheckUser(login, password string) (string, error)
}

// type Api interface {
// 	GetUserData(id string) (ResultInfoAboutUser, error)
// 	GetInfoCollectionUser(id string) (InfoCollectionUser, error)
// 	parceTitleCategories(resInfo *ResultInfoAboutUser, categoriesList []primitive.ObjectID) (*ResultInfoAboutUser, error)
// }

type CategoryMethonds interface {
	GetUserCategories(id string) (UserCategories, error)
	GetUserCategoryById(userId, categoryName string) (Category, error)
	CreateCategory(userId, categoryName string) (string, error)
	UpdateTitle(userId, oldName, newName string) error
	DeleteCategory(userId, categoryName string) error
}

type QuestionMethonds interface {
	GetQuestions(userId, categoryName string) (ResUserQuestions, error)
	GetQuestion(userId, qestionName string) (Question, error)
	CreateQuestion(userId, categoryName string, input anki.CreateQuestion) (string, error)
	UpdateQustion(userId, categoryName, questionName string, input anki.UpdateQuestion) error
	DeleteQuestion(userId, categoryName, questionName string) error
}

type Repository struct {
	Autorization
	// Api
	CategoryMethonds
	QuestionMethonds
}

func NuwRepository(client *mongo.Client) *Repository {
	return &Repository{
		Autorization: NewAuthDb(client),
		// Api:          NewAuthDb(client),
		CategoryMethonds: NewCategoryDb(client),
		QuestionMethonds: NewQuestionDb(client),
	}
}
