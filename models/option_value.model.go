package models

type OptionValue struct {
	Base
	Value      string
	ExtraPrice float32
}

func (o *OptionValue) FetchById() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.First(o).Error

}

func (o *OptionValue) Create() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Create(o).Error
}

func (o *OptionValue) Delete() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Delete(o).Error
}
