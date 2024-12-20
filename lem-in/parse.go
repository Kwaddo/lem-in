package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (graph *Graph) ParseInput(filename string) (int, []Room, []error) {
	if !(strings.HasSuffix(filename, ".txt")) {
		return 0, []Room{}, []error{fmt.Errorf("ERROR: Invalid File Extension, Should be [.txt] file")}
	}
	byteData, err := os.ReadFile(filename)
	if err != nil {
		return 0, []Room{}, []error{fmt.Errorf("ERROR: File Not Found, Please enter a valid file name")}
	}
	lines := strings.Split(string(byteData), "\r\n")
	if len(lines) < 6 {
		return 0, []Room{}, []error{fmt.Errorf("ERROR: Invalid input, Not enough input")}
	}
	start := false
	end := false
	var rooms []Room
	var startRoomArr []string
	var endRoomArr []string
	graph.nodes = make(map[string][]string)
	var ants int
	var index int
	var allErrors []error
	for i := 0; i < len(lines); i++ {
		line := i + 1
		if lines[i] != "" {
			ants, err = strconv.Atoi(lines[i])
			if err != nil {
				allErrors = append(allErrors, fmt.Errorf("ERROR: Invalid input, Enter number of ants"))
				break
			}

			index = i + 1

			if ants > 150000 || ants < 1 {
				allErrors = append(allErrors, fmt.Errorf("ERROR AT LINE %v: Invalid input, Enter number of ants [ 1 - 150000 ]", line))
				break
			}
			
			break
		}
	}
	for i := index; i < len(lines); i++ {
		line := i + 1
		if lines[i] == "" {
			continue
		} else if strings.HasPrefix(lines[i], "##") {
			if i+1 < len(lines) {
				if lines[i] == "##start" && !start {
					start = true
					startRoomArr = strings.Split(lines[i+1], " ")
					if len(startRoomArr) != 3 {
						for j := i + 2; j < len(lines); j++ {
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
						for j := i + 2; j < len(lines); j++ {
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
					allErrors = append(allErrors, fmt.Errorf("ERROR AT LINE %v: Invalid Input, Redundant ##start flag", line))
				} else if lines[i] == "##end" && end {
					allErrors = append(allErrors, fmt.Errorf("ERROR AT LINE %v: Invalid Input, Redundant ##end flag", line))
				} else {
					allErrors = append(allErrors, fmt.Errorf("ERROR AT LINE %v: Invalid Input, Wrong flag entered", line))
				}
				continue
			}
		} else if strings.HasPrefix(lines[i], "#") {
			continue
		} else if strings.Contains(lines[i], "-") {
			negChecker := strings.Split(lines[i], " ")
			if len(negChecker) > 1 {
				if _, err := strconv.Atoi(negChecker[1]); err != nil {
					allErrors = append(allErrors, fmt.Errorf("ERROR AT LINE %v: Invalid input, Spaces between connections input", line))
				} else {
					x, err := strconv.Atoi(negChecker[1])
					if err != nil {
						allErrors = append(allErrors, fmt.Errorf("ERROR AT LINE %v: Invalid X value input", line))
						continue
					}
					y, err := strconv.Atoi(negChecker[2])
					if err != nil {
						allErrors = append(allErrors, fmt.Errorf("ERROR AT LINE %v: Invalid Y value input", line))
						continue
					}

					if x < 0 && y < 0 {
						allErrors = append(allErrors, fmt.Errorf("ERROR AT LINE %v: Invalid input, Negative X [AND] Y values", line))
						continue
					} else if x < 0 {
						allErrors = append(allErrors, fmt.Errorf("ERROR AT LINE %v: Invalid input, Negative X value", line))
						continue
					} else if y < 0 {
						allErrors = append(allErrors, fmt.Errorf("ERROR AT LINE %v: Invalid input, Negative Y value", line))
						continue
					}
				}
			}
			parts := strings.Split(lines[i], "-")
			if len(parts) > 2 {
				allErrors = append(allErrors, fmt.Errorf("ERROR AT LINE %v: Invalid input, Only two rooms can be connected", line))
				continue
			}

			if parts[0] == parts[1] {
				allErrors = append(allErrors, fmt.Errorf("ERROR AT LINE %v: Invalid input, Room is connected to it self", line))
				continue
			}

			graph.nodes[parts[0]] = append(graph.nodes[parts[0]], parts[1])
			graph.nodes[parts[1]] = append(graph.nodes[parts[1]], parts[0])

			for _, conn := range graph.nodes {
				for j := 0; j < len(conn); j++ {
					for k := j + 1; k < len(conn); k++ {
						if conn[j] == conn[k] {
							return 0, []Room{}, []error{fmt.Errorf("ERROR AT LINE %v: Invalid input, Same connection appeared more than once", line)}
						}
					}
				}
			}
		} else {
			room, err := ParseRoom(lines[i])
			if err != nil {
				allErrors = append(allErrors, err)
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
			allErrors = append(allErrors, fmt.Errorf("ERROR: Invalid input, Room %v is Unconnected", unconnected))
			continue
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
			allErrors = append(allErrors, fmt.Errorf("ERROR: Invalid input, Room [%v] connected and does not exist", notEx))
			continue
		}
	}

	for i := 0; i < len(rooms); i++ {
		for j := i + 1; j < len(rooms); j++ {
			if rooms[i].x == rooms[j].x && rooms[i].y == rooms[j].y {
				allErrors = append(allErrors, fmt.Errorf("ERROR: Invalid input, Rooms %v [AND] %v have the same X and Y values", rooms[i], rooms[j]))
			} else if rooms[i].name == rooms[j].name {
				allErrors = append(allErrors, fmt.Errorf("ERROR: Invalid input, Rooms %v [AND] %v have the same name", rooms[i], rooms[j]))
			}
		}
	}
	if !start && !end {
		allErrors = append(allErrors, fmt.Errorf("ERROR: Invalid input, ##start and ##end flags are missing"))
	} else if !start {
		allErrors = append(allErrors, fmt.Errorf("ERROR: Invalid input, ##start is missing"))
	} else if !end {
		allErrors = append(allErrors, fmt.Errorf("ERROR: Invalid input, ##end is missing"))
	}

	if len(allErrors) == 0 {
		return ants, rooms, nil
	} else {
		return 0, []Room{}, allErrors
	}

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
