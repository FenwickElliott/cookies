package main

import (
	"log"

	"github.com/fenwickelliott/cookies/sync"
)

func main() {
	service, err := sync.GetService("remote")
	if err != nil {
		panic(err)
	}
	log.Fatal(sync.Serve(service))
}
