package routing

import (
    "database-example/handler"
    "github.com/gorilla/mux"
    "net/http"
)

func SetupRoutes(handler *handler.EquipmentHandler) http.Handler {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/equipment/{id}", handler.FindEquipmentHandler).Methods("GET")
    router.HandleFunc("/equipment", handler.CreateEquipmentHandler).Methods("POST")
    router.HandleFunc("/equipment", handler.FindAllEquipmentHandler).Methods("GET")

    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

    return router
}
