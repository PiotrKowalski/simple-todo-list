package main

import (
	"simple-todo-list/internal/app"
	"simple-todo-list/internal/router"
	server "simple-todo-list/internal/server"
	"simple-todo-list/pkg/config"
)

var (
	mongoURI, _ = config.ReadEnvString("MONGODB_URI")
)

func main() {
	if err := start(); err != nil {
		panic(err)
	}
}

func start() error {
	application, err := app.New(app.Config{MongoDBURI: mongoURI})
	if err != nil {
		return err
	}

	rout, err := router.New(application)
	if err != nil {
		return err
	}

	srv := server.NewRESTService(rout, server.WithPort("9000"))
	err = srv.Run()
	if err != nil {
		return err
	}
	return nil
}
