package main

import (
	"log"

	HTTP "github.com/xueyiyao/safekeep/http"
)

func main() {
	server := HTTP.NewServer()

	err := server.Run()

	if err != nil {
		log.Fatal("Could not run server.")
	}
}
