package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kamilwoloszyn/photo-cms/configs"
)

func Initialize(conf configs.DbConfig) (*gorm.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.Dbuser, conf.Dbpassword, conf.Database)
	db, err := gorm.Open(conf.Driver, url)
	if err != nil {
		return nil, err
	}
	return db, nil
}
