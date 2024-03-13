package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Equipment struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name" gorm:"not null;type:string"`
	Description  string    `json:"desc" gorm:"not null;type:string"`
}

func (eq *Equipment) BeforeCreate(scope *gorm.DB) error {
	eq.ID = uuid.New()
	return nil
}

func (Equipment) TableName() string {
    return "tours.equipment"
}