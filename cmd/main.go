package main

import (
	"github.com/bayamasa/proglog/internal/server"
	"log"
)

func main() {
	srv := server.NewHttpServer("localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
