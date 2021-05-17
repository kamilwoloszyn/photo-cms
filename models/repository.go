package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	handler *gorm.DB
)

func SetHandler(newHandler *gorm.DB) error {
	if newHandler == nil {
		return errors.New("Handler is empty (db is not connected)")
	}
	handler = newHandler
	return nil
}

func GetHandler() *gorm.DB {
	return handler
}

func CloseDB(db *gorm.DB) error {
	return db.Close()
}
