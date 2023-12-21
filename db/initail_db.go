package db

import (
	"database/sql"
	"log"
)

func Migrate(conn *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS auto_config(
        id VARCHAR(60) PRIMARY KEY,
        license VARCHAR(4) NOT NULL UNIQUE,
        path_file_send VARCHAR(255) NOT NULL,
        path_file_receive VARCHAR(255) NOT NULL,
        interval_time VARCHAR(100) NOT NULL
    );
    `

	_, err := conn.Exec(query)
	if err != nil {
		log.Fatal("erro ao conectar")
	}
	log.Println("tablea criada")
}
