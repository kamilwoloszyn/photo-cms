package migrations

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Option{})
	db.AutoMigrate(&models.OptionValue{})
	db.AutoMigrate(&models.ProductOption{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Image{})
	db.AutoMigrate(&models.DeliveryMethod{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.PaymentMethod{})
	db.AutoMigrate(&models.Payment{})
	db.AutoMigrate(&models.Delivery{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Order{})

}
