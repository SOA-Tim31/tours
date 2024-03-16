package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TourReviewHandler struct {
	TourReviewService *service.TourReviewService
}

func (handler *TourReviewHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tour model.TourReview
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourReviewService.TourReviewRepository.Create(&tour)
	if err != nil {
		println("Error while creating a new review")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourReviewHandler) Delete(writer http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    idTourParam := vars["id"]

	println(idTourParam)

    if idTourParam == ""  {
        writer.WriteHeader(http.StatusBadRequest)
        writer.Write([]byte("Missing tour ID "))
        return
    }
	
	idTour, err := strconv.Atoi(idTourParam)
    if err != nil {
        writer.WriteHeader(http.StatusBadRequest)
        writer.Write([]byte("Invalid tour ID"))
        return
    }

    err = handler.TourReviewService.TourReviewRepository.Delete(idTour)
    if err != nil {
        writer.WriteHeader(http.StatusInternalServerError)
        writer.Write([]byte("Error deleting tour review: " + err.Error()))
        return
    }

    writer.WriteHeader(http.StatusOK)
    writer.Write([]byte("Tour review deleted successfully"))
}

func (h *TourReviewHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    reviews, err := h.TourReviewService.FindAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(reviews)
}