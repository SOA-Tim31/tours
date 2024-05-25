package main

import (
	"context"
	"database-example/handler"
	"database-example/repo"
	"database-example/routing"
	"database-example/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Tour struct {
	Name        string
	Description string
	Status      int
	Price       float32
	UserId      int
}

func main() {
	uri := "mongodb+srv://ninakatarina:12345678NN@clusternn.dc6iczv.mongodb.net/"
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	fmt.Println("Connected!")

	collection := client.Database("MONGODB").Collection("tours")
	fmt.Println(collection.Name())

	tour1 := Tour{"Tura1", "lepa", 1, 400.5, 1}
	tour2 := Tour{"Tura2", "lepa ju", 1, 450.5, 1}
	tours := []interface{}{tour1, tour2}
	insertManyResult, err := collection.InsertMany(context.TODO(), tours)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	tourRepo, err := repo.NewTourRepository(ctx)
	if err != nil {
		log.Fatal("Error creating tour repository:", err)
	}

	tourService := &service.TourService{TourRepository: tourRepo}
	tourHandler := &handler.TourHandler{TourService: tourService}

	router := routing.SetupRoutes(tourHandler)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
