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
		return HandlerNotFound
	}
	return handler.Delete(o).Error
}

func (o *Order) Create() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Create(o).Error
}

func (o *Order) FetchById() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Find(o).Error
}
