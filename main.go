package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
    connectionStr := "postgres://postgres:super@localhost:5432/explorer-v1?sslmode=disable"

    database, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
        return nil
    }

    if err := database.Exec("CREATE SCHEMA IF NOT EXISTS your_schema").Error; err != nil {
        log.Fatalf("Failed to create schema: %v", err)
        return nil
    }

    if err := database.Table("your_schema.students").AutoMigrate(&model.Student{}); err != nil {
        log.Fatalf("Failed to auto migrate: %v", err)
        return nil
    }

    if err := database.Exec("INSERT INTO your_schema.students (id, name, major) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING",
        "aec7e123-233d-4a09-a289-75308ea5b7e6", "Marko Markovic", "Graficki dizajn").Error; err != nil {
        log.Fatalf("Failed to insert record: %v", err)
        return nil
    }

    return database
}



func startServer(handler *handler.StudentHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/students/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/students", handler.Create).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	repo := &repo.StudentRepository{DatabaseConnection: database}
	service := &service.StudentService{StudentRepo: repo}
	handler := &handler.StudentHandler{StudentService: service}

	startServer(handler)
}