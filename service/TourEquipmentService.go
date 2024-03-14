package service

import (
	"database-example/repo"
	"fmt"
	"database-example/model"

)

type TourEquipmentService struct {
	TourEquipmentRepository *repo.TourEquipmentRepository
}

func (service *TourEquipmentService) AddEquipmentToTourAsync(tourID string, equipmentID string) error {
	allTourEquipment, _ := service.TourEquipmentRepository.FindAll()

	connectionExists := false
	for _, te := range allTourEquipment {
		if te.TourId == tourID && te.EquipmentId == equipmentID {
			connectionExists = true
			break
		}
	}

	if !connectionExists {
		newTourEquipment := &model.TourEquipment{
			TourId:      tourID,
			EquipmentId: equipmentID,
			IsSelected:  true,
		}
		
		service.TourEquipmentRepository.CreateTourEquipemnt(newTourEquipment)
	}

	fmt.Println("Equipment added to tour successfully.")
	return nil
}

func (service *TourEquipmentService) DeleteTourEquipment(idTour, equipmentID string) error {
    return service.TourEquipmentRepository.DeleteTourEquipment(idTour, equipmentID)
}


