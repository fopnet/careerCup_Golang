package graph

func NewUndirectedGraph(maxNodes int) Graph {
	return &unDirectedGraphImpl{graphImpl{edges: make(map[string][]edge), maxNodes: maxNodes}}
}

func (g *graphImpl) AddEdge(origin, destiny string, w int) Graph {
	g.edges[origin] = append(g.edges[origin], edge{node: destiny, weight: w})
	g.edges[destiny] = append(g.edges[destiny], edge{node: origin, weight: w})
	return g
}
