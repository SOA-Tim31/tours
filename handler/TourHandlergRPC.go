package handler

import (
	"context"
	"database-example/model"
	tours "database-example/proto"
	"database-example/service"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
	"github.com/lib/pq"
)


type TourHandlergRPC struct {
	TourService *service.TourService
	tours.UnimplementedTourServiceServer
}

func NewTourHandlergRPC(t *service.TourService) *TourHandlergRPC {
	return &TourHandlergRPC{
		TourService: t,
	}
}


func(th *TourHandlergRPC) CreateTour(ctx context.Context, req *tours.Tour)(*tours.TourResponse, error){
	layout := "2006-01-02 15:04:05"
	 publishTime, err := time.Parse(layout, req.PublishedDateTime)
	if err != nil {
		return nil, err
	}

	archiveTime, err := time.Parse(layout, req.PublishedDateTime)
	if err != nil {
		return nil, err
	}

	tour := model.Tour{
		ID: int(req.Id),
        Name: req.Name,
		Description: req.Description,
		DifficultyLevel : model.DifficultyLevel(req.Level),
		Status: int(req.Status),
		Price: float64(req.Price),
        UserId: int(req.UserId),
        PublishedDateTime: &publishTime,
        ArchivedDateTime: &archiveTime,
		Tags: pq.StringArray{},
		TourPoints: []model.TourPoint{},
		TourReview: []model.TourReview{},
		TourCharacteristics: pgtype.JSONB{},

	}

	err = th.TourService.Create(&tour)
	if err != nil {
		fmt.Printf("Database exception %s\n", err)
		return nil, err
	}

	return &tours.TourResponse{
		Message: "SUCCESS",
	}, nil
}


func (th *TourHandlergRPC) GetAllTours(ctx context.Context, req *tours.GetAllRequest)(*tours.GetAllResponse, error){
	allTours, err := th.TourService.FindAll()

	if err != nil {
		return nil, fmt.Errorf("error while finding all tours: %v", err)
	}

	var toursDto []*tours.Tour

	for _, tour := range allTours {
		toursDto = append(toursDto, &tours.Tour{
			Id: uint64(tour.ID),
			Name: tour.Name,
			Description: tour.Description,
			Level: tours.DifficultyLevel(tour.DifficultyLevel),
			Status: uint64(tour.Status),
			Price: uint64(tour.Price),
			PublishedDateTime: tour.PublishedDateTime.Format("2006-01-02 15:04:05"),
			ArchivedDateTime: tour.ArchivedDateTime.Format("2006-01-02 15:04:05"),
			UserId: uint64(tour.UserId),
		})
	}

	return &tours.GetAllResponse{
		Tours: toursDto,
	}, nil
}
