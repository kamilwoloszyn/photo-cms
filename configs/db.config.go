package configs

import (
	"errors"
	"os"
)

type DbConfig struct {
	Driver      string `envconfig:"STORAGE_DRIVER"`
	HandlerName string `envconfig:"STORAGE_HANDLERNAME"`
	Host        string `default:"127.0.0.1:5433" envconfig:"STORAGE_HOST"`
	Port        string `envconfig:"STORAGE_PORT"`
	Dbuser      string `envconfig:"STORAGE_DBUSER"`
	Dbpassword  string `envconfig:"STORAGE_DBPASSWORD"`
	Database    string `envconfig:"STORAGE_DATABASE"`
}

func LoadDbConfig() (*DbConfig, error) {
	if _, found := os.LookupEnv("STORAGE_DRIVER"); !found {
		return nil, errors.New("Variable STORAGE_DRIVER is not present")
	}
	if _, found := os.LookupEnv("STORAGE_HANDLERNAME"); !found {
		return nil, errors.New("Variable STORAGE_HANDLERNAME is not present")
	}
	if _, found := os.LookupEnv("STORAGE_HOST"); !found {
		return nil, errors.New("Variable STORAGE_HOST is not present")
	}
	if _, found := os.LookupEnv("STORAGE_DBUSER"); !found {
		return nil, errors.New("Variable STORAGE_DBUSER is not present")
	}
	if _, found := os.LookupEnv("STORAGE_DBPASSWORD"); !found {
		return nil, errors.New("Variable STORAGE_DBPASSWORD is not present")
	}
	if _, found := os.LookupEnv("STOARGAGE_DATABASE"); !found {
		return nil, errors.New("Variable STOARGAGE_DATABASE is not present")
	}
	var config = DbConfig{
		Driver:      os.Getenv("STORAGE_DRIVER"),
		HandlerName: os.Getenv("STORAGE_HANDLERNAME"),
		Host:        os.Getenv("STORAGE_HOST"),
		Dbuser:      os.Getenv("STORAGE_DBUSER"),
		Dbpassword:  os.Getenv("STORAGE_DBPASSWORD"),
		Database:    os.Getenv("STOARGAGE_DATABASE"),
	}
	return &config, nil
}
