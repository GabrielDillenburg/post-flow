package main

import (
	"log"
)

func main() {
	server := InitializeServer()
	log.Fatal(server.ListenAndServe())
}
