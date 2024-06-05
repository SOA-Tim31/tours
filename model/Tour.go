package model

import (
	"time"

	"github.com/jackc/pgtype"
	"github.com/lib/pq"
)
	




type DifficultyLevel int

const (
	Easy DifficultyLevel = iota
	Moderate
	Hard
)

type Status int

const (
	Draft Status = iota
	Published
	Archived
)



type Tour struct {
	ID               int          `json:"id" gorm:"primaryKey;autoIncrement;column:Id"`
	Name              string          `json:"name" gorm:"not null;type:string;column:Name"`
	Description       string          `json:"description" gorm:"type:string;column:Description"`
	DifficultyLevel   DifficultyLevel      `json:"difficultyLevel" gorm:"not null;type:int;column:DifficultyLevel"`
	Status            int         `json:"status" gorm:"not null;type:int;column:Status"`
	Price             float64         `json:"price" gorm:"not null;type:float;column:Price"`
	UserId            int             `json:"userId" gorm:"not null;type:int;column:UserId"`
	PublishedDateTime *time.Time       `json:"publishedDateTime" gorm:"column:PublishedDateTime"`
	ArchivedDateTime  *time.Time       `json:"archivedDateTime" gorm:"column:ArchivedDateTime"`
	Tags              pq.StringArray        `json:"tags" gorm:"type:text[];column:Tags"`
	TourPoints        []TourPoint         `json:"tourPoints,omitempty"`
	TourReview        []TourReview         `json:"tourReview,omitempty"`
	TourCharacteristics pgtype.JSONB  `json:"tourCharacteristics" gorm:"type:jsonb;column:TourCharacteristics;default:'[]'::jsonb"`
	
}


type TourDTO struct {
	ID               int             `json:"id"`
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	DifficultyLevel   string      `json:"difficultyLevel"`
	Status            string       `json:"status"`
	Price             float64         `json:"price"`
	UserId            int             `json:"userId"`
	PublishedDateTime *time.Time       `json:"publishedDateTime"`
	ArchivedDateTime  *time.Time       `json:"archivedDateTime"`
	Tags              []string        `json:"tags"`
	TourPoints        []string         `json:"tourPoints"`
	TourReview        []string          `json:"tourReview"`
	TourCharacteristics pgtype.JSONB  `json:"tourCharacteristics"`
	
}


func (Tour) TableName() string {
	return "tours.Tours"
}