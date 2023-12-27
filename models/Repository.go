package models

import (
	"database/sql"
)

func GetAll(db *sql.DB) (configs []AutoConfig, err error) {
	row, err := db.Query("select * from auto_config")
	if err != nil {
		return
	}

	for row.Next() {
		var config AutoConfig
		err = row.Scan(&config.ID, &config.License, &config.PathFileSend, &config.PathFileReceive, &config.IntervalTime)
		if err != nil {
			return
		}

		configs = append(configs, config)
	}

	return
}
