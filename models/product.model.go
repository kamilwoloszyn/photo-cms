package models

import (
	"github.com/kamilwoloszyn/photo-cms/pkg/checkers"
	_ "github.com/kamilwoloszyn/photo-cms/pkg/database"
	"github.com/pkg/errors"
)

type Product struct {
	Base
	UnitPrice     float32
	ProductName   string
	Quantity      uint32
	ProductOption []ProductOption `gorm:"foreignKey:ProductId"`
	CategoryId    string          `gorm:"type:uuid;not null"`
	ImageId       string          `gorm:"type:uuid;not null;unique"`
	CustomerId    string          `gorm:"type:uuid;not null"`
	OrderId       *string         `gorm:"type:uuid;default:null"`
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
	if p.IsEmptyId() {
		return ErrIdEmpty
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
	if p.IsEmptyId() {
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
	if customerId := checkers.UuidString(p.CustomerId); customerId.IsEmpty() {
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

	if orderId := checkers.UuidString(*p.OrderId); orderId.IsEmpty() {
		return ErrIdEmpty
	}
	o.ID = *p.OrderId
	if err := o.FetchById(); err != nil {
		errWrapped := errors.Wrap(err, "Fetching order details")
		return errWrapped
	}
	return nil
}
func (p *Product) GetCategoryDetails(c *Category) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if categoryId := checkers.UuidString(p.CategoryId); categoryId.IsEmpty() {
		return ErrIdEmpty
	}
	c.SetID(p.CategoryId)
	if err := c.FetchById(); err != nil {
		errWrapped := errors.Wrap(err, "Fetching category")
		return errWrapped
	}
	return nil
}

func (p *Product) GetImageDetails(i *Image) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if imageId := checkers.UuidString(p.ImageId); imageId.IsEmpty() {
		return ErrIdEmpty
	}
	i.SetID(p.ImageId)
	if err := i.FetchById(); err != nil {
		errWrapped := errors.Wrap(err, "Fetching Image")
		return errWrapped
	}
	return nil
}

func (p *Product) GetProductOptions() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if p.IsEmptyId() {
		return ErrIdEmpty
	}
	tx := handler.Model(p).Select("products.id,product_options.id,product_options.created_at,product_options.updated_at,product_options.deleted_at,product_options.option_value_id,product_options.product_id").Joins("left join product_options on product_options.product_id=products.id").Where("products.id=?", p.GetID()).Find(&p.ProductOption)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "Join products with product_options")
	}
	return nil
}
