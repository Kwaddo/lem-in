package lemin

import (
	"math"
	"sort"
)

/*
First thing this is set out to make is an array of paths, so that every validated and confirmed path is thrown into it and used accordingly for the antmove
function. The function FindPaths focuses on finding every possible path, and within it validates the paths if there are three or mnore possible paths
to begin with since there doesn't need to be any validation if there are only two possible paths.

By using DFS, the given function below, it focuses mainly on returning every possible path with no discrimination. Within it is a function that returns the
distance between the two rooms, which is in accordance to pythogoreas's theorem. Once all paths are found, they are to be validated.

The ValidatePaths function is a slightly complex algorithm that only keeps paths that do not overlap each other, keeping each of them unique. "uPaths" is the
Paths struct array that has only the ones that are chosen. "roomUsedCount" is a map that is meant to keep the rooms that are used in their specific places,
for rememberance that it does not overlap in the future. It checks how many times they are used and keeps a count. A for-loop is made for the paths, to
check each path accordingly one-by-one.

Within the for-loop is the overlap score, which has the total from "roomUsedCount" so that it can be used soon. The "overlapRatio" focuses on dividing the
overlap score by the total amount of internal rooms, obviously disregarding the start and end rooms since they do not count. After constant trial and error,
we found the best ratio for getting only the unique paths to be smaller than or equal to 0.3. This is because within that given path, there are a small amount
of overlaps and by making the number bigger we are only allowing for more overlaps, so 0.3 is meant to be that sweet spot. Within the if statement, it adds the
path to "uPaths" and then keeps another for loop to increase the "roomUsedCount". Once all of it is done, it returns the validated paths only.
*/

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
    uPaths := []Path{}
    roomUsedCount := make(map[string]int)
    for _, path := range paths {
        overlapScore := 0
        totalInternalRooms := len(path.Rooms) - 2
        for _, room := range path.Rooms[1:len(path.Rooms)-1] { 
            overlapScore += roomUsedCount[room]
        }
        overlapRatio := float64(overlapScore) / float64(totalInternalRooms)
        if overlapRatio <= 0.3 {
            uPaths = append(uPaths, path)
            for _, room := range path.Rooms[1:len(path.Rooms)-1] {
                roomUsedCount[room]++
            }
        }
    }
    return CullPaths(uPaths)
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
