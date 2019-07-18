package graph

func NewDirectedGraph(maxNodes int) Graph {
	return &directedGraphImpl{graphImpl{edges: make(map[string][]edge), maxNodes: maxNodes}}
}

func (g *directedGraphImpl) AddEdge(origin, destiny string, w int) Graph {
	g.edges[origin] = append(g.edges[origin], edge{node: destiny, weight: w})
	return g
}
