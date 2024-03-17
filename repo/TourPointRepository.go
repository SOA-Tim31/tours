package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)



type TourPointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourPointRepository) Create(tourPoint *model.TourPoint) error {
	dbResult := repo.DatabaseConnection.Create(tourPoint)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourPointRepository) FindById(id int) (model.TourPoint, error) {
	tourPoint := model.TourPoint{}
	dbResult := repo.DatabaseConnection.First(&tourPoint, id)
	if dbResult != nil {
		return tourPoint, dbResult.Error
	}
	return tourPoint, nil
}


func (repo *TourPointRepository) FindAll() ([]model.TourPoint, error){
	var tourPoints []model.TourPoint
	dbResult := repo.DatabaseConnection.Find(&tourPoints)
	if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    return tourPoints, nil

}

func (repo *TourPointRepository) FindByTourId(tourId int) ([]model.TourPoint, error){
	var tourPoints []model.TourPoint
	dbResult := repo.DatabaseConnection.Find(&tourPoints, `"TourId" = ?`, tourId)
	if dbResult.Error != nil {
        return nil, dbResult.Error
    }
	// jsonval, _ := json.Marshal(tourPoints)
	// os.WriteFile("gas.json", jsonval, 0666)

    return tourPoints, nil

}