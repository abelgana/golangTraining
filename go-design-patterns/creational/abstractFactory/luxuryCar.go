package abstractFactory

type LuxuryCar struct {
}

func (c *LuxuryCar) NumWheel() int64 {
	return 0
}
func (c *LuxuryCar) NumSeats() int64 {
	return 0
}

func (c *LuxuryCar) GetNumDoors() int64 {
	return 0
}
