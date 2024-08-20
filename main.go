package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatalf("could not connect to DB: %s", err.Error())
	}
	fmt.Println(store)

	server := NewAPIServer(":3000", store)
	server.Run()
}
