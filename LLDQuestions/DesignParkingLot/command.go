package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HandleCommand(command string) {
	parts := strings.Split(command, " ")
	switch parts[0] {
	case "create_parking_lot":
		if len(parts) == 4 {
			noOfFloors := parseInt(parts[2])
			slotsPerFloor := parseInt(parts[3])
			parkingLot = NewParkingLot(parts[1], noOfFloors, slotsPerFloor)
			fmt.Printf("Created parking lot with %d floors and %d slots per floor\n", noOfFloors, slotsPerFloor)
		}
	case "park_vehicle":
		if len(parts) == 4 {
			vehicle := Vehicle{Type: parts[1], Registration: parts[2], Colour: parts[3]}
			ticketID := parkingLot.ParkVehicle(vehicle)
			fmt.Println("Parked vehicle. Ticket ID:", ticketID)
		}
	case "unpark_vehicle":
		if len(parts) == 2 {
			message := parkingLot.UnparkVehicle(parts[1])
			fmt.Println(message)
		}
	case "display":
		if parkingLot == nil {
			fmt.Println("Parking lot has not been created yet.")
			return
		}
		if len(parts) == 3 {
			displayType := parts[1]
			vehicleType := parts[2]
			switch displayType {
			case "free_count":
				parkingLot.DisplayFreeCount(vehicleType)
			case "free_slots":
				parkingLot.DisplayFreeSlots(vehicleType)
			case "occupied_slots":
				parkingLot.DisplayOccupiedSlots(vehicleType)
			}
		}
	}
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error parsing integer:", err)
	}
	return n
}
