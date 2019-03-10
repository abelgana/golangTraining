package builder

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWeels() BuilProcess {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() BuilProcess {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() BuilProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}
