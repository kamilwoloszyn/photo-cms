package models

type Image struct {
	Base
	Name     string
	FullPath string
	Size     uint32
}

func (i *Image) Create() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.Create(i).Error
}

func (i *Image) FetchById() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	return handler.First(i).Error
}

func (i *Image) Delete() error {
	if handler == nil {
		return ErrHandlerNotFound
	}
	if len(i.ID) == 0 {
		return ErrIdEmpty
	}
	return handler.Delete(i).Error
}
