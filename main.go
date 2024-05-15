package main

import (
	"log"

	"github.com/BrettS89/api"
)

func main() {
	server := api.NewServer(":4000")

	log.Fatal(server.Start())
}
