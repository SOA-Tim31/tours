package service

import (
	"database-example/model"
	"database-example/repo"
	"encoding/json"
	"fmt"
	"os"
)

type StudentService struct {
	StudentRepo *repo.StudentRepository
}

func (service *StudentService) FindStudent(id string) (*model.Student, error) {
	student, err := service.StudentRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &student, nil
}

func (service *StudentService) Create(student *model.Student) error {
	err := service.StudentRepo.CreateStudent(student)
	jsonval, _ := json.Marshal(student)
	os.WriteFile("gas.json", jsonval, 0666)

	if err != nil {
		return err
	}
	return nil
}
