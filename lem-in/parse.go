package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
The idea of the parse is a checker, to see if there are any conflicting connections, any wrong values, any repetitions, and any formats that completely go against
what should follow through.
*/

type Graph struct {
	nodes map[string][]string
	start string
	end   string
}

type Room struct {
	name string
	x, y int
}

func (graph *Graph) ParseInput(filename string) (int, []Room, error) {
	if !(strings.HasSuffix(filename, ".txt")) {
		return 0, []Room{}, fmt.Errorf("ERROR: Invalid File Extension, Should be [.txt] file")
	}
	byteData, err := os.ReadFile(filename)
	if err != nil {
		return 0, []Room{}, fmt.Errorf("ERROR: File Not Found, Please enter a valid file name")
	}
	lines := strings.Split(string(byteData), "\n")
	if len(lines) < 6 {
		return 0, []Room{}, fmt.Errorf("ERROR: Invalid input, Not enough input")
	}
	start := false
	end := false
	var rooms []Room
	var startRoomArr []string
	var endRoomArr []string
	graph.nodes = make(map[string][]string)
	var ants int
	var index int
	for i := 0; i < len(lines); i++ {
		line := i+1
		if lines[i] != "" {
			ants, err = strconv.Atoi(lines[i])
			if err != nil {
				return 0, []Room{}, fmt.Errorf("ERROR: Invalid input, Enter number of ants")
			}

			index = i + 1

			if ants > 5000000 || ants < 1 {
				return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid input, Enter number of ants [ 1 - 500000 ]", line)
			}
			break
		}
	}
	for i := index; i < len(lines); i++ {
		line := i+1
		if lines[i] == "" {
			continue
		} else if strings.HasPrefix(lines[i], "##") {
			if i+1 < len(lines) {
				if lines[i] == "##start" && !start {
					start = true
					startRoomArr = strings.Split(lines[i+1], " ")
					if len(startRoomArr) != 3 {
						for j := i+2; j < len(lines); j++ {
							startRoomArr = strings.Split(lines[j], " ")
							if len(startRoomArr) != 3 {
								continue
							} else {
								break
							}
						}
					}
					graph.start = startRoomArr[0]
				} else if lines[i] == "##end" && !end && start {
					end = true
					endRoomArr = strings.Split(lines[i+1], " ")
					if len(endRoomArr) != 3 {
						for j := i+2; j < len(lines); j++ {
							endRoomArr = strings.Split(lines[j], " ")
							if len(endRoomArr) != 3 {
								continue
							} else {
								break
							}
						}
					}
					graph.end = endRoomArr[0]
				} else if lines[i] == "##start" && start {
					return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid Input, Redundant ##start flag", line)
				} else if lines[i] == "##end" && end {
					return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid Input, Redundant ##end flag", line)
				} else {
					return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid Input, Wrong flag entered", line)
				}
				continue
			}
		} else if strings.HasPrefix(lines[i], "#") {
			continue
		} else if strings.Contains(lines[i], "-") {
			negChecker := strings.Split(lines[i], " ")
			if len(negChecker) > 1 {
				if _, err := strconv.Atoi(negChecker[1]); err != nil {
					return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid input, Spaces between connections input", line)
				} else {
					x, err := strconv.Atoi(negChecker[1])
					if err != nil {
						return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid X value input", line)
					}
					y, err := strconv.Atoi(negChecker[2])
					if err != nil {
						return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid Y value input", line)
					}

					if x < 0 && y < 0 {
						return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid input, Negative X [AND] Y values", line)
					} else if x < 0 {
						return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid input, Negative X value", line)
					} else if y < 0 {
						return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid input, Negative Y value", line)
					}
				}
			}
			parts := strings.Split(lines[i], "-")
			if len(parts) > 2 {
				return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid input, Only two rooms can be connected", line)
			}
			graph.nodes[parts[0]] = append(graph.nodes[parts[0]], parts[1])
			graph.nodes[parts[1]] = append(graph.nodes[parts[1]], parts[0])
			for _, arr := range graph.nodes {
				for j := 0; j < len(arr); j++ {
					for k := j + 1; k < len(arr); k++ {
						if arr[j] == arr[k] {
							return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid input, Same connection appeared more than once", line)
						}
					}
				}
			}
		} else {
			room, err := ParseRoom(lines[i])
			if err != nil {
				return 0, []Room{}, err
			}
			rooms = append(rooms, room)
		}
	}
	found := false
	var unconnected Room
	for _, room := range rooms {
		for k := range graph.nodes {
			if room.name == k {
				found = true
				break
			} else {
				found = false
				unconnected = room
			}
		}
		if !found {
			return 0, []Room{}, fmt.Errorf("ERROR: Invalid input, Room %v is Unconnected", unconnected)
		}
	}
	notEx := ""
	for k := range graph.nodes {
		for _, room := range rooms {
			if room.name == k {
				found = true
				break
			} else {
				found = false
				notEx = k
			}
		}
		if !found {
			return 0, []Room{}, fmt.Errorf("ERROR: Invalid input, Room [%v] connected and does not exist", notEx)
		}
	}
	for i := 0; i < len(rooms); i++ {
		for j := i + 1; j < len(rooms); j++ {
			if rooms[i].x == rooms[j].x && rooms[i].y == rooms[j].y {
				return 0, []Room{}, fmt.Errorf("ERROR: Invalid input, Rooms %v [AND] %v have the same X and Y values", rooms[i], rooms[j])
			} else if rooms[i].name == rooms[j].name {
				return 0, []Room{}, fmt.Errorf("ERROR: Invalid input, Rooms %v [AND] %v have the same name", rooms[i], rooms[j])
			}
		}
	}
	if !start {
		return 0, []Room{}, fmt.Errorf("ERROR: Invalid input, ##start is missing")
	} else if !end {
		return 0, []Room{}, fmt.Errorf("ERROR: Invalid input, ##end is missing")
	}
	return ants, rooms, nil
}

func ParseRoom(line string) (Room, error) {
	room := strings.Split(line, " ")
	if len(room) != 3 {
		return Room{}, fmt.Errorf("ERROR: Invalid input, Room input should be [RoomName X Y]")
	}
	x, err := strconv.Atoi(room[1])
	if err != nil {
		return Room{}, err
	}
	y, err := strconv.Atoi(room[2])
	if err != nil {
		return Room{}, err
	}
	name := room[0]
	return Room{name: name, x: x, y: y}, nil
}
