package lemin

import (
	"fmt"
	"math"
)

func (graph *Graph) FindPaths(rooms []Room) map[string]float64 {
	paths := make(map[string]float64)
	visited := make(map[string]bool)
	roomMap := make(map[string]Room)
	for _, room := range rooms {
		roomMap[room.name] = room
	}
	DFS(graph, roomMap, visited, graph.start, []string{}, 0, paths)
	return paths
}

func DFS(graph *Graph, mapRoom map[string]Room, visited map[string]bool, currentRoom string, path []string, totalDistance float64, paths map[string]float64) {
	if visited[currentRoom] {
		return
	}
	visited[currentRoom] = true
	path = append(path, currentRoom)
	if currentRoom == graph.end {
		jpaths := fmt.Sprintf("%v", path)
		paths[jpaths] = totalDistance
	} else {
		for _, connectedRoom := range graph.nodes[currentRoom] {
			dist := ReturnDistance(mapRoom[currentRoom], mapRoom[connectedRoom])
			visited2 := make(map[string]bool)
			for k, v := range visited {
				visited2[k] = v
			}
			DFS(graph, mapRoom, visited2, connectedRoom, path, totalDistance+dist, paths)
		}
	}
}

func ReturnDistance(a, b Room) float64 {
	xf := a.x - b.x
	yf := a.y - b.y
	return math.Sqrt(float64(xf*xf + yf*yf))
}
