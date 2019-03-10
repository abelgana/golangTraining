package abstractFactory

type CruiseBike struct {
}

func (c *CruiseBike) NumWheel() int64 {
	return 0
}
func (c *CruiseBike) NumSeats() int64 {
	return 0
}

func (c *CruiseBike) GetBikType() int64 {
	return 0
}
