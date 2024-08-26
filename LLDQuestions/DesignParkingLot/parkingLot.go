package main

import "fmt"

type Slot struct {
	SlotNumber int
	SlotType   string
	IsOccupied bool
	Ticket     *Ticket
}

type Floor struct {
	FloorNumber int
	Slots       []Slot
}

type ParkingLot struct {
	ID     string
	Floors []Floor
}

func NewParkingLot(id string, numberOfFloors int, slotsPerFloor int) *ParkingLot {

	floors := make([]Floor, numberOfFloors)

	for i := 0; i < numberOfFloors; i++ {
		floors[i] = Floor{
			FloorNumber: i,
			Slots:       createSlots(slotsPerFloor),
		}

	}

	return &ParkingLot{
		ID:     id,
		Floors: floors,
	}
}

func createSlots(slotsPerFloor int) []Slot {

	slots := make([]Slot, slotsPerFloor)

	for i := 0; i < slotsPerFloor; i++ {
		slotType := getSlotType(i)
		slots[i] = Slot{
			SlotNumber: i,
			SlotType:   slotType,
			IsOccupied: false,
		}
	}

	return slots
}

func getSlotType(slotNumber int) string {
	if slotNumber == 0 {
		return "TRUCK"
	} else if slotNumber == 1 || slotNumber == 2 {
		return "BIKE"
	} else {
		return "CAR"
	}
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) string {
	for _, floor := range p.Floors {
		for i := 0; i < len(floor.Slots); i++ {
			if !floor.Slots[i].IsOccupied && floor.Slots[i].SlotType == vehicle.Type {
				ticket := &Ticket{
					ID:             fmt.Sprintf("%s_%d_%d", p.ID, floor.FloorNumber, i+1),
					FloorNumber:    floor.FloorNumber,
					SlotNumber:     i + 1,
					VehicleDetails: vehicle,
				}
				floor.Slots[i].IsOccupied = true
				floor.Slots[i].Ticket = ticket
				return ticket.ID
			}
		}
	}
	return "Parking Lot is Full"
}

func (p *ParkingLot) UnparkVehicle(ticketID string) string {
	for _, floor := range p.Floors {
		for i := 0; i < len(floor.Slots); i++ {
			if floor.Slots[i].IsOccupied && floor.Slots[i].Ticket.ID == ticketID {
				regNo := floor.Slots[i].Ticket.VehicleDetails.Registration
				color := floor.Slots[i].Ticket.VehicleDetails.Colour
				floor.Slots[i].IsOccupied = false
				floor.Slots[i].Ticket = nil
				return fmt.Sprintf("Unparked vehicle with Registration Number: %s and Color: %s", regNo, color)
			}
		}
	}
	return "Invalid Ticket"
}

func (p *ParkingLot) DisplayFreeCount(vehicleType string) {
	for _, floor := range p.Floors {
		count := 0
		for _, slot := range floor.Slots {
			if !slot.IsOccupied && slot.SlotType == vehicleType {
				count++
			}
		}
		fmt.Printf("No. of free slots for %s on Floor %d: %d\n", vehicleType, floor.FloorNumber, count)
	}
}

func (p *ParkingLot) DisplayFreeSlots(vehicleType string) {
	for _, floor := range p.Floors {
		slots := []int{}
		for _, slot := range floor.Slots {
			if !slot.IsOccupied && slot.SlotType == vehicleType {
				slots = append(slots, slot.SlotNumber)
			}
		}
		fmt.Printf("Free slots for %s on Floor %d: %v\n", vehicleType, floor.FloorNumber, slots)
	}
}

func (p *ParkingLot) DisplayOccupiedSlots(vehicleType string) {
	for _, floor := range p.Floors {
		slots := []int{}
		for _, slot := range floor.Slots {
			if slot.IsOccupied && slot.SlotType == vehicleType {
				slots = append(slots, slot.SlotNumber)
			}
		}
		fmt.Printf("Occupied slots for %s on Floor %d: %v\n", vehicleType, floor.FloorNumber, slots)
	}
}
