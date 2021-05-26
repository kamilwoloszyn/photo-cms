package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string `sql:"type:uuid;primary_key;default:uuid_generate_v4()" validate:"omitempty,uuid,required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `sql:"index" json:"deleted_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	if err != nil {
		errWrapped := fmt.Sprintf("Cannot generate new Id : %s", err.Error())
		return errors.New(errWrapped)
	}
	tx.Statement.SetColumn("id", id)
	return nil

}

func (b *Base) GetID() string {
	return b.ID
}

func (b *Base) SetID(newId string) error {
	if len(newId) == 0 {
		return errors.New("new id is empty")
	}
	b.ID = newId
	return nil
}

func (b *Base) SetCreatedAt(t time.Time) {
	b.CreatedAt = t.String()
}

func (b *Base) SetUpdatedAt(t time.Time) {
	b.UpdatedAt = t.String()
}

func (b *Base) SetDeletedAt(t time.Time) {
	b.DeletedAt = t.String()
}
