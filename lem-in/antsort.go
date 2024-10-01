package lemin

import (
	"fmt"
	"strconv"
	"strings"
)

func (graph *Graph) MoveAnts(antAmount int, paths []Path) {
	antPositions := make([]int, antAmount)
	antPaths := make([]int, antAmount)
	finished := make([]bool, antAmount)
	enteredFirstRoom := make([]bool, antAmount)
	for i := 0; i < antAmount; i++ {
		antPaths[i] = i % len(paths)
		if i == antAmount-1 && len(paths) < 3 {
			shortestPathIndex := 0
			shortestPathLength := len(paths[0].Rooms)
			for j := 1; j < len(paths); j++ {
				if len(paths[j].Rooms) < shortestPathLength {
					shortestPathIndex = j
					shortestPathLength = len(paths[j].Rooms)
				}
			}
			antPaths[i] = shortestPathIndex
		}
	}
	totalFinished := 0
	for totalFinished < antAmount {
		movement := []string{}
		occupiedRooms := map[string]bool{}
		for antID := 0; antID < antAmount; antID++ {
			if finished[antID] {
				continue
			}
			pathID := antPaths[antID]
			if antPositions[antID] < len(paths[pathID].Rooms)-1 {
				nextRoom := paths[pathID].Rooms[antPositions[antID]+1]
				if !enteredFirstRoom[antID] {
					if !occupiedRooms[nextRoom] {
						antPositions[antID]++
						currentRoom := paths[pathID].Rooms[antPositions[antID]]
						movement = append(movement, "L"+strconv.Itoa(antID+1)+"-"+currentRoom)
						occupiedRooms[currentRoom] = true
						enteredFirstRoom[antID] = true
					}
				} else {
					antPositions[antID]++
					currentRoom := paths[pathID].Rooms[antPositions[antID]]
					movement = append(movement, "L"+strconv.Itoa(antID+1)+"-"+currentRoom)
				}
			} else {
				finished[antID] = true
				totalFinished++
			}
		}
		if len(movement) > 0 {
			fmt.Printf("%s\n", strings.Trim(fmt.Sprint(movement), "[]"))
		}
	}
}
