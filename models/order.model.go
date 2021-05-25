package models

type Order struct {
	Base
	Fvat     bool
	Price    float32
	Payment  Payment
	Delivery Delivery
}

func (o *Order) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Delete(o).Error
}

func (o *Order) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(o).Error
}

func (o *Order) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(o.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.Find(o).Error
}

func (o *Order) AssignTo(p *Product) error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(o.ID) == 0 {
		return ErrIdEmpty
	}
}
