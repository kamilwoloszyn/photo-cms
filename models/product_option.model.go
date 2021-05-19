package models

type ProductOption struct {
	Base
	OptionValues []OptionValue
	Products     []Product
}

func (p *ProductOption) Create() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Create(p).Error
}
func (p *ProductOption) Delete() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Delete(p).Error
}
func (p *ProductOption) FetchById() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.First(p).Error

}
