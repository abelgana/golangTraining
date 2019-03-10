package abstractFactory

import "testing"

func TestBikeFactory(t *testing.T) {
	bikeF, err := BuildFactory(BikeFactoryType)
	if err != nil {
		t.Error("Failed to create a factory")
	} else {

		_, err := bikeF.NewVehicle(SportBikeType)
		if err != nil {
			t.Error("Failed to create Bike")
		}
	}

}

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Error("Failed to create a factory")
	} else {

		_, err := carF.NewVehicle(LuxuryCarType)
		if err != nil {
			t.Error("Failed to create Car")
		}
	}
}
