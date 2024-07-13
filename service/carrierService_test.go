package service

import (
	"carrier/database"
	"carrier/model/carrier"
	"carrier/model/destination"
	"carrier/repository"
	"testing"
	"testing/quick"
)

func TestGetAvailableCarriers(t *testing.T) {

	database.DB = database.ConnectDB()

	defer database.DB.Close()

	f := func() bool {
		carriers, err := repository.GetFreeCarriers()
		if err != nil {
			t.Error(err)
			return false
		}
		if carriers == nil {
			t.Errorf("no available carriers found in db.")
			return false
		}
		for _, carrierObj := range carriers {
			if *carrierObj.IsBusy == true {
				t.Errorf("carrier with id %v is busy", carrierObj.ID)
				return false
			}
		}
		return true
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUpdateCarrier(t *testing.T) {
	isBusy := true

	var testCarrier = carrier.Carrier{
		ID:     1,
		X:      22,
		Y:      22,
		IsBusy: &isBusy,
	}

	database.DB = database.ConnectDB()

	defer database.DB.Close()

	defer func() {
		isBusy = false
		testCarrier.IsBusy = &isBusy
		UpdateCarrier(testCarrier.ID, testCarrier)
	}()
	f := func() bool {

		err := UpdateCarrier(testCarrier.ID, testCarrier)

		if err != nil {
			t.Error(err)
			return false
		}

		return true
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}

}

func TestInitializeCarriers(t *testing.T) {
	f := func() bool {

		err := InitializeCarriers()

		if err != nil {
			t.Error(err)
			return false
		}

		return true
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestGetDistance(t *testing.T) {
	f := func() bool {
		expected := 5.0
		result := GetDistance(2.0, 4.0, 5.0, 8.0)
		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
			return false
		}
		return true
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestFindMinDistance(t *testing.T) {

	distances := []float64{2.0, 4.0, 3.0, 2.0, 1.0}
	var minDistance float64 = 3.0

	f := func() bool {
		for _, distance := range distances {
			minDistance = FindMinDistance(minDistance, distance)
		}
		for _, distance := range distances {
			if distance < minDistance {
				t.Errorf("minimum distance (%v) is grater then %v :/", minDistance, distance)
				return false
			}
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

//func TestGetAverageCounter(t *testing.T) {
//	testDestination := destination.Destination{
//		Xd: 24,
//		Yd: 41,
//	}
//
//	testCarrier := carrier.Carrier{
//		ID: 1,
//		X:  22,
//		Y:  52,
//	}
//
//	distance := service.GetDistance(testCarrier.X, testCarrier.Y, testDestination.Xd, testDestination.Yd)
//
//	f := func() bool {
//		counterX := service.GetAverageCounter(testCarrier.X, testDestination.Xd, distance)
//		counterY := service.GetAverageCounter(testCarrier.Y, testDestination.Yd, distance)
//
//		fmt.Println("x counter : ", counterX, "y counter : ", counterY)
//
//		for i := 0; i <= int(math.Ceil(distance)); i++ {
//			testCarrier.X += counterX
//			testCarrier.Y += counterY
//		}
//		if testDestination.Xd != service.RoundFloat(testCarrier.X, 0) || testDestination.Yd != service.RoundFloat(testCarrier.Y, 0) {
//			t.Errorf("got %v, %v, expected %v, %v", testCarrier.X, testCarrier.Y, testDestination.Xd, testDestination.Yd)
//			return false
//		}
//		return true
//	}
//	if err := quick.Check(f, nil); err != nil {
//		t.Error(err)
//	}
//}

func TestCarrierTaskMaker(t *testing.T) {

	database.DB = database.ConnectDB()

	defer database.DB.Close()

	isBusy := true
	var testCarrier = carrier.Carrier{
		ID:     1,
		X:      22,
		Y:      22,
		IsBusy: &isBusy,
	}

	testDestination := destination.Destination{
		Xd: 24,
		Yd: 41,
	}

	distance := GetDistance(testCarrier.X, testCarrier.Y, testDestination.Xd, testDestination.Yd)
	counterX := GetAverageCounter(testCarrier.X, testDestination.Xd, distance)
	counterY := GetAverageCounter(testCarrier.Y, testDestination.Yd, distance)

	f := func() bool {
		err := CarrierTaskMaker(counterX, counterY, distance, testCarrier)
		if err != nil {
			t.Error(err)
			return false
		}
		return true
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
