package main

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

	ants, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, []Room{}, fmt.Errorf("please enter number of ants")
	}

	if ants > 1000 || ants < 1 {
		return 0, []Room{}, fmt.Errorf("please enter number of ants [ 1 - 1000 ]")
	}

	for i := 1; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "##") {
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

		} else if strings.Contains(lines[i], "-") {
			parts := strings.Split(lines[i], "-")
			if len(parts) > 2 {
				return 0, []Room{}, fmt.Errorf("error: invalid input")
			}
			graph.nodes[parts[0]] = append(graph.nodes[parts[0]], parts[1])
			graph.nodes[parts[1]] = append(graph.nodes[parts[1]], parts[0])
			continue
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

func main() {
	var graph Graph
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please enter a file name")
		return
	}
	ants, rooms, err := graph.ParseInput(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ants)
	fmt.Println(graph)
	fmt.Println(rooms)

}
