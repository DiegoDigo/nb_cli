package main

import (
	"github.com/jasonlvhit/gocron"
	"nb_client/config"
	logger "nb_client/config/logger"
	"nb_client/db"
	"nb_client/internal/read"
	"nb_client/internal/scheduler"
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

	if len(all) > 0 {
		var sc *gocron.Scheduler
		for _, item := range all {
			sc = scheduler.CreateScheduler(item, read.Files)
		}
		<-sc.Start()
	} else {
		logger.Info("Não tem revenda cadastrada.")
	}

	logger.Info("Finalizando programa")
}
