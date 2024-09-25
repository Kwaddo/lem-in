package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		return 0, []Room{}, fmt.Errorf("error: Invalid file extension, Please enter a [ .txt ] file extension")
	}

	byteData, err := os.ReadFile(filename)
	if err != nil {
		return 0, []Room{}, fmt.Errorf("error: File not found, Please enter a valid file name")
	}

	lines := strings.Split(string(byteData), "\r\n")
	if len(lines) < 6 {
		return 0, []Room{}, fmt.Errorf("error: Invalid file input, Not enough input")
	}

	start := false
	end := false
	var rooms []Room
	startRoom := ""
	endRoom := ""
	graph.nodes = make(map[string][]string)
	var ants int
	var index int

	for i := 0; i < len(lines); i++ {
		if lines[i] != "" {
			ants, err = strconv.Atoi(lines[i])
			if err != nil {
				return 0, []Room{}, fmt.Errorf("please enter number of ants")
			}

			index = i + 1

			if ants > 1000 || ants < 1 {
				return 0, []Room{}, fmt.Errorf("please enter number of ants [ 1 - 1000 ]")
			}
			break
		}
	}

	for i := index; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		} else if strings.HasPrefix(lines[i], "##") {
			if i+1 < len(lines) {
				if lines[i] == "##start" && !start {
					start = true
					startRoom = strings.Split(lines[i+1], " ")[0]
					graph.start = startRoom
				} else if lines[i] == "##end" && !end {
					end = true
					endRoom = strings.Split(lines[i+1], " ")[0]
					graph.end = endRoom
				} else {
					return 0, []Room{}, fmt.Errorf("\nPlease make sure you entered the right flags\n\nOR\n\nMake sure there is no redundent flags")
				}
				continue
			}
		} else if strings.HasPrefix(lines[i], "#") {
			continue
		} else if strings.Contains(lines[i], "-") {
			negChecker := strings.Split(lines[i], " ")
			if len(negChecker) > 1 {
				return 0, []Room{}, fmt.Errorf("error: invalid input")
			}
			parts := strings.Split(lines[i], "-")
			fmt.Println(parts)
			if len(parts) > 2 {
				return 0, []Room{}, fmt.Errorf("error: invalid input")
			}
			graph.nodes[parts[0]] = append(graph.nodes[parts[0]], parts[1])
			graph.nodes[parts[1]] = append(graph.nodes[parts[1]], parts[0])
			for _, arr := range graph.nodes {
				for i := 0; i < len(arr); i++ {
					for j := i+1; j < len(arr); j++ {
						if arr[i] == arr[j] {
							return 0, []Room{}, fmt.Errorf("error: duplicated room connections")
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

	for _, room := range rooms {
		for k := range graph.nodes {
			if room.name == k {
				found = true
				break
			} else {
				found = false
			}
		}
		if !found {
			return 0, []Room{}, fmt.Errorf("error: unconnected room")
		}
	}

	for k := range graph.nodes {
		for _, room := range rooms {
			if room.name == k {
				found = true
				break
			} else {
				found = false
			}
		}
		if !found {
			return 0, []Room{}, fmt.Errorf("error: connected room that does not exist")
		}
	}

	for i := 0; i < len(rooms); i++ {
		for j := i + 1; j < len(rooms); j++ {
			if rooms[i].x == rooms[j].x && rooms[i].y == rooms[j].y {
				return 0, []Room{}, fmt.Errorf("please make sure you entered unique coordinates for each room")
			}
		}
	}

	fmt.Println(graph.nodes)
	if !start || !end {
		return 0, []Room{}, fmt.Errorf("error: ##start OR ##end flag was not found")
	}

	return ants, rooms, nil

}

func ParseRoom(line string) (Room, error) {
	room := strings.Split(line, " ")
	if len(room) != 3 {
		return Room{}, fmt.Errorf("error: invalid room input")
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
