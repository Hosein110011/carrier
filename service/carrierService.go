package service

import (
	"carrier/constants"
	"carrier/database"
	"carrier/model/carrier"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math"
)

func InitializeCarriers() error {
	db := database.DB
	var id int = 1
	var instance = carrier.Carrier{}
	if _, err := instance.GetById(id, db); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			for i, _ := range []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1} {
				var newCarrier = carrier.Carrier{
					X:      0,
					Y:      0,
					IsBusy: false,
				}
				err := newCarrier.Create(&newCarrier, db)
				if err != nil {
					return err
				}
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
	db := database.DB
	var instance = carrier.Carrier{}
	if carriers, err := instance.GetFreeCarriers(db); err != nil {
		return nil, err
	} else {
		return carriers, nil
	}
}

func UpdateCarrier(carrierID int, carrier *carrier.Carrier) error {
	db := database.DB
	if err := carrier.Update(carrierID, carrier, db); err != nil {
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

func AverageCounterFinder(source, destination, distance float64) float64 {
	T := (constants.Velocity / distance)
	return (destination - source) / T
}

func CarrierTaskMaker(counterX, counterY float64) error {
	//// Start the Scheduler
	//scheduler := tasks.New()
	//defer scheduler.Stop()
	//
	//// Add a task
	//id, err := scheduler.Add(&tasks.Task{
	//	Interval: 60 * time.Second,
	//	TaskFunc: func() error {
	//
	//	},
	//})
	//if err != nil {
	//	// Do Stuff
	//}
	var err error
	return err
}
