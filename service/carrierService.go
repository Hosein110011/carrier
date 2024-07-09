package service

import (
	"carrier/model/carrier"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math"
)

func InitializeCarriers() error {
	var id int = 1
	var instance = carrier.Carrier{}
	if _, err := instance.GetById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			for i, _ := range []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1} {
				var newCarrier = carrier.Carrier{
					X:      0,
					Y:      0,
					IsBusy: false,
				}
				newCarrier.Create(&newCarrier)
				fmt.Println(i)
			}
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}

func GetAvailableCarriers() ([]carrier.Carrier, error) {
	var instance = carrier.Carrier{}
	if carriers, err := instance.GetFreeCarriers(); err != nil {
		return nil, err
	} else {
		return carriers, nil
	}
}

func UpdateCarrier(carrierID int, carrier *carrier.Carrier) error {
	if err := carrier.Update(carrierID, carrier); err != nil {
		return err
	}
	return nil
}

func DistanceMeter(xs float64, ys float64, xd float64, yd float64) float64 {
	distance := math.Sqrt(math.Pow((xd-xs), 2) + math.Pow((yd-ys), 2))
	return distance
}

func MinFinder(leastDistance float64, distance float64) float64 {
	if distance < leastDistance {
		return distance
	} else {
		return leastDistance
	}
}
