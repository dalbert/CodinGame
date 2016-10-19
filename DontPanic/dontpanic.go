package main

import "fmt"

//import "os"

const block = "BLOCK"
const wait = "WAIT"

func main() {
	// nbFloors: number of floors
	// width: width of the area
	// nbRounds: maximum number of rounds
	// exitFloor: floor on which the exit is found
	// exitPos: position of the exit on its floor
	// nbTotalClones: number of generated clones
	// nbAdditionalElevators: ignore (always zero)
	// nbElevators: number of elevators
	var nbFloors, width, nbRounds, exitFloor, exitPos, nbTotalClones, nbAdditionalElevators, nbElevators int
	fmt.Scan(&nbFloors, &width, &nbRounds, &exitFloor, &exitPos, &nbTotalClones, &nbAdditionalElevators, &nbElevators)

	elevators := make(map[int]int, nbElevators)
	for i := 0; i < nbElevators; i++ {
		// elevatorFloor: floor on which this elevator is found
		// elevatorPos: position of the elevator on its floor
		var elevatorFloor, elevatorPos int
		fmt.Scan(&elevatorFloor, &elevatorPos)
		elevators[elevatorFloor] = elevatorPos
	}
	for {
		// cloneFloor: floor of the leading clone
		// clonePos: position of the leading clone on its floor
		// direction: direction of the leading clone: LEFT or RIGHT
		var cloneFloor, clonePos int
		var direction, command string
		fmt.Scan(&cloneFloor, &clonePos, &direction)
		if cloneFloor == exitFloor {
			if clonePos > exitPos {
				if direction == "RIGHT" {
					command = block
				} else {
					command = wait
				}
			} else {
				if direction == "LEFT" {
					command = block
				} else {
					command = wait
				}
			}
		} else {
			if clonePos > elevators[cloneFloor] {
				if direction == "RIGHT" {
					command = block
				} else {
					command = wait
				}
			} else if clonePos < elevators[cloneFloor] {
				if direction == "LEFT" {
					command = block
				} else {
					command = wait
				}
			} else { // clonePos == elevatorPos
				command = wait
			}
		}
		fmt.Println(command)
	}
}
