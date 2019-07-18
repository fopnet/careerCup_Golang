package graph

import (
	"fmt"
	"testing"
)

/***
Graph image
https://www.geeksforgeeks.org/dijkstras-shortest-path-algorithm-in-java-using-priorityqueue/
https://www.geeksforgeeks.org/java-program-for-dijkstras-shortest-path-algorithm-greedy-algo-7/

**/

func TestUndirectedAlfaNumGraph(t *testing.T) {
	fmt.Println("Dijkstra")
	// Example
	graph := NewUndirectedGraph(7)
	graph.AddEdge("S", "B", 4)
	graph.AddEdge("S", "C", 2)
	graph.AddEdge("B", "C", 1)
	graph.AddEdge("B", "D", 5)
	graph.AddEdge("C", "D", 8)
	graph.AddEdge("C", "E", 10)
	graph.AddEdge("D", "E", 2)
	graph.AddEdge("D", "T", 6)
	graph.AddEdge("E", "T", 2)
	fmt.Println(graph.GetShortestPath("S", "T"))
}

func TestPriorityQueue(t *testing.T) {
	// Example
	graph := NewDirectedGraph(9)
	// graph := graph.NewUndirectedGraph()

	graph.AddEdge("A", "B", 4)
	graph.AddEdge("A", "H", 8)

	graph.AddEdge("B", "H", 11)
	graph.AddEdge("B", "C", 8)

	graph.AddEdge("H", "I", 7)
	graph.AddEdge("H", "B", 11)
	graph.AddEdge("H", "G", 1)
	graph.AddEdge("H", "A", 8)

	graph.AddEdge("C", "I", 2)
	graph.AddEdge("C", "D", 7)
	graph.AddEdge("C", "F", 4)

	graph.AddEdge("D", "E", 9)
	graph.AddEdge("D", "F", 14)
	// graph.AddEdge("D", "C", 7)

	graph.AddEdge("E", "D", 9)
	graph.AddEdge("E", "F", 10)

	graph.AddEdge("F", "E", 10)
	graph.AddEdge("F", "D", 14)

	graph.AddEdge("G", "F", 2)
	graph.AddEdge("G", "H", 1)

	graph.AddEdge("I", "G", 6)
	graph.AddEdge("I", "H", 7)
	graph.AddEdge("I", "C", 2)

	minWeight, minPath, distances := graph.GetShortestPath("A", "I")
	printPaths(distances, "A")
	fmt.Println(minWeight, minPath)
}

func printPaths(paths map[string]int, source string) {
	const A = int('A')
	// fmt.Println("len(paths)", paths)

	for i := A; i < A+len(paths); i++ {
		nodeZeroBased := i - 65
		node := string(rune(i))
		v := paths[node]
		fmt.Printf("%d to %d is %d \n", 0, nodeZeroBased, v)
	}
}
