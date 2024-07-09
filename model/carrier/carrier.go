package carrier

import (
	"carrier/database"
	"errors"
	"gorm.io/gorm"
)

type Carrier struct {
	gorm.Model
	X      float64 `gorm:"not null" json:"x"`
	Y      float64 `gorm:"not null" json:"y"`
	IsBusy bool    `gorm:"not null" json:"is_busy"`
}

func (c *Carrier) Create(carrier *Carrier) error {
	db := database.DB
	if err := db.Create(&carrier); err != nil {
		return err.Error
	}
	return nil
}

func (c *Carrier) GetList() ([]Carrier, error) {
	var carriers []Carrier
	db := database.DB
	if err := db.Find(&carriers); err != nil {
		return nil, err.Error
	}
	return carriers, nil
}

func (c *Carrier) GetById(id int) (*Carrier, error) {
	var carrier Carrier
	db := database.DB
	if err := db.First(&carrier, id).Error; err != nil {
		return nil, err
	}
	return &carrier, nil
}

func (c *Carrier) Update(id int, carrier *Carrier) error {
	if carrierRow, err := carrier.GetById(id); err != nil {
		return err
	} else if carrierRow.ID == 0 {
		return errors.New("The carrier not found")
	}
	db := database.DB
	if err := db.Updates(&carrier).Error; err != nil {
		return err
	}
	return nil
}

func (c *Carrier) Delete(id int) error {
	if carrier, err := c.GetById(id); err != nil {
		return err
	} else if carrier.ID == 0 {
		return errors.New("The carrier not found")
	} else {
		db := database.DB
		if err := db.Delete(&carrier).Error; err != nil {
			return err
		}
	}
	return nil
}

func (c *Carrier) GetFreeCarriers() ([]Carrier, error) {
	db := database.DB
	carriers := []Carrier{}
	if result := db.Where("is_busy = ?", false).Find(&carriers); result.Error != nil {
		return nil, result.Error
	}
	return carriers, nil
}
