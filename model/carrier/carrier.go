package carrier

import (
	"gorm.io/gorm"
)

type Carrier struct {
	gorm.Model
	ID     int     `gorm:"not null" json:"id"`
	X      float64 `gorm:"not null" json:"x"`
	Y      float64 `gorm:"not null" json:"y"`
	IsBusy bool    `gorm:"not null" json:"is_busy"`
}
