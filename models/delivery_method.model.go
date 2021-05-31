package models

type DeliveryMethod struct {
	Base
	Name       string     `gorm:"not null"`
	FixedPirce float32    `gorm:"not null"`
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

func (d *DeliveryMethod) UpdateAll() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(d).Error
}
