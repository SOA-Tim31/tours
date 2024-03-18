package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)


type CompetitionRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *CompetitionRepository) Create(competition *model.Competition) error {
	dbResult := repo.DatabaseConnection.Create(competition)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *CompetitionRepository) FindById(id int) (model.Competition, error) {
	competition := model.Competition{}
	
	dbResult := repo.DatabaseConnection.First(&competition, id)
	if dbResult != nil {
		return competition, dbResult.Error
	}
	return competition, nil
}

func (repo *CompetitionRepository) FindAll() ([]model.Competition, error){
	var competitions []model.Competition
	dbResult := repo.DatabaseConnection.Find(&competitions)
	if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    return competitions, nil

}