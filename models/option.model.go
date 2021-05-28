package models

type Option struct {
	Base
	Name        string
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

func (o *Option) UpdateAll() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(o).Error
}
