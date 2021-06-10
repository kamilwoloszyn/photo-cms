package models

import (
	"github.com/pkg/errors"
)

type Image struct {
	Base
	Name     string `gorm:"not null"`
	FullPath string `gorm:"not null"`
	Size     uint32
	Product  Product `gorm:"foreignKey:ImageId"`
}

func (i *Image) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(i).Error
}

func (i *Image) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if i.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.First(i).Error
}

func (i *Image) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if i.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Delete(i).Error
}

func (i *Image) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if i.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Save(i).Error
}

func (i *Image) AssignTo(p *Product) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if i.IsEmptyId() {
		return ErrIdEmpty
	}
	p.ImageId = i.GetID()
	if err := p.UpdateInstance(); err != nil {
		errWrapped := errors.Wrap(err, "Updating product instance")
		return errWrapped
	}
	return nil
}

func (i *Image) GetProduct() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	tx := handler.Model(i).Select("images.id,products.id,products.created_at,products.updated_at,products.deleted_at,products.unit_price,products.product_name,products.quantity,products.category_id,products.image_id,products.customer_id,products.order_id").Joins("left join products on products.image_id=images.id").Where("images.id=?", i.GetID()).Find(&i.Product)
	if tx != nil {
		errWrapped := errors.Wrap(tx.Error, "Joining image with products")
		return errWrapped
	}
	return nil
}
