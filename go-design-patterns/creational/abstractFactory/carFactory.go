package abstractFactory

import "errors"

const (
	LuxuryCarType = iota
	FamilyCarType
)

type CarFactory struct {
}

func (cf *CarFactory) NewVehicle(vt int64) (Vehicle, error) {
	switch vt {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New("Unknown type of car")
	}
}
