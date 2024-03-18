package model

type TourObject struct {
	Id int `json:"Id" gorm:"column:Id;primaryKey;autoIncrement"`
	Name string    `json:"Name" gorm:"column:Name"`
	Description string    `json:"Description" gorm:"column:Description"`
	ImageUrl string    `json:"ImageUrl" gorm:"column:ImageUrl"`
	Category string    `json:"Category" gorm:"column:Category"`
    Latitude  float64 `json:"Latitude" gorm:"column:Latitude"`
	Longitude  float64 `json:"Longitude" gorm:"column:Longitude"`
	TourId      int    `json:"TourId" gorm:"column:TourId"` 
}

func (TourObject) TableName() string {
	return "tours.TourObject"
}