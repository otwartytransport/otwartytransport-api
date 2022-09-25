package main

import (
	"github.com/otwartytransport/otwartytransport-api/internal/server"
	"log"
)

func main() {
	srv := server.NewServer()
	log.Fatal(srv.Listen(":8010"))
}
