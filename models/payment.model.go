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
		return HandlerNotFound
	}
	return handler.Create(p).Error
}

func (p *Payment) Delete() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Delete(p).Error
}

func (p *Payment) FetchByID() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.First(p).Error
}
