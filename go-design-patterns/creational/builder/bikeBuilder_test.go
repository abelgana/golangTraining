package builder

import "testing"

func TestBikeBuilder(t *testing.T) {
	manuFactoringComplex := ManufactoringDirector{}

	bikeBuilder := &BikeBuilder{}
	manuFactoringComplex.SetBuilder(bikeBuilder)
	manuFactoringComplex.Construct()
	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Error("Expect 2 wheels received ", bike.Wheels)
	}

	if bike.Structure != "Bike" {
		t.Error("Expected bike structure received ", bike.Structure)
	}

	if bike.Seats != 2 {
		t.Error("Expected 2 seats received ", bike.Seats)
	}
}
