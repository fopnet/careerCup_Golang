package graph

type edge struct {
	node   string
	weight int
}

type graphImpl struct {
	edges    map[string][]edge
	maxNodes int
}

type directedGraphImpl struct {
	graphImpl
}

type unDirectedGraphImpl struct {
	graphImpl
}

type Graph interface {
	AddEdge(origin, destiny string, w int) Graph
	GetEdge(node string) []edge
	GetShortestPath(origin, destiny string) (int, []string, map[string]int)
}
