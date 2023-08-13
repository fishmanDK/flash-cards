package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (a *AuthDb) SearchEmail(email string) bool {
	collection := a.client.Database("anki").Collection("user")
	filter := bson.M{
		"email": email,
	}

	var result bson.M
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err == nil {
		return true
	} else if err != mongo.ErrNoDocuments {
		log.Println(err)
	}

	return false
}

func (a *AuthDb) SearchUsername(username string) bool {
	collection := a.client.Database("anki").Collection("user")
	filter := bson.M{
		"username": username,
	}

	var result bson.M
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err == nil {
		return true
	} else if err != mongo.ErrNoDocuments {
		log.Println(err)
	}

	return false
}
