package config

import "github.com/spf13/viper"

var cfg *config

type config struct {
	Firebase firebaseConfig
	DB       dbConfig
}

type firebaseConfig struct {
	Url string
}

type dbConfig struct {
	Name string
}

func init() {
	viper.SetDefault("firebase.url", "https://api-wn6zribkia-uc.a.run.app")
	viper.SetDefault("db.name", "bi_cli.sqlite")
}

func Load() error {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.Firebase = firebaseConfig{
		Url: viper.GetString("firebase.url"),
	}

	cfg.DB = dbConfig{
		Name: viper.GetString("db.name"),
	}

	return nil
}

func GetFirebase() firebaseConfig {
	return cfg.Firebase
}

func GetDb() dbConfig {
	return cfg.DB
}
