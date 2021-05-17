package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/kamilwoloszyn/photo-cms/models"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.PaperType{})
	db.AutoMigrate(&models.Format{})
	db.AutoMigrate(&models.Photo{})
	db.AutoMigrate(&models.Print{})
	db.AutoMigrate(&models.Order{})

}
