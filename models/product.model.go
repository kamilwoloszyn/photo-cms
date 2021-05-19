package models

type Product struct {
	Base
	UnitPrice   float32
	ProductName string
	Quantity    uint32
	Category    []Category
	Images      []Image
	Customer    []Customer
	Order       Order
}

func (p *Product) Create() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Create(p).Error
}
func (p *Product) Delete() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Delete(p).Error
}
func (p *Product) FetchByID() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.First(p).Error
}
