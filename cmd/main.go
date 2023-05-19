package main

import (
	"log"
	"github.com/bayamasa/proglog/internal/server"
)

func main() {
	srv := server.NewHttpServer("localhost:8080")
	log.Fatal(srv.ListenAndServe())
}

