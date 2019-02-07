package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/go-yaml/yaml"
)

type ServerConfig struct {
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

type DatabaseConfig struct {
	Url          string        `yaml:"url"`
	OpenConns    int           `yaml:"open_conns"`
	IdleConns    int           `yaml:"idle_conns"`
	ConnLifetime time.Duration `yaml:"conn_lifetime"`
}

type EnvironmentConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

func LoadConfig(name string) (EnvironmentConfig, error) {

	cwd, _ := os.Getwd()
	cwd = filepath.Join(cwd, "config")

	data, err := ioutil.ReadFile(filepath.Join(cwd, name))

	var config EnvironmentConfig

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return config, err
	}

	return config, nil

}
