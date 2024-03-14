package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)


type TourService struct{
	TourRepository *repo.TourRepository
}



func (service *TourService) Create(tour *model.Tour) error {
	err := service.TourRepository.Create(tour)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) FindById(id string) (*model.Tour, error){
	tour, err := service.TourRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &tour, nil

}


func (service *TourService) FindAll() ([]model.Tour, error){
   return service.TourRepository.FindAll()
}