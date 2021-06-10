package models

import (
	"github.com/pkg/errors"
)

type ProductOption struct {
	Base
	OptionValueId string `gorm:"type:uuid;not null"`
	ProductId     string `gorm:"type:uuid;not null"`
}

func (po *ProductOption) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(po).Error
}
func (po *ProductOption) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if po.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Delete(po).Error
}
func (po *ProductOption) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if po.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.First(po).Error
}

func (po *ProductOption) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if po.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Save(po).Error
}

func (po *ProductOption) GetProductDetails(p *Product) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if po.IsEmptyId() {
		return ErrIdEmpty
	}
	p.SetID(po.ProductId)
	if err := p.FetchByID(); err != nil {
		errWrapped := errors.Wrap(err, "Fetching product")
		return errWrapped
	}
	return nil
}

func (po *ProductOption) GetOptionValueDetails(ov *OptionValue) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(po.OptionValueId) == 0 {
		return ErrIdEmpty
	}
	ov.SetID(po.OptionValueId)
	if err := ov.FetchById(); err != nil {
		errWrapped := errors.Wrap(err, "Fetching OptionValue")
		return errWrapped
	}
	return nil
}
