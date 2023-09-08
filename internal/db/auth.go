package db

import (
	"context"

	anki "github.com/fishmanDK/anki_telegram"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (a *AuthDb) CheckUser(login, password string) (string, error) {
	var result bson.M

	collection := a.client.Database("anki").Collection("user")
	filter := bson.D{
		{"email", login},
		{"password", password},
	}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return "", err
	}

	id := result["_id"].(primitive.ObjectID)
	idString := id.Hex()

	return idString, nil
}
