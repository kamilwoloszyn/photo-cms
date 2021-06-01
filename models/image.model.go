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
	return handler.First(i).Error
}

func (i *Image) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(i.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.Delete(i).Error
}

func (i *Image) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(i).Error
}

func (i *Image) AssignTo(p *Product) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if i.IsEmptyId() || p.IsEmptyId() {
		return ErrIdEmpty
	}
	p.ImageId = i.GetID()
	if err := p.UpdateInstance(); err != nil {
		errWrapped := errors.Wrap(err, "Updating product instance")
		return errWrapped
	}
	return nil
}
