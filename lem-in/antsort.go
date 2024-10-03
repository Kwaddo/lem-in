package lemin

import (
	"fmt"
	"strconv"
	"strings"
)

func MoveAnts(antAmount int, paths []Path) {
	positions := make([]int, antAmount)
	aPaths := make([]int, antAmount)
	finishedchecker := make([]bool, antAmount)
	enteredFirstRoom := make([]bool, antAmount)
	for i := 0; i < antAmount; i++ {
		aPaths[i] = i % len(paths)
		if i == antAmount-1 && len(paths) < 3 {
			count := 0
			for j := 1; j < len(paths); j++ {
				if len(paths[j].Rooms) < len(paths[0].Rooms) {
					count = j
				}
			}
			aPaths[i] = count
		}
	}
	totalFinished := 0
	for totalFinished < antAmount {
		var movement []string
		occupied := map[string]bool{}
		for i := 0; i < antAmount; i++ {
			if finishedchecker[i] {
				continue
			}
			pathID := aPaths[i]
			if positions[i] < len(paths[pathID].Rooms)-1 {
				nextRoom := paths[pathID].Rooms[positions[i]+1]
				if !enteredFirstRoom[i] {
					if !occupied[nextRoom] {
						positions[i]++
						currentRoom := paths[pathID].Rooms[positions[i]]
						movement = append(movement, "L"+strconv.Itoa(i+1)+"-"+currentRoom)
						occupied[currentRoom] = true
						enteredFirstRoom[i] = true
					}
				} else {
					positions[i]++
					currentRoom := paths[pathID].Rooms[positions[i]]
					movement = append(movement, "L"+strconv.Itoa(i+1)+"-"+currentRoom)
				}
			} else {
				finishedchecker[i] = true
				totalFinished++
			}
		}
		if len(movement) > 0 {
			fmt.Printf("%s\n", strings.Trim(fmt.Sprint(movement), "[]"))
		}
	}
}
