package main

import (
	"log"

	"github.com/BrettS89/api"
	"github.com/BrettS89/db"
)

func main() {
	db.Client.Connect("mongodb://mongodb:27017")

	server := api.NewServer(":4000")

	log.Fatal(server.Start())
}
