package db

import (
	anki "github.com/fishmanDK/anki_telegram"
	"go.mongodb.org/mongo-driver/mongo"
)

type Autorization interface {
	CreateUser(user anki.FinalUser) error
	SearchEmail(email string) bool
	SearchUsername(username string) bool
}

type Repository struct {
	Autorization
}

func NuwRepository(client *mongo.Client) *Repository {
	return &Repository{
		Autorization: NewAuthDb(client),
	}
}
