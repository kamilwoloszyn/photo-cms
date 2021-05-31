package models

import "github.com/google/uuid"

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
	return handler.Delete(c).Error
}

func (c *Customer) SetId(id uuid.UUID) {
	if len(id) > 0 {
		c.ID = id
	}
}

func (c *Customer) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(c.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.Find(c).Error

}

func (c *Customer) UpdateAll() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(c).Error
}
