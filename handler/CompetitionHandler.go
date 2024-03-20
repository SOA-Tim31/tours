package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CompetitionHandler struct {
	CompetitionService *service.CompetitionService
}


func (h *CompetitionHandler) CreateCompetition(w http.ResponseWriter, r *http.Request) {
    var (competition model.Competition
         competitionDTO model.CompetitionDTO   
    )

    err := json.NewDecoder(r.Body).Decode(&competitionDTO)

    switch competitionDTO.Status{
	case "Open":
		competition.Status=0
	case "Close":
		competition.Status=1
	}


    competition.ID = competitionDTO.ID
    competition.TourID = competitionDTO.TourID
    competition.StartDate = competitionDTO.StartDate
    competition.Duration = competitionDTO.Duration
    competition.CompetitionApplies = competitionDTO.CompetitionApplies
     
  



    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = h.CompetitionService.Create(&competition)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}


func (handler *CompetitionHandler) FindByID(writer http.ResponseWriter, req *http.Request) {
    
    CompetitionID := mux.Vars(req)["id"]

	idCompetition, err := strconv.Atoi(CompetitionID)
    if err != nil {
        writer.WriteHeader(http.StatusBadRequest)
        writer.Write([]byte("Invalid tour ID"))
        return
    }

    competition, err := handler.CompetitionService.FindById(idCompetition)
    if err != nil {
        http.Error(writer, "Failed to find tour", http.StatusInternalServerError)
        return
    }

	
    
    writer.Header().Set("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(competition)
}


func (handler *CompetitionHandler) FindAllCompetitions(writer http.ResponseWriter, req *http.Request){
	competitions, err := handler.CompetitionService.FindAll()
	
	
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
	}

	writer.Header().Set("Content-Type", "application/json")
    encoder := json.NewEncoder(writer)
	encoder.Encode(competitions)
}