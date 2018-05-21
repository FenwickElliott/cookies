package main

import (
	"fmt"
	"os"

	"github.com/fenwickelliott/cookies/sync"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("A command is required")
		return
	}
	switch os.Args[1] {
	case "new":
		new()
	case "update":
		update()
	case "delete":
		delete()
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
	defer session.Close()
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

func update() {
	var name, address, port, redirect, mongoServer string
	fmt.Print("Name: ")
	session, err := mgo.Dial("cookies.fenwickelliott.io")
	check(err)
	defer session.Close()
	c := session.DB("services").C("services")
	fmt.Scanln(&name)
	var res bson.M
	err = c.FindId(name).One(&res)
	if err != nil && err.Error() == "not found" {
		fmt.Println("Service", name, "not found")
		return
	}
	check(err)
	fmt.Print("Address: ")
	fmt.Scanln(&address)
	if address != "" {
		err = c.UpdateId(name, bson.M{"$set": bson.M{"host": address}})
		check(err)
	}
	fmt.Print("Port: ")
	fmt.Scanln(&port)
	if port != "" {
		err = c.UpdateId(name, bson.M{"$set": bson.M{"port": port}})
		check(err)
	}
	fmt.Print("Redirect: ")
	fmt.Scanln(&redirect)
	if redirect != "" {
		err = c.UpdateId(name, bson.M{"$set": bson.M{"redirect": redirect}})
		check(err)
	}
	fmt.Print("MongoServer: ")
	fmt.Scanln(&mongoServer)
	if mongoServer != "" {
		err = c.UpdateId(name, bson.M{"$set": bson.M{"mongoserver": mongoServer}})
		check(err)
	}
	fmt.Println(name, "successfully updated")
}
func delete() {
	var name string
	fmt.Print("Name of service to delete: ")
	fmt.Scanln(&name)
	session, err := mgo.Dial("cookies.fenwickelliott.io")
	check(err)
	defer session.Close()
	c := session.DB("services").C("services")
	err = c.RemoveId(name)
	if err != nil && err.Error() == "not found" {
		fmt.Println("No service found named", name)
		return
	}
	check(err)
	fmt.Println(name, "successfully deleted")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
