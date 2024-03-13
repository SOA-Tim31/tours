package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type EquipmentService struct {
	EquipmentRepository *repo.EquipmentRepository
}

func (service *EquipmentService) FindEquipment(id string) (*model.Equipment, error) {
	equipment, err := service.EquipmentRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &equipment, nil
}

func (service *EquipmentService) Create(equipment *model.Equipment) error {
	err := service.EquipmentRepository.CreateEquipemnt(equipment)
	if err != nil {
		return err
	}
	return nil
}

func (service *EquipmentService) FindAllEquipment() ([]model.Equipment, error) {
    return service.EquipmentRepository.FindAll()
}