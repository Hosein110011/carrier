package service

import (
	"carrier/constants"
	"carrier/model/carrier"
	"carrier/repository"
	"errors"
	"fmt"
	"sync"
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
			for i, _ := range []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1} {
				var newCarrier = carrier.Carrier{
					X:      0,
					Y:      0,
					IsBusy: false,
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

func UpdateCarrier(carrierID int, carrier *carrier.Carrier) error {
	if err := repository.Update(carrierID, carrier); err != nil {
		return err
	}
	return nil
}

func GetDistance(xs float64, ys float64, xd float64, yd float64) float64 {
	distance := math.Sqrt(math.Pow(xd-xs, 2) + math.Pow(yd-ys, 2))
	return distance
}

func FindMinDistance(leastDistance float64, distance float64) float64 {
	if distance < leastDistance {
		return distance
	} else {
		return leastDistance
	}
}

func GetTotalTime(distance float64) float64 {
	T := constants.Velocity / distance
	return T
}

func GetAverageCounter(source, destination, distance, T float64) float64 {
	return (destination - source) / T
}

func taskProcessor(selectedCarrier carrier.Carrier, counterX, counterY float64, totalTime float64) {
	println("new task proccess started...")
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Done()
	go func() {
		fmt.Println("Goroutine")
		fmt.Println("totalTime == ", totalTime)
		for i := 0; i < int(math.Ceil(totalTime)); i++ {
			i += 1
			fmt.Println("i == ", i)
			selectedCarrier.X += counterX
			selectedCarrier.Y += counterY
			selectedCarrier.IsBusy = true
			err := UpdateCarrier(selectedCarrier.ID, &selectedCarrier)
			if err != nil {
				return
			}
			println("new X = ", selectedCarrier.X)
			println("new Y = ", selectedCarrier.Y)
			fmt.Println(i, "seconds passed...")
			time.Sleep(1 * time.Second)
		}
		selectedCarrier.IsBusy = false
		err := UpdateCarrier(selectedCarrier.ID, &selectedCarrier)
		if err != nil {
			return
		}
		println("Updated to false>>>>")
	}()
}

func CarrierTaskMaker(counterX,
	counterY float64,
	selectedCarrier carrier.Carrier,
	totalTime float64) error {

	taskProcessor(selectedCarrier, counterX, counterY, totalTime)
	selectedCarrier.IsBusy = false
	if err := UpdateCarrier(selectedCarrier.ID, &selectedCarrier); err != nil {
		return err
	}
	return nil
}
