package model



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
	Id             int             `json:"Id" gorm:"primaryKey;autoIncrement;column:Id"`
	Name           string          `json:"Name" gorm:"not null;type:string;column:Name"`
	Description    string          `json:"Description" gorm:"type:string;column:Description"`
	DifficultyLevel DifficultyLevel `json:"DifficultyLevel" gorm:"not null;type:int;column:DifficultyLevel"`
	Status         Status          `json:"Status" gorm:"not null;type:int;column:Status"`
	Price          float64         `json:"Price" gorm:"not null;type:float;column:Price"`
}



func (Tour) TableName() string {
    return "tours.Tours"
}