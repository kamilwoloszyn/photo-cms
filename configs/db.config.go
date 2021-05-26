package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	HandlerName string
	Host        string
	Port        string
	Dbuser      string
	Dbpassword  string
	Database    string
}

func LoadDbConfig() (*DbConfig, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}
	if _, found := os.LookupEnv("STORAGE_HANDLER"); !found {
		return nil, errors.New("Variable STORAGE_HANDLER is not present")
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
	if _, found := os.LookupEnv("STORAGE_DATABASE"); !found {
		return nil, errors.New("Variable STORAGE_DATABASE is not present")
	}
	var config = DbConfig{
		HandlerName: os.Getenv("STORAGE_HANDLER"),
		Host:        os.Getenv("STORAGE_HOST"),
		Dbuser:      os.Getenv("STORAGE_DBUSER"),
		Dbpassword:  os.Getenv("STORAGE_DBPASSWORD"),
		Database:    os.Getenv("STORAGE_DATABASE"),
		Port:        os.Getenv("STORAGE_PORT"),
	}
	return &config, nil
}
