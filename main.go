package main

import (
	"fmt"
	"nb_client/config"
	logger "nb_client/config/Logger"
	"nb_client/db"
	"nb_client/models"
)

func main() {
	logger.Info("iniciando programa")
	if err := config.Load(); err != nil {
		panic("Error")
	}

	conn, err := db.OpenConnection()
	if err != nil {
		panic(err)
	}

	db.Migrate(conn)

	all, err := models.GetAll(conn)
	if err != nil {
		logger.Error("Erro ao buscar dados", err)
	}

	fmt.Println(all)

	logger.Info("Finalizando programa")
}
