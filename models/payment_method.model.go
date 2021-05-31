package models

type PaymentMethod struct {
	Base
	Name        string    `gorm:"not null"`
	Provider    string    `gorm:"not null"`
	PosId       string    `gorm:"not null"`
	KeyMd5      string    `gorm:"not null"`
	ClientId    string    `gorm:"not null"`
	OauthSecret string    `gorm:"not null"`
	Payment     []Payment `gorm:"foreignKey:PaymentMethodId"`
}

func (p *PaymentMethod) FetchByID() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.First(p).Error
}

func (p *PaymentMethod) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Delete(p).Error
}

func (p *PaymentMethod) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(p.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.Create(p).Error
}

func (p *PaymentMethod) UpdateAll() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(p).Error
}
