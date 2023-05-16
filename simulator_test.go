package main

import (
	"github.com/sidartaoss/maratona-fullcycle-1/simulator/entity"
	"testing"
)

func TestDestinationToJson(t *testing.T) {
	// arrange
	order := entity.Order{
		Uuid:        "a10e4124-6f1f-4b37-8188-d10842f2ed29",
		Destination: "1",
	}
	lat := "28.36588"
	lng := "-81.2654"

	// act
	actual := destinationToJson(order, lat, lng)

	// assert
	if actual == nil {
		t.Error("should return json")
	}
}

func TestStringInSlice(t *testing.T) {
	// arrange
	expected := "a10e4124-6f1f-4b37-8188-d10842f2ed29"
	s := []string{"a10e4124-6f1f-4b37-8188-d10842f2ed29", "762d0004-b421-4e9d-be8a-4359272b4380"}

	// act
	actual := stringInSlice(expected, s)

	// assert
	if actual == false {
		t.Error("should have been found")
	}
}
