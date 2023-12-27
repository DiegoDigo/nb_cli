package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"nb_client/config"
	logger "nb_client/config/logger"
)

func OpenConnection() (*sql.DB, error) {
	conf := config.GetDb()

	conn, err := sql.Open("sqlite3", conf.Name)
	if err != nil {
		logger.Error("Error ao abrir a conexão com bando de dados.", err)
	}
	err = conn.Ping()
	return conn, err
}
