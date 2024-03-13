package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	TourService *service.TourService
}


func (handler *TourHandler) CreateTour(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		println("Error while parsing json")
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

    tour, err := handler.TourService.FindById(tourID)
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