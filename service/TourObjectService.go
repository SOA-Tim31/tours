package service

import (
	"database-example/model"
	"database-example/repo"
)

type TourObjectService struct {
	TourObjectRepository *repo.TourObjectRepository
}

func (service *TourObjectService) Create(object *model.TourObject) error {
	err := service.TourObjectRepository.Create(object)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourObjectService) FindAll() ([]model.TourObject, error){
	return service.TourObjectRepository.FindAll()
 }