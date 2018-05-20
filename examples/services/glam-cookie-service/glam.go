package main

import (
	"log"

	"github.com/fenwickelliott/cookies/sync"
)

func main() {
	service, err := sync.GetService("glam")
	if err != nil {
		panic(err)
	}
	log.Fatal(sync.Serve(service))
}
