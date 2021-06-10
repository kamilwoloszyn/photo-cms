package models

import (
	"github.com/kamilwoloszyn/photo-cms/pkg/checkers"
	"github.com/pkg/errors"
)

type Customer struct {
	Base
	City         string `gorm:"not null"`
	Address      string `gorm:"not null"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	PostalCode   string `gorm:"not null"`
	CompanyName  string `gorm:"not null"`
	PhoneNumber  string `gorm:"not null"`
	EmailAddress string `gorm:"unique; not null"`
	Employed     bool   `gorm:"default:false;"`
	NIP          string `gorm:"unique; not null"`
	Regon        string `gorm:"unique"`
	Products     []Product
}

func (c *Customer) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(c).Error
}
func (c *Customer) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Delete(c).Error
}

func (c *Customer) SetId(id string) {
	if newId := checkers.UuidString(id); newId.IsValid() {
		c.ID = id
	}
}

func (c *Customer) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Find(c).Error

}

func (c *Customer) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Save(c).Error
}

func (c *Customer) GetProducts() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	tx := handler.Model(c).Select("customers.id,products.id,products.created_at,products.updated_at,products.deleted_at,products.unit_price,products.product_name,products.quantity,products.category_id,products.image_id,products.customer_id,products.order_id").Joins("left join products on products.customer_id=customers.id").Where("customers.id=?", c.GetID()).Find(&c.Products)

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (c *Customer) AssignTo(p *Product) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	p.CustomerId = c.GetID()
	if err := p.UpdateInstance(); err != nil {
		errWrapped := errors.Wrapf(err, "Updating instance customer -> product")
		return errWrapped
	}
	return nil
}
