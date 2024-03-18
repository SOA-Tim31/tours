package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type CompetitionService struct {
	CompetitionRepository *repo.CompetitionRepository
}



func (service *CompetitionService) FindById(id int) (*model.Competition, error) {
	competition, err := service.CompetitionRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &competition, nil
}



func (service *CompetitionService) Create(competition *model.Competition) error {
	err := service.CompetitionRepository.Create(competition)
	if err != nil {
		return err
	}
	return nil
}



func (service *CompetitionService) FindAll() ([]model.Competition, error) {
    return service.CompetitionRepository.FindAll()
}