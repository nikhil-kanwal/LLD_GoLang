// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var parkingLot *ParkingLot

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command: ")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]

		if input == "exit" {
			break
		}

		HandleCommand(input)
	}
}
