package routing

import (
	"database-example/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(tourHandler *handler.TourHandler) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tours", tourHandler.CreateTour).Methods("POST")
	router.HandleFunc("/tours/{id}", tourHandler.FindByID).Methods("GET")
	router.HandleFunc("/toursByUser/{userId}", tourHandler.FindByUserId).Methods("GET")
	router.HandleFunc("/tours", tourHandler.FindAllTours).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	return router
}
