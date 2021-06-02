package models

import "github.com/pkg/errors"

type Category struct {
	Base
	CategoryName string    `gorm:"not null"`
	Product      []Product `gorm:"foreignKey:CategoryId"`
}

func (c *Category) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.First(c).Error
}

func (c *Category) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Create(c).Error
}

func (c *Category) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Delete(c).Error
}

func (c *Category) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Save(c).Error
}

func (c *Category) AssignTo(p *Product) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if c.IsEmptyId() {
		return ErrIdEmpty
	}
	p.CategoryId = c.GetID()
	if err := p.UpdateInstance(); err != nil {
		errWrapped := errors.Wrap(err, "Update instance product")
		return errWrapped
	}
	return nil
}
