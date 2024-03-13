package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"

)

type EquipmentHandler struct {
    EquipmentService *service.EquipmentService
}

func NewEquipmentHandler(equipmentService *service.EquipmentService) *EquipmentHandler {
    return &EquipmentHandler{
        EquipmentService: equipmentService,
    }
}

func (h *EquipmentHandler) FindEquipmentHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "Missing ID parameter", http.StatusBadRequest)
        return
    }

    equipment, err := h.EquipmentService.FindEquipment(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(equipment)
}

func (h *EquipmentHandler) CreateEquipmentHandler(w http.ResponseWriter, r *http.Request) {
    var equipment model.Equipment
    err := json.NewDecoder(r.Body).Decode(&equipment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = h.EquipmentService.Create(&equipment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (h *EquipmentHandler) FindAllEquipmentHandler(w http.ResponseWriter, r *http.Request) {
    equipments, err := h.EquipmentService.FindAllEquipment()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(equipments)
}