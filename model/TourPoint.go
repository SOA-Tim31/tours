package model

type TourPoint struct {
	ID          int     `json:"id" gorm:"primaryKey;autoIncrement;column:Id"`
	Name        string  `json:"name" gorm:"type:string;column:Name"`
	TourID      int     `json:"tourId" gorm:"type:int;column:TourId"`
	Description string  `json:"description" gorm:"type:string;column:Description"`
	Latitude    float64 `json:"latitude" gorm:"type:float;column:Latitude"`
	Longitude   float64 `json:"longitude" gorm:"type:float;column:Longitude"`
	ImageUrl    string  `json:"imageUrl" gorm:"type:string;column:ImageUrl"`
	Secret      string  `json:"secret" gorm:"type:string;column:Secret"`
}

func (TourPoint) TableName() string {
	return "tours.TourPoint"
}