package main

import (
	"nb_client/config"
	"nb_client/db"
	"nb_client/internal"
)

func main() {
	if err := config.Load(); err != nil {
		panic("Error")
	}

	conn, err := db.OpenConnection()
	if err != nil {
		panic(err)
	}

	db.Migrate(conn)

	internal.HttpServer()
}
