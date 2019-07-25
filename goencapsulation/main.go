package main

import (
	"fmt"

	"goslides/goencapsulation/realstate"
)

func main() {
	var property realstate.PropertyOption

	//PropertyOption interface with GetName() method
	property = realstate.CreateBuilding(
		"Building 1",
		"Pune",
		"150 acr",
	)

	fmt.Println(property.GetName())

	property = realstate.CreateHouse(
		"House 1",
		"123 house",
		15, //number of rooms
	)

	fmt.Println(property.GetName())
}
