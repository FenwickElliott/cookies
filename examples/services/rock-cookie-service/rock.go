package main

import (
	"log"

	"github.com/fenwickelliott/cookies/sync"
)

func main() {
	service, err := sync.GetService("rock")
	if err != nil {
		panic(err)
	}
	log.Fatal(sync.Serve(service))
}
