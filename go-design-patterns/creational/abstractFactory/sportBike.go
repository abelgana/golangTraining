package abstractFactory

type SportBike struct {
}

func (c *SportBike) NumWheel() int64 {
	return 0
}
func (c *SportBike) NumSeats() int64 {
	return 0
}

func (c *SportBike) GetBikType() int64 {
	return 0
}
