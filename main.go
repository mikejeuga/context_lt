package main

import (
	"github.com/mikejeuga/context_lt/server"
	"log"
)

func main() {

	newServer := server.NewServer()
	err := newServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
