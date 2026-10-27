package config

import "os"

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string

	DbUrl string
}

func NewDatabaseConfig() (*DatabaseConfig, error) {
	dbConfig := &DatabaseConfig{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Name:     os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		DbUrl:    os.Getenv("DB_URL"),
	}
	return dbConfig, nil
}
