package graphExamples

/**
https://levelup.gitconnected.com/simple-implementation-of-dijkstra-using-heap-in-go-fa6bea892909
ATENÇÃO : Este código não funciona corretamente para o exemplo abaixo

https://www.geeksforgeeks.org/dijkstras-shortest-path-algorithm-in-java-using-priorityqueue/
*/
import (
	hp "container/heap"

	"fmt"
)

type edge struct {
	node   string
	weight int
}

type graph struct {
	nodes map[string][]edge
}

func newGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) addEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
	g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

func (g *graph) getPath(origin, destiny string) (int, []string) {
	h := newHeap()
	h.push(path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.value, p.nodes
		}

		for _, e := range g.getEdges(node) {
			if !visited[e.node] {
				// We calculate the total spent so far plus the cost and the path of getting here
				h.push(path{value: p.value + e.weight, nodes: append(p.nodes, e.node)})
			}
		}

		visited[node] = true
	}

	return 0, nil
}

type path struct {
	value int
	nodes []string
}

type minPath []path

func (h minPath) Len() int           { return len(h) }
func (h minPath) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minPath) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}

func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type heap struct {
	values *minPath
}

func newHeap() *heap {
	return &heap{values: &minPath{}}
}

func (h *heap) push(p path) {
	hp.Push(h.values, p)
}

func (h *heap) pop() path {
	i := hp.Pop(h.values)
	return i.(path)
}

func main() {
	fmt.Println("Dijkstra")
	// Example
	graph := newGraph()
	// graph.addEdge("S", "B", 4)
	// graph.addEdge("S", "C", 2)
	// graph.addEdge("B", "C", 1)
	// graph.addEdge("B", "D", 5)
	// graph.addEdge("C", "D", 8)
	// graph.addEdge("C", "E", 10)
	// graph.addEdge("D", "E", 2)
	// graph.addEdge("D", "T", 6)
	// graph.addEdge("E", "T", 2)

	graph.addEdge("A", "B", 4)
	graph.addEdge("A", "H", 8)

	graph.addEdge("B", "H", 11)
	graph.addEdge("B", "C", 8)

	graph.addEdge("H", "I", 7)
	graph.addEdge("H", "B", 11)
	graph.addEdge("H", "G", 1)
	// graph.addEdge("H", "A", 8)

	graph.addEdge("C", "I", 2)
	graph.addEdge("C", "D", 7)
	graph.addEdge("C", "F", 14)

	graph.addEdge("D", "E", 9)
	graph.addEdge("D", "F", 14)
	graph.addEdge("D", "C", 7)

	graph.addEdge("E", "D", 9)
	graph.addEdge("E", "F", 10)

	graph.addEdge("F", "E", 10)
	graph.addEdge("F", "D", 14)

	graph.addEdge("G", "F", 2)
	graph.addEdge("G", "H", 1)
	graph.addEdge("G", "I", 6)

	graph.addEdge("I", "G", 6)
	graph.addEdge("I", "H", 7)
	graph.addEdge("I", "C", 2)

	fmt.Println(graph.getPath("A", "I"))
}
