package model


type Equipment struct {
    Id           int `json:"Id" gorm:"primaryKey;column:Id";autoIncrement`
    Name         string    `json:"Name" gorm:"not null;type:string;column:Name"`
    Description  string    `json:"Description" gorm:"not null;type:string;column:Description"`
}


func (Equipment) TableName() string {
    return "tours.Equipment"
}