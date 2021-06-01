package models

import "github.com/google/uuid"

type Order struct {
	Base
	Fvat       bool    `gorm:"default:false"`
	Price      float32 `gorm:"not null"`
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

func (o *Order) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(o).Error
}

func (o *Order) GetProducts() error {
	if handler == nil {
		return ErrHandlerNotFound
	}

	return nil
}

func (o *Order) AssignTo(p *Product) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return nil
}

func (o *Order) GetPaymentDetails(p *Payment) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return nil
}

func (o *Order) GetDeliveryDetails(d *Delivery) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return nil
}
