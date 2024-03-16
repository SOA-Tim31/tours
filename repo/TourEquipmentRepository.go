package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourEquipmentRepository struct {
	DatabaseConnection *gorm.DB
}



func (repo *TourEquipmentRepository) CreateTourEquipemnt(equipmentTour *model.TourEquipment) error {
    equipmentTour.IsSelected = true 
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

func (repo *TourEquipmentRepository) GetTourEquipment(tourID int) ([]model.TourEquipment, error) {
    allTourEquipment, err := repo.FindAll() 
    if err != nil {
        return nil, err
    }

    var tourEquipmentForTour []model.TourEquipment
    for _, te := range allTourEquipment {
        if te.TourId == tourID {
            tourEquipmentForTour = append(tourEquipmentForTour, te)
        }
    }

    return tourEquipmentForTour, nil
}

// func (repo *TourEquipmentRepository) UpdateTourEquipment(equipment *model.TourEquipment) error {
//     var existingEquipment model.TourEquipment
//     dbResult := repo.DatabaseConnection.First(&existingEquipment, equipment.Id)
//     if dbResult.Error != nil {
//         return dbResult.Error
//     }

//     existingEquipment.TourId = equipment.TourId
//     existingEquipment.EquipmentId = equipment.EquipmentId
//     existingEquipment.IsSelected = equipment.IsSelected

//     dbResult = repo.DatabaseConnection.Save(&existingEquipment)
//     if dbResult.Error != nil {
//         return dbResult.Error
//     }

//     println("Rows affected: ", dbResult.RowsAffected)
//     return nil
// }



func (repo *TourEquipmentRepository) DeleteTourEquipment(idTour, equipmentID int) error {
    var equipment model.TourEquipment
    equipment.IsSelected = false 
    dbResult := repo.DatabaseConnection.First(&equipment,idTour,equipmentID)
    if dbResult.Error != nil {
        return dbResult.Error
    }
    
    equipment.IsSelected = false
    
    dbResult = repo.DatabaseConnection.Save(&equipment)
    if dbResult.Error != nil {
        return dbResult.Error
    }
    
    println("Rows affected: ", dbResult.RowsAffected)
    return nil
}



