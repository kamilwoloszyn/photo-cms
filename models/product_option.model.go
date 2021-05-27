package models

import "github.com/google/uuid"

type ProductOption struct {
	Base
	OptionValueId uuid.UUID
	ProductId     uuid.UUID
}

func (p *ProductOption) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(p).Error
}
func (p *ProductOption) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Delete(p).Error
}
func (p *ProductOption) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(p.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.First(p).Error

}
