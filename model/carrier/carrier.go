package carrier

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Carrier struct {
	gorm.Model
	X      float64 `gorm:"not null" json:"x"`
	Y      float64 `gorm:"not null" json:"y"`
	IsBusy bool    `gorm:"not null" json:"is_busy"`
}

func (c *Carrier) Create(carrier *Carrier, db *gorm.DB) error {
	if err := db.Create(&carrier); err != nil {
		return err.Error
	}
	return nil
}

func (c *Carrier) GetList(db *gorm.DB) ([]Carrier, error) {
	var carriers []Carrier
	//db := database.DB
	if err := db.Find(&carriers); err != nil {
		return nil, err.Error
	}
	return carriers, nil
}

func (c *Carrier) GetById(id int, db *gorm.DB) (*Carrier, error) {
	//var carrier Carrier
	//db := database.DB
	fmt.Println(db, "errrrrrrrrrr")
	if err := db.First(&c, id).Error; err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Carrier) Update(id int, carrier *Carrier, db *gorm.DB) error {
	if carrierRow, err := carrier.GetById(id, db); err != nil {
		return err
	} else if carrierRow.ID == 0 {
		return errors.New("The carrier not found")
	}
	//db := database.DB
	if err := db.Updates(&carrier).Error; err != nil {
		return err
	}
	return nil
}

func (c *Carrier) Delete(id int, db *gorm.DB) error {
	if carrier, err := c.GetById(id, db); err != nil {
		return err
	} else if carrier.ID == 0 {
		return errors.New("The carrier not found")
	} else {
		if err := db.Delete(&carrier).Error; err != nil {
			return err
		}
	}
	return nil
}

func (c *Carrier) GetFreeCarriers(db *gorm.DB) ([]Carrier, error) {
	//db := database.DB
	carriers := []Carrier{}
	if result := db.Where("is_busy = ?", false).Find(&carriers); result.Error != nil {
		return nil, result.Error
	}
	return carriers, nil
}
