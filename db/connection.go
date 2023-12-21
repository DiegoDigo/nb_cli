package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"nb_client/config"
)

func OpenConnection() (*sql.DB, error) {
	conf := config.GetDb()

	conn, err := sql.Open("sqlite3", conf.Name)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping()
	return conn, err
}
