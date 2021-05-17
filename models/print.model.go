package models

type Print struct {
	BaseCols   Base
	Paper      PaperType
	FormatType Format
}

func (p *Print) New(paperType PaperType, formatType Format) {
	p.FormatType = formatType
	p.Paper = paperType
}

func (p *Print) Delete() {
	handler.Delete(p)
}
