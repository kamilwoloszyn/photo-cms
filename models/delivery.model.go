package models

import "github.com/google/uuid"

type Delivery struct {
	Base
	ShippedVia               string `gorm:"not null"`
	TrackingCode             string
	DestinationPostalCode    string `gorm:"not null"`
	DestinationConturyRegion string
	DestinationAddress       string `gorm:"not null"`
	DestinationCity          string `gorm:"not null"`
	PaymentMethodId          uuid.UUID
	DeliveryMethodId         uuid.UUID
	Order                    []Order `gorm:"foreignKey:DeliveryId"`
}

func (d *Delivery) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.First(d).Error
}

func (d *Delivery) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Delete(d).Error
}

func (d *Delivery) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(d.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.Create(d).Error
}

func (d *Delivery) UpdateAll() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(d).Error
}
