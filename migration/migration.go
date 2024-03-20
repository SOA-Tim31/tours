package migration

import (
	"database-example/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {

    if err := db.AutoMigrate(&model.Student{}, &model.Equipment{}, &model.Tour{},&model.TourEquipment{},&model.TourReview{},&model.TourObject{},&model.TourPoint{}, &model.Competition{}); err != nil {
        return err
    }

    if err := db.Exec(`ALTER TABLE "tours"."Tours" ALTER COLUMN "TourCharacteristics" SET DEFAULT '[]'::jsonb`).Error; err != nil {

        return err
    }
    return nil
}
