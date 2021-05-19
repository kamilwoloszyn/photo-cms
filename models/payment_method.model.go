package models

type PaymentMethod struct {
	Base
	Name        string
	Provier     string
	PostId      string
	KeyMd5      string
	ClientId    string
	OauthSecret string
}

func (p *PaymentMethod) FetchByID() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.First(p).Error
}

func (p *PaymentMethod) Delete() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Delete(p).Error
}

func (p *PaymentMethod) Create() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Create(p).Error
}
