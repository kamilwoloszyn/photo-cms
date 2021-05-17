package models

type PaperType struct {
	Name string
}

func (p *PaperType) Delete() error {
	return handler.Delete(p).Error
}

func (p *PaperType) Save() error {
	return handler.Save(p).Error
}
