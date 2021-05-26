package models

import (
	"errors"
	"fmt"

	"github.com/kamilwoloszyn/photo-cms/configs"
	"github.com/kamilwoloszyn/photo-cms/pkg/database"
	"gorm.io/gorm"
)

var (
	handler            *gorm.DB
	ErrHandlerNotFound = errors.New("db is not connected")
	ErrIdEmpty         = errors.New("id is required to fetching from model db")
)

func SetHandler(newHandler *gorm.DB) error {
	if newHandler == nil {
		return errors.New("handler is empty (db is not connected)")
	}
	handler = newHandler
	return nil
}

func GetHandler() *gorm.DB {
	return handler
}

func CloseDB(db *gorm.DB) error {
	instance, err := db.DB()
	if err != nil {
		return err
	}
	instance.Close()
	return nil
}

func Connect() error {
	if handler != nil {
		return nil
	}
	cfg, err := configs.LoadDbConfig()
	if err != nil {
		errWrapped := fmt.Sprintf("Cannot load database config: %s", err.Error())
		return errors.New(errWrapped)
	}
	db, err := database.Initialize(*cfg)
	if err != nil {
		errWrapped := fmt.Sprintf("Couldn't initialize database: %s", err.Error())
		return errors.New(errWrapped)
	}
	if err := SetHandler(db); err != nil {
		errWrapped := fmt.Sprintf("Couldn't set handler: %s", err.Error())
		return errors.New(errWrapped)
	}
	return nil

}
