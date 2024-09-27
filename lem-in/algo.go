package lemin

import (
	"math"
	"sort"
)

type Path struct {
	Rooms    []string 
	Distance float64 
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
    uniquePaths := []Path{}
    roomUsage := make(map[string]int)

    for _, path := range paths {
        overlapScore := 0
        totalInternalRooms := len(path.Rooms) - 2 // Exclude start and end

        // Calculate overlap score
        for _, room := range path.Rooms[1:len(path.Rooms)-1] { // Exclude start and end
            overlapScore += roomUsage[room]
        }

        // Calculate overlap ratio
        overlapRatio := float64(overlapScore) / float64(totalInternalRooms)

        // Allow paths with overlap ratio between 0.1 and 0.5
		count := 0.0
		if len(paths) > 15 {
			count = 0.35
		} else {
			count = 0.55
		}
        if overlapRatio <= count {
            uniquePaths = append(uniquePaths, path)
            
            // Increment usage count for internal rooms in this path
            for _, room := range path.Rooms[1:len(path.Rooms)-1] {
                roomUsage[room]++
            }
        }
    }

    return uniquePaths
}
