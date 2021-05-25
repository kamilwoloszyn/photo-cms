package models

type OptionValue struct {
	Base
	Value      string
	ExtraPrice float32
}

func (o *OptionValue) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.First(o).Error

}

func (o *OptionValue) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(o).Error
}

func (o *OptionValue) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(o.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.Delete(o).Error
}
