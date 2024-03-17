package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TourEquipmentHandler struct {
	TourEquipmentService *service.TourEquipmentService
}

func (handler *TourEquipmentHandler) CreateTourEquipemnt(writer http.ResponseWriter, req *http.Request) {
	var tourEq model.TourEquipment
	err := json.NewDecoder(req.Body).Decode(&tourEq)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourEquipmentService.TourEquipmentRepository.CreateTourEquipemnt(&tourEq)
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourEquipmentHandler) DeleteTourEquipment(writer http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    idTourParam := vars["idTour"]
    idEquipmentParam := vars["idEquipment"]

    idTour, err := strconv.Atoi(idTourParam)
    if err != nil {
        writer.WriteHeader(http.StatusBadRequest)
        writer.Write([]byte("Invalid tour ID"))
        return
    }

    idEquipment, err := strconv.Atoi(idEquipmentParam)
    if err != nil {
        writer.WriteHeader(http.StatusBadRequest)
        writer.Write([]byte("Invalid equipment ID"))
        return
    }

    err = handler.TourEquipmentService.DeleteTourEquipment(idTour, idEquipment)
    if err != nil {
        writer.WriteHeader(http.StatusInternalServerError)
        writer.Write([]byte("Error deleting tour equipment: " + err.Error()))
        return
    }

    writer.WriteHeader(http.StatusOK)
    writer.Write([]byte("Tour equipment deleted successfully"))
}

func (h *TourEquipmentHandler) GetTourEquipment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    idTourParam := vars["tourID"]
    if idTourParam == "" {
        http.Error(w, "Missing tour ID", http.StatusBadRequest)
        return
    }

    tourID, err := strconv.Atoi(idTourParam)
    if err != nil {
        http.Error(w, "Invalid tour ID", http.StatusBadRequest)
        return
    }

    equipment, err := h.TourEquipmentService.GetTourEquipment(tourID)
    if err != nil {
        http.Error(w, "Error fetching tour equipment: "+err.Error(), http.StatusInternalServerError)
        return
    }


    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(equipment)

}


