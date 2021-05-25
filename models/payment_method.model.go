package models

type PaymentMethod struct {
	Base
	Name        string
	Provider    string
	PosId       string
	KeyMd5      string
	ClientId    string
	OauthSecret string
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
