package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourObjectRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourObjectRepository) Create(object *model.TourObject) error {
	dbResult := repo.DatabaseConnection.Create(object)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourObjectRepository) FindAll() ([]model.TourObject, error){
	var tours []model.TourObject
	dbResult := repo.DatabaseConnection.Find(&tours)
	if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    return tours, nil

}