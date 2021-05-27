package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	Base
	PaymentDate     *time.Time
	PaymentAmount   float32
	PaymentError    bool
	PaymentFinished bool
	PaymentMethodId uuid.UUID
	Order           []Order `gorm:"foreignKey:PaymentId"`
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
	return handler.Delete(p).Error
}

func (p *Payment) FetchByID() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(p.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.First(p).Error
}
