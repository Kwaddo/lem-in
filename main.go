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
	
	for _, path := range paths {
		fmt.Printf("Path: %v, Distance: %.12f\n", path.Rooms, path.Distance)
	}
	a := graph.MoveAnts(ants, paths)
	for _, b := range a {
		fmt.Println(b)
	}
}
