package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter := GetInstace()

	if counter == nil {
		t.Error("Expected pointer to singleton received ", counter)
	}
}

func TestGetSameInstace(t *testing.T) {
	expectedCounter := GetInstace()
	counter := GetInstace()
	if counter != expectedCounter {
		t.Error("Expected to receive ", expectedCounter, "and received ", counter)
	}
}

func TestAddOne(t *testing.T) {
	counter := GetInstace()
	currentCount := counter.AddOne()
	if currentCount != 1 {
		t.Error("Expected 1 receved ", currentCount)
	}
}

func TestAddTwo(t *testing.T) {
	counter1 := GetInstace()
	counter1.AddOne()
	counter2 := GetInstace()
	if counter2.counter != 2 {
		t.Error("Expected to receive ", counter1.counter, "and received ", counter2.counter)
	}
}
