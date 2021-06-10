package models

import (
	"time"

	"github.com/pkg/errors"
)

type Payment struct {
	Base
	PaymentDate     *time.Time
	PaymentAmount   float32 `gorm:"not null"`
	PaymentError    bool    `gorm:"default:false"`
	PaymentFinished bool    `gorm:"default:false"`
	PaymentMethodId string  `gorm:"type:uuid;default:null"`
	Order           Order   `gorm:"foreignKey:PaymentId"`
}

func (p *Payment) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(p).Error
}

func (p *Payment) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if p.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Delete(p).Error
}

func (p *Payment) FetchByID() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if p.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.First(p).Error
}

func (p *Payment) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if p.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Save(p).Error
}

func (p *Payment) GetPaymentMethodDetails(pm *PaymentMethod) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(p.PaymentMethodId) == 0 {
		return ErrIdEmpty
	}
	pm.SetID(p.PaymentMethodId)
	if err := pm.FetchByID(); err != nil {
		errWrapped := errors.Wrap(err, "Fetching payment method")
		return errWrapped
	}
	return nil
}

func (p *Payment) AssignTo(o *Order) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if p.IsEmptyId() {
		return ErrIdEmpty
	}
	paymentId := p.GetID()
	o.PaymentId = &paymentId
	if err := o.UpdateInstance(); err != nil {
		errWrapped := errors.Wrap(err, "Updating payment instance")
		return errWrapped
	}
	return nil
}

func (p *Payment) GetOrders() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	tx := handler.Model(p).Select("payments.id,orders.id,orders.created_at,orders.updated_at,orders.deleted_at,orders.fvat,orders.price,orders.payment_id,orders.delivery_id").Joins("left join orders on orders.payment_id=payments.id").Where("payments.id=?", p.GetID()).Find(&p.Order)
	if tx.Error != nil {
		errWrapped := errors.Wrap(tx.Error, "Getting orders")
		return errWrapped
	}
	return nil
}
