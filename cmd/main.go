package main

import (
	"github.com/fishmanDK/anki_telegram/internal/db"
	"github.com/fishmanDK/anki_telegram/internal/handlers"
	"github.com/fishmanDK/anki_telegram/internal/service"
)

func main() {
	client := db.ConectMongo()
	db := db.NuwRepository(client)
	service := service.NewService(db)
	handlers := handlers.NewHandlers(*service)
	server := handlers.InitRouts()

	server.Run(":8080")
}


