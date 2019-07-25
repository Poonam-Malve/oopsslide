package main

import (
	"fmt"
	"goencapsulation/realestate"
)

func main() {
	var property realestate.PropertyOption
	//PropertyOption interface with GetName() method
	property = realestate.CreateBuilding(
		"Building 1",
		"Pune",
		"150 acr",
	)
	fmt.Println(property.GetName())

	property = realestate.CreateHouse(
		"House 1",
		"123 house",
		15, //number of rooms
	)
	fmt.Println(property.GetName())
}
