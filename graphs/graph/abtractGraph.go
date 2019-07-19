package graph

func (g graphImpl) GetEdge(node string) []edge {
	return g.edges[node]
}

func (g *graphImpl) GetShortestPath(origin, destiny string) (int, []string, map[string]int) {
	queue := newHeap()
	isVisitedMap := make(map[string]bool)
	dist := map[string]int{origin: 0}
	queue.push(path{node: origin, value: dist[origin], nodes: []string{origin}})

	var destinyPath path
	// for len(isVisitedMap) != g.maxNodes {
	for !queue.isEmpty() {
		// encontrar o mais próximo while não estiver visitado
		currPath := queue.pop()
		// último edge adicionado no caminho
		// currNode := currPath.nodes[len(currPath.nodes)-1]

		if isVisitedMap[currPath.node] {
			continue
		}
		if currPath.node == destiny {
			// return currPath.value, currPath.nodes
			destinyPath = currPath
		}

		for _, currEdge := range g.GetEdge(currPath.node) {
			if !isVisitedMap[currEdge.node] {
				newDistance := dist[currPath.node] + currEdge.weight

				if dist[currEdge.node] == 0 || newDistance < dist[currEdge.node] {
					dist[currEdge.node] = newDistance

					path := path{node: currEdge.node, value: dist[currEdge.node], nodes: append(currPath.nodes, currEdge.node)}

					queue.push(path)
				}

			}
		}

		isVisitedMap[currPath.node] = true
	}

	if &destinyPath == nil {
		return 0, nil, dist
	} else {
		return destinyPath.value, destinyPath.nodes, dist
	}
}
