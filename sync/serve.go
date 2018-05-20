package sync

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
)

var (
	err error
	// c       *mgo.Collection
	service Service
)

// Serve starts the service according to the given Service struct
func Serve(serviceVars Service) error {
	service = serviceVars
	if service.Name == "" {
		return errors.New("A service name must be provided")
	}

	session, err := mgo.Dial(service.MongoServer)
	if err != nil {
		return err
	}
	defer session.Close()
	service.c = session.DB(service.Name).C("master")

	http.HandleFunc("/print", service.print)
	http.HandleFunc("/in", service.in)
	http.HandleFunc("/out", service.out)
	http.HandleFunc("/forward", service.forward)
	http.HandleFunc("/back", service.back)
	if service.TLS {
		fmt.Println("Serving:", service.Name, "on port:", service.Port)
		return http.ListenAndServeTLS(":"+service.Port, "server.crt", "server.key", nil)
	}
	fmt.Println("Serving:", service.Name, "on port:", service.Port)
	return http.ListenAndServe(":"+service.Port, nil)
}
