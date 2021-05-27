package models

type DeliveryMethod struct {
	Base
	Name       string
	FixedPirce float32
	Delivery   []Delivery `gorm:"foreignKey:DeliveryMethodId"`
}

func (d *DeliveryMethod) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(d).Error

}

func (d *DeliveryMethod) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Delete(d).Error
}

func (d *DeliveryMethod) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(d.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.First(d).Error
}
