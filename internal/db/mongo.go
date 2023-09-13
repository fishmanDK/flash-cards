package db

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)




func ConectMongo() *mongo.Client {
	if err := initConfig(); err != nil{
		log.Fatal(err.Error())
	}

	uri := viper.GetString("uriMongoDb")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	
	if err != nil {
		log.Fatalf(err.Error())
	}
	// defer client.Disconnect(ctx)
	
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf(err.Error())
	}
	
	return client
}
 
func initConfig() error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("db")

	return viper.ReadInConfig()
}