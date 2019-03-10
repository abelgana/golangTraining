package builder

type BuilProcess interface {
	SetWeels() BuilProcess
	SetSeats() BuilProcess
	SetStructure() BuilProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int32
	Seats     int32
	Structure string
}

type ManufactoringDirector struct {
	builder BuilProcess
}

func (f *ManufactoringDirector) Construct() {
	f.builder.SetWeels().SetSeats().SetStructure()
}

func (f *ManufactoringDirector) SetBuilder(b BuilProcess) {
	f.builder = b
}
