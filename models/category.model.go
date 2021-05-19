package models

type Category struct {
	Base
	CategoryName string
}

func (c *Category) FetchById() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.First(c).Error
}

func (c *Category) Create() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Create(c).Error
}

func (c *Category) Delete() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Delete(c).Error
}
