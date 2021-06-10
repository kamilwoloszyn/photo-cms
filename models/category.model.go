package models

import "github.com/pkg/errors"

type Category struct {
	Base
	CategoryName string    `gorm:"not null"`
	Product      []Product `gorm:"foreignKey:CategoryId"`
}

func (c *Category) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.First(c).Error
}

func (c *Category) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(c).Error
}

func (c *Category) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Delete(c).Error
}

func (c *Category) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Save(c).Error
}

func (c *Category) AssignTo(p *Product) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	p.CategoryId = c.GetID()
	if err := p.UpdateInstance(); err != nil {
		errWrapped := errors.Wrap(err, "Update instance product")
		return errWrapped
	}
	return nil
}

func (c *Category) FetchProducts() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if tx := handler.Model(c).Select("categories.id,products.category_id,products.id,products.created_at,products.updated_at,products.deleted_at,products.unit_price,products.product_name,products.quantity,products.image_id,products.customer_id,products.order_id").Joins("left join products on products.category_id=categories.id").Where("categories.id=?", c.GetID()).Find(&c.Product); tx.Error != nil {
		errWrapped := errors.Wrap(tx.Error, "Fetching products connected to category")
		return errWrapped
	}
	return nil
}
