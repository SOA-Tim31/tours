package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourReviewRepository struct {
	DatabaseConnection *gorm.DB
}



func (repo *TourReviewRepository) Create(review *model.TourReview) error {
	dbResult := repo.DatabaseConnection.Create(review)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourReviewRepository) FindAll() ([]model.TourReview, error) {
    var reviews []model.TourReview
    dbResult := repo.DatabaseConnection.Find(&reviews)
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    return reviews, nil
}

func (repo *TourReviewRepository) Delete(idTour int) error {
    var review model.TourReview
	dbResult := repo.DatabaseConnection.First(&review, idTour)
    if dbResult.Error != nil {
        return dbResult.Error
    }

    dbResult = repo.DatabaseConnection.Delete(&review)
    if dbResult.Error != nil {
        return dbResult.Error 
    }

    println("Rows affected: ", dbResult.RowsAffected)
    return nil
}


