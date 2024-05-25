package repo

import (
	"context"
	"database-example/model"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type TourRepository struct {
	cli *mongo.Client
}

func NewTourRepository(ctx context.Context) (*TourRepository, error) {
	dburi := "mongodb+srv://ninakatarina:12345678NN@clusternn.dc6iczv.mongodb.net/"

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &TourRepository{
		cli: client,
	}, nil
}

func (br *TourRepository) Disconnect(ctx context.Context) error {
	err := br.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (br *TourRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := br.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}

	databases, err := br.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(databases)
}

func (br *TourRepository) getCollection() *mongo.Collection {
	blogDatabase := br.cli.Database("tours")
	blogsCollection := blogDatabase.Collection("tours")
	return blogsCollection
}

func (repo *TourRepository) Create(tour *model.Tour) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	toursCollection := repo.getCollection()

	_, err := toursCollection.InsertOne(ctx, &tour)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

	// dbResult := repo.DatabaseConnection.Create(tour)
	// if dbResult.Error != nil {
	// 	return dbResult.Error
	// }
	// println("Rows affected: ", dbResult.RowsAffected)
	// return nil
}

func (repo *TourRepository) FindById(id int) (*model.Tour, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	toursCollection := repo.getCollection()

	var tour model.Tour
	err := toursCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&tour)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &tour, nil

	// tour := model.Tour{}

	// dbResult := repo.DatabaseConnection.Preload("TourPoints").First(&tour, id)
	// if dbResult != nil {
	// 	return tour, dbResult.Error
	// }
	// return tour, nil
}

func (repo *TourRepository) FindByUserId(userId int) ([]model.Tour, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	toursCollection := repo.getCollection()

	var tours []model.Tour
	toursCursor, err := toursCollection.Find(ctx, bson.M{"userId": userId})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err = toursCursor.All(ctx, &tours); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return tours, nil

	// var tours []model.Tour
	// dbResult := repo.DatabaseConnection.Preload("TourPoints").Find(&tours, `"UserId" = ?`, userId)
	// if dbResult.Error != nil {
	// 	return nil, dbResult.Error
	// }

	// return tours, nil
}

func (repo *TourRepository) FindAll() ([]model.Tour, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	toursCollection := repo.getCollection()

	var tours []model.Tour

	toursCursor, err := toursCollection.Find(ctx, toursCollection)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err = toursCursor.All(ctx, &tours); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return tours, nil

	// dbResult := repo.DatabaseConnection.Preload("TourPoints").Find(&tours)
	// if dbResult.Error != nil {
	// 	return nil, dbResult.Error
	// }
	// return tours, nil

}
