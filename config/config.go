package config

import (
	"os"
	"strconv"
	"time"
)

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type DatabaseConfig struct {
	Url          string
	OpenConns    int
	IdleConns    int
	ConnLifetime time.Duration
}

type EnvironmentConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
}

func LoadConfig(name string) (EnvironmentConfig, error) {

	serverPort := os.Getenv("SERVER_PORT")
	serverReadTimeout, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	serverWriteTimeout, _ := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
	serverIdleTimeout, _ := strconv.Atoi(os.Getenv("SERVER_IDLE_TIMEOUT"))

	databaseUrl := os.Getenv("DATABASE_URL")
	databaseOpenConns, _ := strconv.Atoi(os.Getenv("DATABASE_OPEN_CONNS"))
	databaseIdleConns, _ := strconv.Atoi(os.Getenv("DATABASE_IDLE_CONNS"))
	databaseConnLifetime, _ := strconv.Atoi(os.Getenv("DATABASE_CONN_LIFETIME"))

	config := EnvironmentConfig{
		Server: ServerConfig{
			Port:         serverPort,
			ReadTimeout:  time.Duration(serverReadTimeout) * time.Second,
			WriteTimeout: time.Duration(serverWriteTimeout) * time.Second,
			IdleTimeout:  time.Duration(serverIdleTimeout) * time.Second,
		},
		Database: DatabaseConfig{
			Url:          databaseUrl,
			OpenConns:    databaseOpenConns,
			IdleConns:    databaseIdleConns,
			ConnLifetime: time.Duration(databaseConnLifetime) * time.Second,
		},
	}

	return config, nil

}
