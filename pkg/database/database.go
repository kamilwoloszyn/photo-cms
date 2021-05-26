package database

import (
	"fmt"

	"github.com/kamilwoloszyn/photo-cms/configs"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

func Initialize(conf configs.DbConfig) (*gorm.DB, error) {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Host, conf.Dbuser, conf.Dbpassword, conf.Database, conf.Port)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
