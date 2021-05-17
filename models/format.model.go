package models

type Format struct {
	BaseCols Base
	Name     string
}

func (f *Format) Delete() error {
	return handler.Delete(f).Error
}

func (f *Format) Save() error {
	return handler.Save(f).Error
}
