package migrations

import (
	"github.com/kamilwoloszyn/photo-cms/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.PaymentMethod{},
		&models.Payment{},
		&models.DeliveryMethod{},
		&models.Delivery{},
		&models.Category{},
		&models.Customer{},
		&models.Image{},
		&models.Product{},
		&models.Order{},
		&models.ProductOption{},
		&models.Option{},
		&models.OptionValue{},
	); err != nil {
		return err
	}
	return nil

}
