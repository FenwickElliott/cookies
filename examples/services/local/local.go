package main

import (
	"log"

	"github.com/fenwickelliott/cookies/sync"
)

func main() {
	log.Fatal(sync.Serve(sync.Service{
		Name:        "local",
		Port:        "443",
		Address:     "https://cookies.fenwickelliott.io",
		Redirect:    "https://cookie-sync-203420.appspot.com",
		MongoServer: "cookies.fenwickelliott.io",
		TLS:         true,
	}))
}
