package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TourPointHandler struct {
	TourPointService *service.TourPointService
}


func (h *TourPointHandler) CreateTourPoint(w http.ResponseWriter, r *http.Request) {
    var tourPoint model.TourPoint
    err := json.NewDecoder(r.Body).Decode(&tourPoint)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = h.TourPointService.Create(&tourPoint)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}


func (handler *TourPointHandler) FindByID(writer http.ResponseWriter, req *http.Request) {
    
    tourPointID := mux.Vars(req)["id"]

	idTourPoint, err := strconv.Atoi(tourPointID)
    if err != nil {
        writer.WriteHeader(http.StatusBadRequest)
        writer.Write([]byte("Invalid tour ID"))
        return
    }

    tourPoint, err := handler.TourPointService.FindById(idTourPoint)
    if err != nil {
        http.Error(writer, "Failed to find tour", http.StatusInternalServerError)
        return
    }

	
    
    writer.Header().Set("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(tourPoint)
}


func (handler *TourPointHandler) FindAllTourPoints(writer http.ResponseWriter, req *http.Request){
	tourPoints, err := handler.TourPointService.FindAll()
	
	
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
	}

	writer.Header().Set("Content-Type", "application/json")
    encoder := json.NewEncoder(writer)
    for _, tourPoint := range tourPoints {
        if err := encoder.Encode(tourPoint); err != nil {
            http.Error(writer, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}


func (handler *TourPointHandler) FindByTourId(writer http.ResponseWriter, req *http.Request){

	tourID := mux.Vars(req)["tourId"]
    
	idTour, err := strconv.Atoi(tourID)
    if err != nil {
        writer.WriteHeader(http.StatusBadRequest)
        writer.Write([]byte("Invalid tour ID"))
        return
    }
	tourPoints, err := handler.TourPointService.FindByTourId(idTour)
	
	
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
	}

	writer.Header().Set("Content-Type", "application/json")
    encoder := json.NewEncoder(writer)
	encoder.Encode(tourPoints)
    // for _, tour := range tours {
    //     if err := encoder.Encode(tour); err != nil {
    //         http.Error(writer, err.Error(), http.StatusInternalServerError)
    //         return
    //     }
    // }
}