package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
)

type TourObjectHandler struct{
	TourObjectService *service.TourObjectService
}

func (handler *TourObjectHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var object model.TourObject
	err := json.NewDecoder(req.Body).Decode(&object)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourObjectService.TourObjectRepository.Create(&object)
	if err != nil {
		println("Error while creating a new object")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (h *TourObjectHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    reviews, err := h.TourObjectService.FindAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(reviews)
}