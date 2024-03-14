package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourEquipmentRepository struct {
	DatabaseConnection *gorm.DB
}



func (repo *TourEquipmentRepository) CreateTourEquipemnt(equipmentTour *model.TourEquipment) error {
	dbResult := repo.DatabaseConnection.Create(equipmentTour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourEquipmentRepository) FindAll() ([]model.TourEquipment, error) {
    var equipments []model.TourEquipment
    dbResult := repo.DatabaseConnection.Find(&equipments)
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    return equipments, nil
}

func (repo *TourEquipmentRepository) DeleteTourEquipment(idTour, equipmentID string) error {
    var equipment model.TourEquipment
    dbResult := repo.DatabaseConnection.Where("equipment_id = ? AND tour_id = ?", equipmentID, idTour).First(&equipment)
    if dbResult.Error != nil {
        return dbResult.Error
    }
    
    dbResult = repo.DatabaseConnection.Delete(&equipment)
    if dbResult.Error != nil {
        return dbResult.Error
    }
    
    println("Rows affected: ", dbResult.RowsAffected)
    return nil
}



