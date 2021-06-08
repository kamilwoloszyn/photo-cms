package models

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type OptionValue struct {
	Base
	Value         string
	ExtraPrice    float32
	OptionId      uuid.UUID
	ProductOption []ProductOption `gorm:"foreginKey:OptionValueId"`
}

func (o *OptionValue) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if o.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.First(o).Error

}

func (o *OptionValue) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if o.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Create(o).Error
}

func (o *OptionValue) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if o.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Delete(o).Error
}

func (o *OptionValue) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if o.IsEmptyId() {
		return ErrIdEmpty
	}
	return handler.Save(o).Error
}

func (o *OptionValue) AssignTo(po *ProductOption) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if o.IsEmptyId() {
		return ErrIdEmpty
	}
	po.OptionValueId = o.GetID()
	if err := po.UpdateInstance(); err != nil {
		errWrapped := errors.Wrap(err, "Updating instance Product Option")
		return errWrapped
	}
	return nil
}

func (o *OptionValue) GetProductOptions() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if o.IsEmptyId() {
		return ErrIdEmpty
	}
	tx := handler.Model(o).Select("option_values.id,product_options.id,product_options.product_id,product_options.option_value_id").Joins("left join product_options on product_options.option_value_id=option_values.id").Where("option_values.id=?", o.GetID()).Find(&o.ProductOption)
	if tx.Error != nil {
		errWrapped := errors.Wrap(tx.Error, "Join OptionValue with ProductOption")
		return errWrapped
	}
	return nil
}
