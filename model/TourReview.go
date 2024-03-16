package model

import "time"

type TourReview struct {
    Id int `json:"Id" gorm:"column:Id;primaryKey;autoIncrement"`
	Grade          float64   `json:"Grade" gorm:"column:Grade"`
	Comment        string    `json:"Comment" gorm:"column:Comment"`
	TouristId      int64     `json:"TouristId" gorm:"column:TouristId"`
	AttendanceDate time.Time `json:"AttendanceDate" gorm:"column:AttendanceDate"`
	ReviewDate     time.Time `json:"ReviewDate" gorm:"column:ReviewDate"`
	TourId         string     `json:"TourId" gorm:"column:TourId"`
}

func (TourReview) TableName() string {
	return "tours.TourReviews"
}
