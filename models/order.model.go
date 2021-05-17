package models

type Order struct {
	Price    float64
	BaseCols Base
	Printing Print
}

func (o *Order) Delete() error {
	return handler.Delete(o).Error
}

func (o *Order) Save() error {
	return handler.Save(o).Error
}
