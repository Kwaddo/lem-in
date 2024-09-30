package main

import (
	"fmt"
	li "lemin/lem-in"
	"os"
)

func main() {
	var graph li.Graph
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Print("\nERROR: Invalid number of arguments\n\nUSAGE: go run main.go [INPUT FILENAME]\n")
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
	
	for _, path := range paths {
		fmt.Printf("Path: %v, Distance: %.12f\n", path.Rooms, path.Distance)
	}
	graph.MoveAnts(ants, paths)
}
