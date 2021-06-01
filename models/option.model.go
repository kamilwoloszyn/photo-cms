package models

import "github.com/pkg/errors"

type Option struct {
	Base
	Name        string        `gorm:"not null"`
	OptionValue []OptionValue `gorm:"foreignKey:OptionId"`
}

func (o *Option) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.First(o).Error
}

func (o *Option) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Delete(o).Error
}

func (o *Option) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(o.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.Create(o).Error
}

func (o *Option) UpdateInstance() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(o).Error
}

func (o *Option) GetOptionValues() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	err := handler.Model(o).Select("options.id, options.name,option_values.value").Joins("left join option_values on option_values.option_id = options.id").Where("options.id= ?", o.GetID()).Find(&o.OptionValue).Error
	if err != nil {
		errWrapped := errors.Wrap(err, "GetOptionValues : Join tables")
		return errWrapped
	}
	return nil
}

func (o *Option) AssignTo(ov *OptionValue) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if o.IsEmptyId() || ov.IsEmptyId() {
		return ErrIdEmpty
	}
	ov.ID = o.GetID()
	if err := ov.UpdateInstance(); err != nil {
		errWrapped := errors.Wrap(err, "Updating instance")
		return errWrapped
	}
	return nil
}
