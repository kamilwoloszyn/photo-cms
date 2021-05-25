package models

import "time"

type Payment struct {
	Base
	PaymentDate     *time.Time
	PaymentAmount   float32
	PaymentMethods  []PaymentMethod
	PaymentError    bool
	PaymentFinished bool
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
