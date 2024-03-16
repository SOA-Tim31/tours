package model



type TourEquipment struct {
    Id          int `json:"Id" gorm:"column:Id;primaryKey";autoIncrement`
    TourId      int    `json:"TourId" gorm:"column:TourId"`
    EquipmentId int    `json:"EquipmentId" gorm:"column:EquipmentId"`
    IsSelected  bool `json:"IsSelected" gorm:"column:IsSelected;default:false"`
}




func (TourEquipment) TableName() string {
	return "tours.TourEquipments"
}
