package lemin

import (
	"fmt"
	"strconv"
	"strings"
)

// var antMovement []string

// MoveAnts simulates the movement of ants across the given paths.
func (graph *Graph) MoveAnts(antAmount int, paths []Path) {
	// Track the position of each ant (initially all ants are at the start)
	antPositions := make([]int, antAmount)
	// Track which path each ant is assigned to
	antPaths := make([]int, antAmount)
	// Track if the ant has finished moving
	finished := make([]bool, antAmount)
	// Track which ants have entered their first room
	enteredFirstRoom := make([]bool, antAmount)

	// Assign ants to paths in a round-robin fashion
	for i := 0; i < antAmount; i++ {
		antPaths[i] = i % len(paths)
		if i == antAmount-1 && len(paths) < 3 {
			antPaths[i] = 0
		}
	}

	totalFinished := 0

	// Simulate the steps for the ants to move until all finish
	for totalFinished < antAmount {
		movement := []string{}
		// A map to track which rooms are occupied in this step (only for the first room entry)
		occupiedRooms := map[string]bool{}

		// Simulate movement for each ant
		for antID := 0; antID < antAmount; antID++ {
			// Skip finished ants
			if finished[antID] {
				continue
			}

			pathID := antPaths[antID]
			// Check if the ant has not reached the end of its path
			if antPositions[antID] < len(paths[pathID].Rooms)-1 {
				// If this is the ant's first move, check if the first room is available
				nextRoom := paths[pathID].Rooms[antPositions[antID]+1]
				if !enteredFirstRoom[antID] {
					// Only allow one ant to enter a new first room per step
					if !occupiedRooms[nextRoom] {
						antPositions[antID]++
						currentRoom := paths[pathID].Rooms[antPositions[antID]]
						movement = append(movement, "L"+strconv.Itoa(antID+1)+"-"+currentRoom)
						occupiedRooms[currentRoom] = true
						enteredFirstRoom[antID] = true
					}
				} else {
					// After entering the first room, ants can move freely (no room restrictions)
					// No need to check if it is occupied since the overlapping rooms were already validated
					antPositions[antID]++
					currentRoom := paths[pathID].Rooms[antPositions[antID]]
					movement = append(movement, "L"+strconv.Itoa(antID+1)+"-"+currentRoom)
				}
			} else {
				// Mark ant as finished
				finished[antID] = true
				totalFinished++
			}
		}

		// Print the current step movement
		if len(movement) > 0 {
			fmt.Printf("%s\n", strings.Trim(fmt.Sprint(movement), "[]"))
		}

		// return antMovement
	}
}
