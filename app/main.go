package main

import (
	"log"

	"github.com/priyanshu360/urlShortnerApp.git/server"
	"github.com/priyanshu360/urlShortnerApp.git/storage"
)




func main() {
	apiServer := server.NewServer()

	mongo, err := storage.InitMongo()
	if err != nil {
		log.Fatal("Error initializing MongoDB:", err)
	}

	apiServer.Controller.DB = mongo
	apiServer.Run()
}
