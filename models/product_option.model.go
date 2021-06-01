package models

import "github.com/google/uuid"

type ProductOption struct {
	Base
	OptionValueId uuid.UUID
	ProductId     uuid.UUID
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
	return handler.Delete(po).Error
}
func (po *ProductOption) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(po.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.First(po).Error
}

func (po *ProductOption) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(po).Error
}

func (po *ProductOption) GetProductDetails() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return nil
}

func (po *ProductOption) GetOptionValueDetails() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return nil
}

func (po *ProductOption) GetOptionDetails() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return nil
}
