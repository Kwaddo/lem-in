package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Path struct {
	Rooms    []string
	Distance float64
}

type Graph struct {
	nodes map[string][]string
	start string
	end   string
}

type Room struct {
	name string
	x, y int
}

func main() {
	var graph Graph

	file, err := os.OpenFile("tester/testerOuput.txt", os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 10; i++ {
		inputFile := fmt.Sprintf("tester/examples/example0%v.txt", i)
		ants, rooms, err := graph.ParseInput(inputFile)
		if err != nil {
			file.WriteString(err.Error())
			file.WriteString("\n\n------------------------------------------------------------------------------------------------------------\n\n")
			continue
		}

		paths := graph.FindPaths(rooms)

		allMovement := MoveAnts(ants, paths)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := 0; i < len(allMovement); i++ {
			_, err = file.WriteString(fmt.Sprintf("%s\n", allMovement[i]))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		_, err = file.WriteString("\n\nNumber of Turns: " + strconv.Itoa(len(allMovement)) + "\n\n------------------------------------------------------------------------------------------------------------\n\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	file.Close()

}

func (graph *Graph) FindPaths(rooms []Room) []Path {
	paths := []Path{}
	visited := make(map[string]bool)
	roomMap := make(map[string]Room)
	for _, room := range rooms {
		roomMap[room.name] = room
	}
	DFS(graph, roomMap, visited, graph.start, []string{}, 0, &paths)
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].Distance < paths[j].Distance
	})
	if len(paths) >= 3 {
		paths = ValidatePaths(paths)
	}
	return paths
}

func DFS(graph *Graph, roomMap map[string]Room, visited map[string]bool, currentRoom string, path []string, totalDistance float64, paths *[]Path) {
	if visited[currentRoom] {
		return
	}
	visited[currentRoom] = true
	path = append(path, currentRoom)
	if currentRoom == graph.end {
		*paths = append(*paths, Path{Rooms: append([]string{}, path...), Distance: totalDistance})
	} else {
		for _, connectedRoom := range graph.nodes[currentRoom] {
			dist := ReturnDistance(roomMap[currentRoom], roomMap[connectedRoom])
			visited2 := make(map[string]bool)
			for k, v := range visited {
				visited2[k] = v
			}
			DFS(graph, roomMap, visited2, connectedRoom, path, totalDistance+dist, paths)
		}
	}
}

func ReturnDistance(a, b Room) float64 {
	xf := a.x - b.x
	yf := a.y - b.y
	return math.Sqrt(float64(xf*xf + yf*yf))
}

func ValidatePaths(paths []Path) []Path {
	uPaths := []Path{}
	roomUsedCount := make(map[string]int)
	for _, path := range paths {
		overlapScore := 0
		totalInternalRooms := len(path.Rooms) - 2
		for _, room := range path.Rooms[1 : len(path.Rooms)-1] {
			overlapScore += roomUsedCount[room]
		}
		overlapRatio := float64(overlapScore) / float64(totalInternalRooms)
		if overlapRatio <= 0.3 {
			uPaths = append(uPaths, path)
			for _, room := range path.Rooms[1 : len(path.Rooms)-1] {
				roomUsedCount[room]++
			}
		}
	}
	return SmallestPaths(CullPaths(uPaths))
}

func CullPaths(paths []Path) []Path {
	culledPaths := []Path{}
	for _, path := range paths {
		shouldAdd := true
		for _, culledPath := range culledPaths {
			minLength := len(path.Rooms)
			if len(culledPath.Rooms) < minLength {
				minLength = len(culledPath.Rooms)
			}
			for i := 1; i < minLength-1; i++ {
				if path.Rooms[i] == culledPath.Rooms[i] {
					shouldAdd = false
					break
				}
			}
			if !shouldAdd {
				break
			}
		}
		if shouldAdd {
			culledPaths = append(culledPaths, path)
		}
	}
	return culledPaths
}

func SmallestPaths(paths []Path) []Path {
	var sPaths []Path
	for _, path := range paths {
		if len(path.Rooms) == 2 || len(path.Rooms) == 3 {
			sPaths = append(sPaths, path)
			break
		} else {
			sPaths = append(sPaths, path)
		}
	}
	return sPaths
}

func MoveAnts(antAmount int, paths []Path) []string {
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
	var allMovement []string

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
		allMovement = append(allMovement, strings.Join(movement, " "))

	}
	return allMovement
}

func (graph *Graph) ParseInput(filename string) (int, []Room, error) {
	if !(strings.HasSuffix(filename, ".txt")) {
		return 0, []Room{}, fmt.Errorf("ERROR: Invalid File Extension, Should be [.txt] file")
	}
	byteData, err := os.ReadFile(filename)
	if err != nil {
		return 0, []Room{}, fmt.Errorf("ERROR: File Not Found, Please enter a valid file name")
	}
	lines := strings.Split(string(byteData), "\r\n")
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
		line := i + 1
		if lines[i] != "" {
			ants, err = strconv.Atoi(lines[i])
			if err != nil {
				return 0, []Room{}, fmt.Errorf("ERROR: Invalid input, Enter number of ants")
			}

			index = i + 1

			if ants > 1500000 || ants < 1 {
				return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid input, Enter number of ants [ 1 - 150000 ]", line)
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

			if parts[0] == parts[1] {
				return 0, []Room{}, fmt.Errorf("ERROR AT LINE %v: Invalid input, Room is connected to it self", line)
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
