package model


import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type TourEquipment struct{
	ID    uuid.UUID `json:"id"`
	TourId string `json:"tour_id"`
	EquipmentId string `json:"equipment_id"`
	IsSelected bool `json:"is_selected"`
}

func (t *TourEquipment) BeforeCreate(scope *gorm.DB) error {
	t.ID = uuid.New()
	return nil
}

func (TourEquipment) TableName() string {
    return "tours.tourequipments"
}
