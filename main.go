package main

import (
	"database-example/handler"
	"database-example/migration"
	"database-example/repo"
	"database-example/routing"
	"database-example/service"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "postgres://postgres:super@localhost:5432/explorer-v1?sslmode=disable&search_path=tours"

    database, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
        return nil
    }

    if err := database.Exec("CREATE SCHEMA IF NOT EXISTS tours").Error; err != nil {
        log.Fatalf("Failed to create schema: %v", err)
        return nil
    }

    if err := migration.AutoMigrate(database); err != nil {
        log.Fatalf("Failed to perform auto migration: %v", err)
        return nil
    }
    return database
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	equipmentRepo := &repo.EquipmentRepository{DatabaseConnection: database}
    equipmentService := &service.EquipmentService{EquipmentRepository: equipmentRepo}
    equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}

	tourRepository := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepository: tourRepository}
	tourHandler := &handler.TourHandler{TourService: tourService}

	tourEqRepository := &repo.TourEquipmentRepository{DatabaseConnection: database}
	tourEqService := &service.TourEquipmentService{TourEquipmentRepository: tourEqRepository}
	tourEqHandler := &handler.TourEquipmentHandler{TourEquipmentService: tourEqService}


    router := routing.SetupRoutes(equipmentHandler, tourHandler,tourEqHandler)

    log.Println("Server starting...")
    log.Fatal(http.ListenAndServe(":8000", router))
}