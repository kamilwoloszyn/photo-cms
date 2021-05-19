package models

type Delivery struct {
	Base
	ShippedVia               string
	TrackingCode             string
	DestinationPostalCode    string
	DestinationConturyRegion string
	DestinationAddress       string
	DestinationCity          string
	DeliveryMethods          []DeliveryMethod
}

func (d *Delivery) FetchById() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.First(d).Error
}

func (d *Delivery) Delete() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Delete(d).Error
}

func (d *Delivery) Create() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Create(d).Error
}
