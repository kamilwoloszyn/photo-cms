package models

import (
	"github.com/kamilwoloszyn/photo-cms/pkg/checkers"
	"github.com/pkg/errors"
)

type Delivery struct {
	Base
	ShippedVia               string `gorm:"not null"`
	TrackingCode             string
	DestinationPostalCode    string `gorm:"not null"`
	DestinationCountryRegion string
	DestinationAddress       string  `gorm:"not null"`
	DestinationCity          string  `gorm:"not null"`
	DeliveryMethodId         string  `gorm:"type:uuid;not null"`
	Order                    []Order `gorm:"foreignKey:DeliveryId"`
}

func (d *Delivery) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if d.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.First(d).Error
}

func (d *Delivery) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if d.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Delete(d).Error
}

func (d *Delivery) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(d).Error
}

func (d *Delivery) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if d.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Save(d).Error
}

func (d *Delivery) GetDeliveryMethodDetails(deliveryMethod *DeliveryMethod) error {
	deliveryMethodId := checkers.UuidString(d.DeliveryMethodId)
	if handler == nil {
		return ErrHandlerNotFound
	}
	if deliveryMethodId.IsEmpty() {
		return ErrIdEmpty
	}
	deliveryMethod.ID = d.GetID()
	if err := deliveryMethod.FetchById(); err != nil {
		errWrapped := errors.Wrap(err, "Fetching by id")
		return errWrapped
	}
	return nil
}

func (d *Delivery) GetOrders() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if d.IsEmptyId() {
		return ErrIdEmpty
	}
	tx := handler.Model(d).Select("deliveries.id,orders.id,orders.created_at,orders.fvat,orders.price,orders.payment_id,orders.delivery_id").Joins("left join orders on orders.delivery_id=deliveries.id").Where("deliveries.id=?", d.GetID()).Find(&d.Order)
	if tx.Error != nil {
		errWrapped := errors.Wrap(tx.Error, "Joining table delivery and orders")
		return errWrapped
	}
	return nil
}
