package models

type Option struct {
	Base
	Name          string
	OptionsValues []OptionValue
}

func (o *Option) FetchById() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.First(o).Error
}

func (o *Option) Delete() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Delete(o).Error
}

func (o *Option) Create() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Create(o).Error
}
