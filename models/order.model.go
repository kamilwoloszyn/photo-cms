package models

import "github.com/google/uuid"

type Order struct {
	Base
	Fvat       bool
	Price      float32
	PaymentId  uuid.UUID
	DeliveryId uuid.UUID
	Product    []Product `gorm:"foreignKey:OrderId"`
}

func (o *Order) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Delete(o).Error
}

func (o *Order) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(o).Error
}

func (o *Order) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(o.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.Find(o).Error
}
