package db

import (
	"database/sql"
	logger "nb_client/config/Logger"
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
		logger.Error("Erro ao fazer a migraçao do banco de dados", err)
	}
	logger.Info("Migração rodada com sucesso.")
}
