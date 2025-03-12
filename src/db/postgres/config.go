package postgres

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost       string
	DBPort       string
	DBUsername   string
	DBPassword   string
	DBName       string
	DBSchemaName string
	MaxOpenConns int32
}

func InitConfig() (*Config, error) {
	config := &Config{
		DBHost:       viper.GetString("Database.PostgreSQL.DBHost"),
		DBPort:       viper.GetString("Database.PostgreSQL.DBPort"),
		DBUsername:   viper.GetString("Database.PostgreSQL.DBUsername"),
		DBPassword:   viper.GetString("Database.PostgreSQL.DBPassword"),
		DBName:       viper.GetString("Database.PostgreSQL.DBName"),
		DBSchemaName: viper.GetString("Database.PostgreSQL.DBSchemaName"),
		MaxOpenConns: viper.GetInt32("Database.PostgreSQL.MaxOpenConns"),
	}
	if config.DBHost == "" {
		config.DBHost = "localhost"
	}
	if config.DBPort == "" {
		config.DBPort = "5432"
	}
	if config.DBUsername == "" {
		config.DBUsername = "postgres"
	}
	if config.DBPassword == "" {
		config.DBPassword = "postgres"
	}
	if config.DBName == "" {
		config.DBName = "postgres"
	}
	if config.DBSchemaName == "" {
		config.DBSchemaName = "public"
	}

	return config, nil
}
