package migration

import (
	"database-example/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
    if err := db.AutoMigrate(&model.Student{}, &model.Equipment{}, &model.Tour{},&model.TourEquipment{}); err != nil {
        return err
    }
    return nil
}
