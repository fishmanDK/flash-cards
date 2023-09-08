package db

import (
	anki "github.com/fishmanDK/anki_telegram"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Autorization interface {
	CreateUser(user anki.FinalUser) error
	SearchEmail(email string) bool
	SearchUsername(username string) bool
	CheckUser(login, password string) (string, error)
}

type Api interface {
	GetUserData(id string) (ResultInfoAboutUser, error)
	parceTitleCategories(resInfo *ResultInfoAboutUser, categoriesList []primitive.ObjectID) (*ResultInfoAboutUser, error)
	GetInfoCollectionUser(id string) (InfoCollectionUser, error)

	GetUserCategories(id string) (UserCategories, error)
}

type Repository struct {
	Autorization
	Api
}

func NuwRepository(client *mongo.Client) *Repository {
	return &Repository{
		Autorization: NewAuthDb(client),
		Api:          NewAuthDb(client),
	}
}
