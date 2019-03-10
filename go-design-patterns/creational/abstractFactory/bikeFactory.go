package abstractFactory

import "errors"

const (
	CruiseBikeType = iota
	SportBikeType
)

type BikeFactory struct {
}

func (bf *BikeFactory) NewVehicle(vt int64) (Vehicle, error) {
	switch vt {
	case CruiseBikeType:
		return new(CruiseBike), nil
	case SportBikeType:
		return new(SportBike), nil
	default:
		return nil, errors.New("Unknown type of car")
	}
}
