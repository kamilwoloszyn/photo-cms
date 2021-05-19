package models

type DeliveryMethod struct {
	Base
	Name       string
	FixedPirce float32
}

func (d *DeliveryMethod) Create() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Create(d).Error

}

func (d *DeliveryMethod) Delete() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Delete(d).Error
}

func (d *DeliveryMethod) FetchById() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.First(d).Error
}
