package repository

import (
	"carrier/database"
	"carrier/model/carrier"
	"errors"
)

func Create(carrier *carrier.Carrier) error {
	db := database.DB
	if err := db.Create(&carrier); err != nil {
		return err.Error
	}
	return nil
}

func GetList() ([]carrier.Carrier, error) {
	var carriers []carrier.Carrier
	db := database.DB
	if err := db.Find(&carriers); err != nil {
		return nil, err.Error
	}
	return carriers, nil
}

func GetById(id int) (*carrier.Carrier, error) {
	var carrierInstance carrier.Carrier
	db := database.DB
	if err := db.First(&carrierInstance, id).Error; err != nil {
		return nil, err
	}
	return &carrierInstance, nil
}

func Update(id int, carrierInstance *carrier.Carrier) error {
	if carrierRow, err := GetById(id); err != nil {
		return err
	} else if carrierRow.ID == 0 {
		return errors.New("The carrier not found")
	}
	db := database.DB
	if err := db.Updates(&carrierInstance).Error; err != nil {
		return err
	}
	return nil
}

func Delete(id int) error {
	if carrierInstance, err := GetById(id); err != nil {
		return err
	} else if carrierInstance.ID == 0 {
		return errors.New("the carrier not found")
	} else {
		db := database.DB
		if err := db.Delete(&carrierInstance).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetFreeCarriers() ([]carrier.Carrier, error) {
	db := database.DB
	var carriers []carrier.Carrier
	if result := db.Where("is_busy = ?", false).Find(&carriers); result.Error != nil {
		return nil, result.Error
	}
	return carriers, nil
}
