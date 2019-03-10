package builder

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWeels() BuilProcess {
	b.v.Wheels = 2
	return b
}
func (b *BikeBuilder) SetSeats() BuilProcess {
	b.v.Seats = 2
	return b
}
func (b *BikeBuilder) SetStructure() BuilProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}
