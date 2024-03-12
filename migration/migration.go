package migration

import (
    "database-example/model"
    "gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
    if err := db.AutoMigrate(&model.Student{}, &model.Equipment{}); err != nil {
        return err
    }
    return nil
}
