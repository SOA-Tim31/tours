package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DifficultyLevel int

const (

	Easy DifficultyLevel = iota
	Moderate 
	Hard 

)

type Status int

const (

	Draft Status = iota
	Published 
	Archived

)



type Tour struct {
	ID uuid.UUID `json:"id"`
	Name  string    `json:"name" gorm:"not null;type:string"`
	Description  string    `json:"desc" gorm:"type:string"`
	DifficultyLevel DifficultyLevel `json:"difficultylevel" gorm:"not null;type:int"`
	Status Status `json:"status" gorm:"not null;type:int"`
	Price float64 `json:"price" gorm:"not null;type:float"`
}

//count id before create tour
func (t *Tour) BeforeCreate(scope *gorm.DB) error {
	t.ID = uuid.New()
	return nil
}