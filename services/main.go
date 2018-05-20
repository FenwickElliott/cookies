package main

import (
	"fmt"
	"os"

	"github.com/fenwickelliott/cookies/sync"
	"gopkg.in/mgo.v2"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("A command is required")
		return
	}
	switch os.Args[1] {
	case "new":
		new()
	default:
		fmt.Println("Command not recognized")
	}
}

func new() {
	fmt.Println("Welcome to the cookie sync service generator!")
	fmt.Println("We need to ask a few questions about the service to be generated")
	var name, address, port, redirect, mongoServer string
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	if name == "" {
		fmt.Println("A service name must be provided")
		return
	}
	fmt.Print("Address: ")
	fmt.Scanln(&address)
	if address == "" {
		fmt.Println("A service address must be provided")
		return
	}
	fmt.Print("Port: ")
	fmt.Scanln(&port)
	if port == "" {
		port = "80"
	}
	fmt.Print("MongoServer: ")
	fmt.Scanln(&mongoServer)
	if mongoServer == "" {
		mongoServer = "cookies.fenwickelliott.io"
	}
	fmt.Print("Redirect: ")
	fmt.Scanln(&redirect)

	service := sync.Service{
		Name:        name,
		Address:     address,
		Port:        port,
		Redirect:    redirect,
		MongoServer: mongoServer,
	}

	session, err := mgo.Dial("cookies.fenwickelliott.io")
	check(err)
	c := session.DB("services").C("services")

	var checkExisting sync.Service
	err = c.FindId(service.Name).One(&checkExisting)

	if err == nil {
		fmt.Println("That service name is already taken, please try again")
		return
	} else if err != nil && err.Error() != "not found" {
		check(err)
	}

	err = c.Insert(service)
	check(err)
	fmt.Println(service.Name, "successfully created")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
