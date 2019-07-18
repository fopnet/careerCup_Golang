package graph

func (g graphImpl) GetEdge(node string) []edge {
	return g.edges[node]
}


func (g *graphImpl) GetShortestPath(origin, destiny string) (int, []string, map[string]int) {
	queue := newHeap()
	isVisitedMap := make(map[string]bool)
	dist := map[string]int{origin: 0}
	queue.push(path{value: dist[origin], nodes: []string{origin}})

	var destinyPath path
	// for len(isVisitedMap) != g.maxNodes {
	for !queue.isEmpty() {
		// encontrar o mais próximo while não estiver visitado
		currPath := queue.pop()
		currNode := currPath.nodes[len(currPath.nodes)-1]

		if isVisitedMap[currNode] {
			continue
		}
		if currNode == destiny {
			// return currPath.value, currPath.nodes
			destinyPath = currPath
		}

		for _, currEdge := range g.GetEdge(currNode) {
			if !isVisitedMap[currEdge.node] {
				newDistance := dist[currNode] + currEdge.weight

				if dist[currEdge.node] == 0 || newDistance < dist[currEdge.node] {
					dist[currEdge.node] = newDistance
				}

				path := path{value: dist[currEdge.node], nodes: append(currPath.nodes, currEdge.node)}
				queue.push(path)
			}
		}

		isVisitedMap[currNode] = true
	}

	if &destinyPath == nil {
		return 0, nil, dist
	} else {
		return destinyPath.value, destinyPath.nodes, dist
	}
}
