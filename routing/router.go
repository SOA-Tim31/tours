package routing

import (
	"database-example/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(handler *handler.EquipmentHandler, tourHandler *handler.TourHandler,tourEqHandler *handler.TourEquipmentHandler,reviewHandler *handler.TourReviewHandler, tourPointHandler *handler.TourPointHandler, competitionHandler *handler.CompetitionHandler) http.Handler {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/equipment/{id}", handler.FindEquipmentHandler).Methods("GET")
    router.HandleFunc("/equipment", handler.CreateEquipmentHandler).Methods("POST")
    router.HandleFunc("/equipment", handler.FindAllEquipmentHandler).Methods("GET")

    router.HandleFunc("/tours", tourHandler.CreateTour).Methods("POST")
    router.HandleFunc("/tours/{id}", tourHandler.FindByID).Methods("GET")
    router.HandleFunc("/toursByUser/{userId}", tourHandler.FindByUserId).Methods("GET")
    router.HandleFunc("/tours", tourHandler.FindAllTours).Methods("GET")

    router.HandleFunc("/equipment", handler.CreateEquipmentHandler).Methods("POST")


    router.HandleFunc("/equipmentTours", tourEqHandler.CreateTourEquipemnt).Methods("POST")
    router.HandleFunc("/equipmentTours/{idTour}/{idEquipment}", tourEqHandler.DeleteTourEquipment).Methods("DELETE")
    router.HandleFunc("/equipmentTours/{tourID}", tourEqHandler.GetTourEquipment).Methods("GET")


    router.HandleFunc("/reviews", reviewHandler.Create).Methods("POST")
    router.HandleFunc("/reviews/{id}", reviewHandler.Delete).Methods("DELETE")
    router.HandleFunc("/reviews", reviewHandler.FindAll).Methods("GET")

    router.HandleFunc("/tourPoints", tourPointHandler.CreateTourPoint).Methods("POST")
    router.HandleFunc("/tourPoints/{id}", tourPointHandler.FindByID).Methods("GET")
    router.HandleFunc("/tourPoints", tourPointHandler.FindAllTourPoints).Methods("GET")
    router.HandleFunc("/tourPointByTour/{tourId}", tourPointHandler.FindByTourId).Methods("GET")

    router.HandleFunc("/competitions", competitionHandler.CreateCompetition).Methods("POST")
    router.HandleFunc("/competitions/{id}", competitionHandler.FindByID).Methods("GET")
    router.HandleFunc("/competitions", competitionHandler.FindAllCompetitions).Methods("GET")

    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

    return router
}
