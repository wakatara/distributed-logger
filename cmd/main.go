package main

import (
	"log"

	"github.com/wakatara/distributed-logger/interal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe)
}
