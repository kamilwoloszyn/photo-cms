package models

import (
	"github.com/pkg/errors"
)

type PaymentMethod struct {
	Base
	Name        string    `gorm:"not null"`
	Provider    string    `gorm:"not null"`
	PosId       string    `gorm:"not null"`
	KeyMd5      string    `gorm:"not null"`
	ClientId    string    `gorm:"not null"`
	OauthSecret string    `gorm:"not null"`
	Payment     []Payment `gorm:"foreignKey:PaymentMethodId"`
}

func (pm *PaymentMethod) FetchByID() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if pm.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.First(pm).Error
}

func (pm *PaymentMethod) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if pm.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Delete(pm).Error
}

func (pm *PaymentMethod) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(pm).Error
}

func (pm *PaymentMethod) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if pm.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Save(pm).Error
}

func (pm *PaymentMethod) AssignTo(p *Payment) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if pm.IsEmptyId() {
		return ErrIdEmpty
	}
	p.PaymentMethodId = pm.GetID()
	if err := p.UpdateInstance(); err != nil {
		errWrapped := errors.Wrap(err, "Updating instance Payment")
		return errWrapped
	}
	return nil
}

func (pm *PaymentMethod) GetPayments() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if pm.IsEmptyId() {
		return ErrIdEmpty
	}

	tx := handler.Model(pm).Select("payment_methods.id,payments.id,payments.created_at,payments.updated_at,payments.deleted_at,payments.payment_date,payments.payment_amount,payments.payment_error,payments.payment_finished,payments.payment_method_id").Joins("left join payments on payments.payment_method_id=payment_methods.id").Where("payment_methods.id=?", pm.GetID()).Find(&pm.Payment)
	if tx.Error != nil {
		errWrapped := errors.Wrap(tx.Error, "Getting payment methods")
		return errWrapped
	}
	return nil
}
