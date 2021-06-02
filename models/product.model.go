package models

import (
	"github.com/google/uuid"
	"github.com/kamilwoloszyn/photo-cms/pkg/database"
	_ "github.com/kamilwoloszyn/photo-cms/pkg/database"
	"github.com/pkg/errors"
)

type Product struct {
	Base
	UnitPrice     float32
	ProductName   string
	Quantity      uint32
	ProductOption []ProductOption `gorm:"foreignKey:ProductId"`
	CategoryId    uuid.UUID       `gorm:"not null"`
	ImageId       uuid.UUID       `gorm:"not null"`
	CustomerId    uuid.UUID       `gorm:"not null"`
	OrderId       uuid.UUID       `gorm:"default: null"`
}

func (p *Product) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(p).Error
}
func (p *Product) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Delete(p).Error
}
func (p *Product) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(p).Error
}
func (p *Product) FetchByID() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if p.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.First(p).Error
}
func (p *Product) AssignTo(po *ProductOption) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if po.IsEmptyId() {
		return ErrIdEmpty
	}
	po.ProductId = p.GetID()
	if err := po.UpdateInstance(); err != nil {
		errWrapped := errors.Wrap(err, "Updating ProductOption instance")
		return errWrapped
	}
	return nil
}
func (p *Product) GetCustomerDetails(c *Customer) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if database.IsCorrectId(p.CustomerId.String()) {
		return ErrIdEmpty
	}
	c.SetID(p.CustomerId)
	if err := c.FetchById(); err != nil {
		errWrapped := errors.Wrap(err, "Fetching customer details")
		return errWrapped
	}
	return nil
}
func (p *Product) GetOrderDetails(o *Order) error {
	if handler == nil {
		return ErrHandlerNotFound
	}

	return nil
}
func (p *Product) GetCategoryDetails(c *Category) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return nil
}

func (p *Product) GetImageDetails(c *Category) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return nil
}
