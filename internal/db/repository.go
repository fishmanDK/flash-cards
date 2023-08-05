package db

import "go.mongodb.org/mongo-driver/mongo"


type Autorization interface{

}

type Repository struct{
	Autorization
}

func NuwRepository(client *mongo.Client) *Repository{
	return &Repository{
		Autorization: client,
	}
}