package model

import (
	"time"
)

type CompetitionStatus int

const (
	Open CompetitionStatus = iota
	Close
	
)

type Competition struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;column:Id"`
	TourID    int       `json:"tourId" gorm:"not null;type:int;column:TourId"`
	StartDate *time.Time `json:"startDate" gorm:"column:StartDate"`
	Duration  int        `json:"duration" gorm:"type:int;column:Duration"`
	Status  CompetitionStatus   `json:"status" gorm:"not null;type:int;column:Status"`  
	CompetitionApplies  []string         `json:"competitionApplies" gorm:"-"` 
	//tip []string je privremen dok ne napravimo CompetitionApply struct

}


type CompetitionDTO struct{
	ID        int       `json:"id"`
	TourID    int       `json:"tourId"`
	StartDate *time.Time `json:"startDate"`
	Duration  int        `json:"duration"`
	Status  string   `json:"status"`  
	CompetitionApplies  []string         `json:"competitionApplies"` 
}


func (Competition) TableName() string {
	return "tours.Competitions"
}