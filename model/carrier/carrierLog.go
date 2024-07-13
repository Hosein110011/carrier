package carrier

import "gorm.io/gorm"

type CarrierUpdateLog struct {
	gorm.Model
	CarrierID int
	OldX      float64
	NewX      float64
	OldY      float64
	NewY      float64
	OldIsBusy *bool
	NewIsBusy *bool
}
