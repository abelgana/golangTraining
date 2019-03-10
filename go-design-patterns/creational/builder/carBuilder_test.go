package builder

import "testing"

func TestCarBuilder(t *testing.T) {
	manuFactoringComplex := ManufactoringDirector{}

	carBuilder := &CarBuilder{}
	manuFactoringComplex.SetBuilder(carBuilder)
	manuFactoringComplex.Construct()
	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Error("Expect 4 wheels received ", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Error("Expected Car structure received ", car.Structure)
	}

	if car.Seats != 5 {
		t.Error("Expected 5 seats received ", car.Seats)
	}
}
