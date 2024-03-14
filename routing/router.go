package routing

import (
	"database-example/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(handler *handler.EquipmentHandler, tourHandler *handler.TourHandler,tourEqHandler *handler.TourEquipmentHandler) http.Handler {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/equipment/{id}", handler.FindEquipmentHandler).Methods("GET")
    router.HandleFunc("/equipment", handler.CreateEquipmentHandler).Methods("POST")
    router.HandleFunc("/equipment", handler.FindAllEquipmentHandler).Methods("GET")

    router.HandleFunc("/tours", tourHandler.CreateTour).Methods("POST")
    router.HandleFunc("/tours/{id}", tourHandler.FindByID).Methods("GET")
    router.HandleFunc("/tours", tourHandler.FindAllTours).Methods("GET")

    router.HandleFunc("/equipment", handler.CreateEquipmentHandler).Methods("POST")


    router.HandleFunc("/equipmentTours", tourEqHandler.CreateTourEquipemnt).Methods("POST")

    router.HandleFunc("/equipmentTours/{idTour}/{idEquipment}", tourEqHandler.DeleteTourEquipment).Methods("DELETE")

    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

    return router
}
