package service

import (
	"database-example/repo"

	"database-example/model"

)

type TourReviewService struct {
	TourReviewRepository *repo.TourReviewRepository
}



func (service *TourReviewService) Delete(idTour int) error {
    return service.TourReviewRepository.Delete(idTour)
}

func (service *TourReviewService) Create(tourReview *model.TourReview) error {
	err := service.TourReviewRepository.Create(tourReview)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourReviewService) FindAll() ([]model.TourReview, error) {
    return service.TourReviewRepository.FindAll()
}