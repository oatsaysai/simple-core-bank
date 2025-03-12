package db

import "github.com/spf13/viper"

type Config struct {
	DBType string
}

func InitConfig() (*Config, error) {
	config := &Config{
		DBType: viper.GetString("Database.Type"),
	}

	if config.DBType == "" {
		panic("Database.Type not set")
	} else {
		switch config.DBType {
		case "postgres":
		// case "tidb":
		default:
			panic("Unsupported database type")
		}
	}
	return config, nil
}
