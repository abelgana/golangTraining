package abstractFactory

type FamilyCar struct {
}

func (c *FamilyCar) NumWheel() int64 {
	return 0
}
func (c *FamilyCar) NumSeats() int64 {
	return 0
}

func (c *FamilyCar) GetNumDoors() int64 {
	return 0
}
