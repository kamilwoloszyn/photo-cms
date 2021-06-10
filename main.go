package main

import (
	"fmt"

	"github.com/kamilwoloszyn/photo-cms/configs"
	"github.com/kamilwoloszyn/photo-cms/http/routes"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("[photo-cms]: Server starting ...")
	server, err := configs.NewServer()
	if err != nil {
		errWrapped := errors.Wrap(err, "Creating a new server")
		fmt.Printf("[photo-cms]: Cannot create a new server: %s \n", errWrapped.Error())
	}
	routes.ApplyRoutes(server)
	err = server.Listen()
	if err != nil {
		errWrapped := errors.Wrap(err, "Cannot start a new server")
		fmt.Printf("[photo-cms]: Cannot start a server: %s ", errWrapped.Error())
	}
}
