package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
)

type ReportValidator struct {
	Email    string
	Username string
	Password string
}

func (a *AuthDb) ValidateRegistration(email, username, password, repeatPassword string) (*ReportValidator, bool) {
	var report ReportValidator
	// report := new(ReportValidator)

	flaw := false //это для того, чтобы отследить были ли ошибки

	if password != repeatPassword {
		report.Password = "different passwords"
		flaw = true
	}

	if response := SearchEmail(a, email); response {
		report.Email = "mail already exists"
		flaw = true
	}

	if response := SearchUsername(a, username); response {
		report.Username = "username already exists"
		flaw = true
	}

	return &report, flaw
}

func SearchEmail(a *AuthDb, email string) bool {
	collection := a.client.Database("anki").Collection("user")
	filter := bson.M{
		"email": email,
	}

	result := collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		log.Println(result.Err())
	} else {
		return true
	}

	return false
}

func SearchUsername(a *AuthDb, username string) bool {
	collection := a.client.Database("anki").Collection("user")
	filter := bson.M{
		"username": username,
	}

	result := collection.FindOne(context.Background(), filter)
	if result.Err() == nil {
		log.Println(result.Err())
	} else {
		return true
	}
	return false
}
