package lemin

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
