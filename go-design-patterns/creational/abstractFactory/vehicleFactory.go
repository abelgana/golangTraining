package abstractFactory

import "errors"

type VehicleFactory interface {
	NewVehicle(v int64) (Vehicle, error)
}

const (
	CarFactoryType = iota
	BikeFactoryType
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case BikeFactoryType:
		return new(BikeFactory), nil
	default:
		return nil, errors.New("Unknown type")
	}
}
