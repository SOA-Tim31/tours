package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
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

    if idTourParam == "" || idEquipmentParam == "" {
        writer.WriteHeader(http.StatusBadRequest)
        writer.Write([]byte("Missing tour ID or equipment ID"))
        return
    }

    err := handler.TourEquipmentService.DeleteTourEquipment(idTourParam, idEquipmentParam)
    if err != nil {
        writer.WriteHeader(http.StatusInternalServerError)
        writer.Write([]byte("Error deleting tour equipment: " + err.Error()))
        return
    }

    writer.WriteHeader(http.StatusOK)
    writer.Write([]byte("Tour equipment deleted successfully"))
}



