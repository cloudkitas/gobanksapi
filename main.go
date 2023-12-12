package main

import (
	"log"
)

func main() {

	store, err := NewPosgreStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.init(); err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":9000", store)
	server.Run()
}
