package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	Base
	PaymentDate     *time.Time
	PaymentAmount   float32 `gorm:"not null"`
	PaymentError    bool    `gorm:"default:false"`
	PaymentFinished bool    `gorm:"default:false"`
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

func (p *Payment) UpdateAll() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(p).Error
}
