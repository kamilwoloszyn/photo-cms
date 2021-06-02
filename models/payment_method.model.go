package models

import (
	"github.com/kamilwoloszyn/photo-cms/pkg/checkers"
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
	if pm.IsEmptyId() {
		return ErrIdEmpty
	}
	if payMethodId := checkers.UuidGeneric(pm.GetID()); payMethodId.IsEmpty() {
		return ErrIdEmpty
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
