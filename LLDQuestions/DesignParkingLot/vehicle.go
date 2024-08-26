package main

type Vehicle struct {
	Type         string
	Registration string
	Colour       string
}

type Ticket struct {
	ID             string
	FloorNumber    int
	SlotNumber     int
	VehicleDetails Vehicle
}
