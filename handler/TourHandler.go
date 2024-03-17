package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	TourService *service.TourService
}


func (handler *TourHandler) CreateTour(writer http.ResponseWriter, req *http.Request) {
	var (tour model.Tour
	    tourDTO model.TourDTO
	)

    // body, _:= io.ReadAll(req.Body)
	// fmt.Println("Request body:", string(body))
	
    
    // // jsonval, _ := json.Marshal(string(body))
	//  err:= json.Unmarshal(body, &tour)
	//  if err != nil{
	// 	fmt.Println(err)
	//  }
	//  fmt.Println(tour)
	//  os.Exit(1)
    
	err := json.NewDecoder(req.Body).Decode(&tourDTO)
	switch tourDTO.Status{
	case "Draft":
		tour.Status=0
	case "Published":
		tour.Status=1
	}

	switch tourDTO.DifficultyLevel{
	case "Easy":
		tour.DifficultyLevel=0
	case "Moderate":
		tour.DifficultyLevel=1
	case "Hard":
		tour.DifficultyLevel=2
	}
	tour.TourCharacteristics = tourDTO.TourCharacteristics
	fmt.Println(tour.TourCharacteristics)
    tour.ID = tourDTO.ID
	tour.Name = tourDTO.Name
	tour.Price = tourDTO.Price
	tour.ArchivedDateTime = tourDTO.ArchivedDateTime
	tour.PublishedDateTime = tourDTO.PublishedDateTime
	tour.Description = tourDTO.Description
	tour.Tags = tourDTO.Tags
	tour.UserId = tourDTO.UserId
	
	

	if err != nil {
		println("Error while parsing json", err)
		fmt.Println(err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourService.Create(&tour)
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}


func (handler *TourHandler) FindByID(writer http.ResponseWriter, req *http.Request) {
    
    tourID := mux.Vars(req)["id"]

	idTour, err := strconv.Atoi(tourID)
    if err != nil {
        writer.WriteHeader(http.StatusBadRequest)
        writer.Write([]byte("Invalid tour ID"))
        return
    }

    tour, err := handler.TourService.FindById(idTour)
    if err != nil {
        http.Error(writer, "Failed to find tour", http.StatusInternalServerError)
        return
    }

	
    
    writer.Header().Set("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(tour)
}


func (handler *TourHandler) FindAllTours(writer http.ResponseWriter, req *http.Request){
	tours, err := handler.TourService.FindAll()
	
	
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
	}

	writer.Header().Set("Content-Type", "application/json")
    encoder := json.NewEncoder(writer)
    for _, tour := range tours {
        if err := encoder.Encode(tour); err != nil {
            http.Error(writer, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}


func (handler *TourHandler) FindByUserId(writer http.ResponseWriter, req *http.Request){

	userID := mux.Vars(req)["userId"]

	idUser, err := strconv.Atoi(userID)
    if err != nil {
        writer.WriteHeader(http.StatusBadRequest)
        writer.Write([]byte("Invalid tour ID"))
        return
    }
	tours, err := handler.TourService.FindByUserId(idUser)
	
	
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
	}

	writer.Header().Set("Content-Type", "application/json")
    encoder := json.NewEncoder(writer)
	encoder.Encode(tours)
    // for _, tour := range tours {
    //     if err := encoder.Encode(tour); err != nil {
    //         http.Error(writer, err.Error(), http.StatusInternalServerError)
    //         return
    //     }
    // }
}