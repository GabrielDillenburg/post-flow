package main

import (
	"log"

	server "github.com/gabrieldillenburg/post-flow/cmd"
)

func main() {
	server := server.InitializeServer()
	log.Printf("Server starting on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
