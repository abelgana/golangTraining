package singleton

type singleton struct {
	counter int64
}

type Singleton interface {
	addOne() int64
}

var instance *singleton

func GetInstace() (s *singleton) {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int64 {
	s.counter++
	return s.counter
}
