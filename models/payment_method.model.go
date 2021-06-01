package models

import "github.com/pkg/errors"

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
	return handler.First(pm).Error
}

func (pm *PaymentMethod) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Delete(pm).Error
}

func (pm *PaymentMethod) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(pm.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.Create(pm).Error
}

func (pm *PaymentMethod) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(pm).Error
}

func (pm *PaymentMethod) AssignTo(p *Payment) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if pm.IsEmptyId() || p.IsEmptyId() {
		return ErrIdEmpty
	}
	p.PaymentMethodId = pm.GetID()
	if err := p.UpdateInstance(); err != nil {
		errWrapped := errors.Wrap(err, "Updating instance Payment")
		return errWrapped
	}
	return nil
}
