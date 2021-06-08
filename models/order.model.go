package models

import (
	"github.com/google/uuid"
	"github.com/kamilwoloszyn/photo-cms/pkg/checkers"
	"github.com/pkg/errors"
)

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
	if o.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Delete(o).Error
}

func (o *Order) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if o.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Create(o).Error
}

func (o *Order) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if o.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Find(o).Error
}

func (o *Order) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if o.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Save(o).Error
}

func (o *Order) GetProducts() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if o.IsEmptyId() {
		return ErrIdEmpty
	}

	tx := handler.Model(o).Select("orders.id,products.id,products.created_at,products.updated_at,products.deleted_at,products.unit_price,products.product_name,products.quantity,products.category_id,products.image_id,products.customer_id,products.order_id").Joins("left join products on products.order_id=orders.id").Where("orders.id=?", o.GetID()).Find(&o.Product)
	if tx.Error != nil {
		errWrapped := errors.Wrap(tx.Error, "Join orders with products")
		return errWrapped
	}

	return nil
}

func (o *Order) AssignTo(p *Product) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if o.IsEmptyId() {
		return ErrIdEmpty
	}
	p.OrderId = o.GetID()
	if err := p.UpdateInstance(); err != nil {
		errWrapped := errors.Wrap(err, "Updating order instance")
		return errWrapped
	}
	return nil
}

func (o *Order) GetPaymentDetails(p *Payment) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if paymentId := checkers.UuidGeneric(o.PaymentId); paymentId.IsEmpty() {
		return ErrIdEmpty
	}
	p.SetID(o.PaymentId)
	if err := p.FetchByID(); err != nil {
		errWrapped := errors.Wrap(err, "Fetching product instance")
		return errWrapped
	}
	return nil
}

func (o *Order) GetDeliveryDetails(d *Delivery) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if deliveryId := checkers.UuidGeneric(o.DeliveryId); deliveryId.IsEmpty() {
		return ErrIdEmpty
	}
	d.SetID(o.DeliveryId)
	if err := d.FetchById(); err != nil {
		errWrapped := errors.Wrap(err, "Fetching Delivery")
		return errWrapped
	}
	return nil
}
