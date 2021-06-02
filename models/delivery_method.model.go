package models

import (
	"github.com/pkg/errors"
)

type DeliveryMethod struct {
	Base
	Name       string     `gorm:"not null"`
	FixedPirce float32    `gorm:"not null"`
	Delivery   []Delivery `gorm:"foreignKey:DeliveryMethodId"`
}

func (dm *DeliveryMethod) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if dm.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Create(dm).Error

}

func (dm *DeliveryMethod) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if dm.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Delete(dm).Error
}

func (dm *DeliveryMethod) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if dm.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.First(dm).Error
}

func (dm *DeliveryMethod) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if dm.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Save(dm).Error
}

func (dm *DeliveryMethod) AssignTo(d *Delivery) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if dm.IsEmptyId() {
		return ErrIdEmpty
	}
	d.DeliveryMethodId = dm.GetID()
	if err := d.UpdateInstance(); err != nil {
		errWrapped := errors.Wrap(err, "Updating instance of delivery model")
		return errWrapped
	}
	return nil
}

func (dm *DeliveryMethod) GetDeliveries() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if dm.IsEmptyId() {
		return ErrIdEmpty
	}
	tx := handler.Model(dm).Select("delivery_methods.id,deliveries.id,deliveries.shipped_via,deliveries.tracking_code,deliveries.destination_postal_code,deliveries.destination_country_region,deliveries.destination_address,deliveries.destination_city,deliveries.payment_method_id").Joins("left join deliveries on deliveries.delivery_method_id=delivery_methods.id").Where("delivery_methods.id=?", dm.GetID()).Find(&dm.Delivery)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
