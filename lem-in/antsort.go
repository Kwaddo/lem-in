package lemin

import (
	"fmt"
	"strconv"
	"strings"
)

/*
After using the DFS algorithm to receive every possible path and using the Validation algorithm to only keep paths that aren't overlapping, we reach this point.
The MoveAnts function uses the amount of ants and the given paths as inputs, and then creates four array variables according to the ant amounts which
focus on the ant's positions, their paths, and to check if the ant finished or entered the first room.

The first for loop is in accordance to the amount of ants, and the idea of it is that it assigns each ant to a certain path through round robbin style. Our
DFS algorithm already sorts the paths from shortest to longest from their lenghts, so ideally we would like the first ant to start from the shortest path, and
the next to go to the second-shortest, and so on. There is an extra feature, however, which focuses on assigning the last ant to the shortest path regardless of
the round robbin method, only if there are only two paths.

After that, we reach the movement section of the function. A for loop is created that is only done when all ants have their paths set out for them. The movement
is just a string array, while a variable called occupied is meant to check if the potential room is occupied by a different ant. After that, a nested loop is
created as a normal for-i loop, where "i" here would be the ID of each ant. First thing it checks is if "finishedchecker" is complete, this is because the for loop
repeats on every ant but the only way it moves on to the next upcoming one is if "finishedchecker" returns true. After that, it begins taking the ants through their
chosen paths.

A new variable named pathID is named such because it is the path in accordance to the current chosen ant, from their ID. Then comes an if statement which checks if
the positions, which would be an int and would intentionally increment every single time the given ant follows through a move, for the ID of the ant is smaller than
the amount of rooms that the chosen path is. If this returns false, then the ant is considered finished and would move on to the next ant.

Within the previous if statement, a new variable is created named "nextRoom" which is focuses on returning the next room only, since the statements end before the
final room. "enteredFirstRoom" is defaulted as false, and the if statement next checks if this current ant has entered the first room. If it did not, the if statement
then follows through to another one which checks if the next room is occupied or not. If it isn't occupied, the following happens:
	1) The position of the ants increment, assuming it's next trajectory is the next room.
	2) A new variable called "currentRoom" is created, and it returns the current ant's room rather than the next one done previously.
	3) Since we want to show the ant going through the current movement, it appends the style (L[ID]-[ROOM]) to the "movement" array.
	4) Since it is the first room and we want to show that it is occupied, we keep them as true in accordance to the occupied map-boolean and "enteredFirstRoom".

That is all if the first room was entered and if the next room wasn't occupied, if it were it would have skipped it and if the first room wasn't entered then it
would assume that it is entering for the first time through the first room, and would do all the steps above disregarding the fourth one. In the end, if there
are any inputs in the "movement" array, then it would be printed in according to this until we reach the final row.
*/
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
