package service

import (
	"carrier/model/carrier"
	"carrier/repository"
	"errors"
	"fmt"
	"time"

	//"github.com/madflojo/tasks"
	"gorm.io/gorm"
	"math"
	//"time"
)

func InitializeCarriers() error {
	var id int = 1
	if _, err := repository.GetById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			for i := 1; i <= 10; i++ {
				var newCarrier = carrier.Carrier{
					ID: i,
					X:  0,
					Y:  0,
				}
				if err := repository.Create(&newCarrier); err != nil {
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
	if carriers, err := repository.GetFreeCarriers(); err != nil {
		return nil, err
	} else {
		return carriers, nil
	}
}

func UpdateCarrier(carrierID int, carrier carrier.Carrier) error {
	//fmt.Println("UpdateCarrier : ", carrier)
	if err := repository.Update(carrierID, &carrier); err != nil {
		return err
	}
	return nil
}

func GetDistance(xs float64, ys float64, xd float64, yd float64) float64 {
	distance := math.Sqrt(math.Pow(xd-xs, 2) + math.Pow(yd-ys, 2))
	return distance
}

func FindMinDistance(minDistance float64, distance float64) float64 {
	if distance < minDistance {
		return distance
	} else {
		return minDistance
	}
}

//func GetTotalTime(distance float64) float64 {
//	T := constants.Velocity / distance
//	return T
//}

func GetAverageCounter(source, destination, distance float64) float64 {
	return (destination - source) / distance
}

func taskProcessor(selectedCarrier carrier.Carrier, counterX, counterY, distance float64) {
	println("new task started...")

	go func() {

		defer func() {
			*selectedCarrier.IsBusy = false

			if err := UpdateCarrier(selectedCarrier.ID, selectedCarrier); err != nil {
				return
			}

		}()
		for i := 0; i < int(math.Ceil(distance)); i++ {
			selectedCarrier.X += counterX
			selectedCarrier.Y += counterY
			err := UpdateCarrier(selectedCarrier.ID, selectedCarrier)
			if err != nil {
				return
			}
			fmt.Println("new X = ", selectedCarrier.X)
			fmt.Println("new Y = ", selectedCarrier.Y)

			time.Sleep(1 * time.Second)
		}
	}()
}

func CarrierTaskMaker(
	counterX, counterY, distance float64,
	selectedCarrier carrier.Carrier) error {
	*selectedCarrier.IsBusy = true
	if err := UpdateCarrier(selectedCarrier.ID, selectedCarrier); err != nil {
		return err
	}
	taskProcessor(selectedCarrier, counterX, counterY, distance)

	return nil
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
