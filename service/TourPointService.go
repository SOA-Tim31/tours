package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type TourPointService struct {
	TourPointRepository *repo.TourPointRepository
}

func (service *TourPointService) FindById(id int) (*model.TourPoint, error) {
	tourPoint, err := service.TourPointRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &tourPoint, nil
}

func (service *TourPointService) Create(tourPoint *model.TourPoint) error {
	err := service.TourPointRepository.Create(tourPoint)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourPointService) FindAll() ([]model.TourPoint, error) {
    return service.TourPointRepository.FindAll()
}

func (service *TourPointService) FindByTourId(tourId int) ([]model.TourPoint, error) {
    return service.TourPointRepository.FindByTourId(tourId)
}