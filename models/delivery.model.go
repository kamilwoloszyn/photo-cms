package models

import (
	"github.com/pkg/errors"

	"github.com/google/uuid"
)

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
	if d.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Create(d).Error
}

func (d *Delivery) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(d).Error
}

func (d *Delivery) GetDeliveryMethodDetails(deliveryMethod *DeliveryMethod) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if d.IsEmptyId() || len(d.DeliveryMethodId) == 0 {
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
