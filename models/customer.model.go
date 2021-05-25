package models

type Customer struct {
	Base
	City         string
	Address      string
	FirstName    string
	LastName     string
	PostalCode   string
	CompanyName  string
	PhoneNumber  string
	EmailAddress string
	Employed     bool
	NIP          string
	Regon        string
}

func (c *Customer) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}

	return handler.Create(c).Error
}
func (c *Customer) Delete() error {
	return handler.Delete(c).Error
}

func (c *Customer) SetId(id string) {
	if len(id) > 0 {
		c.ID = id
	}
}

func (c *Customer) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(c.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.Find(c).Error

}
