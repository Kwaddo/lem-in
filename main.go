package main

import (
	"os"
	li "lemin/lem-in"
	"fmt"

)

func main() {
	var graph li.Graph
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please enter a file name")
		return
	}
	fmt.Println(args[0])
	ants, rooms, err := graph.ParseInput(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ants)
	fmt.Println(graph)
	fmt.Println(rooms)
	paths := graph.FindPaths(rooms)

	// Sort the paths by distance.
	// sortedPaths := li.SortPaths(paths)

	// Output the sorted paths and their distances.
	for p, d := range paths {
		fmt.Printf("Path: %s, Distance: %.5f\n", p, d)
	}
}