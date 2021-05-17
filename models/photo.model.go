package models

import (
	"os"
	"path"
)

type Photo struct {
	ID              uint64 `gorm:"primaryKey"`
	Location        string
	FileName        string
	PaperTypeId     string
	ColorCorrection bool
	UserID          string
	Paper           *PaperType
}

func (p *Photo) New(userId string) {
	tmpPath := os.TempDir()
	p.Location = path.Join(tmpPath, p.UserID, p.FileName)

}

func (p *Photo) Save() error {
	return handler.Save(p).Error
}

func (p *Photo) Delete() error {
	return handler.Delete(p).Error
}
