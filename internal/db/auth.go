package db

import (
	"context"
	anki "github.com/fishmanDK/anki_telegram"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthDb struct {
	client *mongo.Client
}

func NewAuthDb(client *mongo.Client) *AuthDb {
	return &AuthDb{
		client: client,
	}
}

func (a *AuthDb) CreateUser(user anki.FinalUser) error {
	collection := a.client.Database("anki").Collection("user")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		panic(err)
	}
	return nil
}
