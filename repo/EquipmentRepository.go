package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EquipmentRepository) FindById(id int) (model.Equipment, error) {
	equipment := model.Equipment{}
	dbResult := repo.DatabaseConnection.First(&equipment, "id = ?", id)
	if dbResult != nil {
		return equipment, dbResult.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) CreateEquipemnt(equipment *model.Equipment) error {
	dbResult := repo.DatabaseConnection.Create(equipment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EquipmentRepository) FindAll() ([]model.Equipment, error) {
    var equipments []model.Equipment
    dbResult := repo.DatabaseConnection.Find(&equipments)
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    return equipments, nil
}