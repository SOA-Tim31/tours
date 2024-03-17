package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)



type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) Create(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) FindById(id int) (model.Tour, error) {
	tour := model.Tour{}
	
	dbResult := repo.DatabaseConnection.Preload("TourPoints").First(&tour, id)
	if dbResult != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

func (repo *TourRepository) FindByUserId(userId int) ([]model.Tour, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.Preload("TourPoints").Find(&tours, `"UserId" = ?`, userId)
	if dbResult.Error != nil {
        return nil, dbResult.Error
    }
	
    return tours, nil
}

func (repo *TourRepository) FindAll() ([]model.Tour, error){
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.Preload("TourPoints").Find(&tours)
	if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    return tours, nil

}