package models

type Category struct {
	Base
	CategoryName string
	Product      []Product `gorm:"foreignKey:CategoryId"`
}

func (c *Category) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.First(c).Error
}

func (c *Category) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(c).Error
}

func (c *Category) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(c.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.Delete(c).Error
}

func (c *Category) UpdateAll() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Save(c).Error
}
