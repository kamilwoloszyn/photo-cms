package models

type Image struct {
	Base
	Name     string
	FullPath string
	Size     uint32
}

func (i *Image) Create() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Create(i).Error
}

func (i *Image) FetchById() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.First(i).Error
}

func (i *Image) Delete() error {
	if handler == nil {
		return HandlerNotFound
	}
	return handler.Delete(i).Error
}
