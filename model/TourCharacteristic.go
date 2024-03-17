package model

import "time"

type TourCharacteristic struct {
	Duration time.Time `json:"Duration"`
	Distance float64     `json:"Distance"`
	TransportType string  `json:"TransportType`

}