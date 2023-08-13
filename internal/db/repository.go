package db

import (
	anki "github.com/fishmanDK/anki_telegram"
	"go.mongodb.org/mongo-driver/mongo"
)


type Autorization interface{
	CreateUser(user anki.User) error
	ValidateRegistration(email, username, password, repeatPassword string) (*ReportValidator, bool)
}

type Repository struct{
	Autorization
}

func NuwRepository(client *mongo.Client) *Repository{
	return &Repository{
		Autorization: NewAuthDb(client),
	}
}