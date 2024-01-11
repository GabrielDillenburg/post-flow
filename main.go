package main

import (
	"log"

	server "github.com/gabrieldillenburg/post-flow/cmd"
)

func main() {
	server := server.InitializeServer()
	log.Fatal(server.ListenAndServe())
}
