package main

import (
	"database-example/handler"
	"database-example/migration"
	"database-example/repo"
	"database-example/service"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	tours "database-example/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "user=postgres password=super dbname=explorer host=database port=5432 sslmode=disable search_path=tours"

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
	// equipmentRepo := &repo.EquipmentRepository{DatabaseConnection: database}
    // equipmentService := &service.EquipmentService{EquipmentRepository: equipmentRepo}
    // equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}

	tourRepository := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepository: tourRepository}
	// tourHandler := &handler.TourHandler{TourService: tourService}

	// tourPointRepository := &repo.TourPointRepository{DatabaseConnection: database}
	// tourPointService := &service.TourPointService{TourPointRepository: tourPointRepository}
	// tourPointHandler := &handler.TourPointHandler{TourPointService: tourPointService}

	// tourEqRepository := &repo.TourEquipmentRepository{DatabaseConnection: database}
	// tourEqService := &service.TourEquipmentService{TourEquipmentRepository: tourEqRepository}
	// tourEqHandler := &handler.TourEquipmentHandler{TourEquipmentService: tourEqService}

	// tourReviewRepository := &repo.TourReviewRepository{DatabaseConnection: database}
	// tourReviewService := &service.TourReviewService{TourReviewRepository: tourReviewRepository}
	// tourReviewHandler := &handler.TourReviewHandler{TourReviewService: tourReviewService}

	// tourObjectRepository := &repo.TourObjectRepository{DatabaseConnection: database}
	// tourObjectService := &service.TourObjectService{TourObjectRepository: tourObjectRepository}
	// tourObjectHandler := &handler.TourObjectHandler{TourObjectService: tourObjectService}


   

	// competitionRepository := &repo.CompetitionRepository{DatabaseConnection: database}
	// competitionService := &service.CompetitionService{CompetitionRepository: competitionRepository}
	// competitionHandler := &handler.CompetitionHandler{CompetitionService: competitionService}


    // router := routing.SetupRoutes(equipmentHandler, tourHandler,tourEqHandler,tourReviewHandler,tourPointHandler,tourObjectHandler, competitionHandler)


	tourHandlergRPC := handler.NewTourHandlergRPC(tourService)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	tours.RegisterTourServiceServer(grpcServer, tourHandlergRPC)

	//GRPC
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)
	fmt.Println("Server listening on port 80")

	//Distribute all the connections to goroutines
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	fmt.Println("Received terminate, graceful shutdown", sig)
}