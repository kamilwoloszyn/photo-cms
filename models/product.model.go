package models

import "github.com/google/uuid"

type Product struct {
	Base
	UnitPrice     float32
	ProductName   string
	Quantity      uint32
	ProductOption []ProductOption `gorm:"foreignKey:ProductId"`
	CategoryId    uuid.UUID
	ImageId       uuid.UUID
	CustomerId    uuid.UUID
	OrderId       uuid.UUID
}

func (p *Product) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(p).Error
}
func (p *Product) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Delete(p).Error
}
func (p *Product) UpdateAll() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(p).Error
}
func (p *Product) FetchByID() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(p.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.First(p).Error
}
func (p *Product) AssignTo(o *Order) error {
	if len(o.ID) == 0 {
		return ErrIdEmpty
	}
	p.OrderId = o.GetID()
	return nil
}
