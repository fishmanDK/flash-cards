package db

import (
	"context"
	"log"

	anki "github.com/fishmanDK/anki_telegram"
	"go.mongodb.org/mongo-driver/mongo"
)


type AuthDb struct{
	client *mongo.Client
}

func NewAuthDb(client *mongo.Client) *AuthDb{
	return &AuthDb{
		client: client,
	}
}

func (a *AuthDb) CreateUser(user anki.User) error{
	collection := a.client.Database("anki").Collection("user")
	log.Println(collection, user)
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
        panic(err)
    }
	return nil
}