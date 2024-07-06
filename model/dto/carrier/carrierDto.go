package carrier

import "carrier/model/dto/destination"

type CarrierDto struct {
	ID          int                         `json:"id"`
	X           float64                     `json:"x"`
	Y           float64                     `json:"y"`
	Destination *destination.DestinationDto `json:"destination"`
	Distance    float64                     `json:"distance"`
}
