package shorterPath

import (
	"fmt"
	"testing"

	"github.com/fopnet/carrerCup_Golang/graphs/graph"
)

/*
https://www.careercup.com/question?id=6215770032308224
A matrix represents a sequence of travel points. One only can travel either left/right or up/down.
Some of those points are dead points which one can't travel any further.
There is a destination point in the matrix.
Find the shortest path from the top left point (1, 1) to the destination.

e.g.
Input
[	  E    F    G    H
A	[‘O’, ‘O’, ‘O’, ‘O’],
B	[‘D’, ‘O’, ‘D’, ‘O’],
C	[‘O’, ‘O’, ‘O’, ‘O’],
D	[‘X’, ‘D’, ‘D’, ‘O’],
]

Output
Route is (0, 0), (0, 1), (1, 1), (2, 1), (2, 0), (3, 0) The minimum route takes 5 steps.

Tips
Horizontal = Line,Column
Vertical = Column,Line
*/

func TestMazeMinRoute(t *testing.T) {
	// Example
	graph := graph.NewDirectedGraph(8)
	// graph := graph.NewUndirectedGraph()
	graph.AddEdge("A", "F", 1)
	graph.AddEdge("A", "G", 2)
	graph.AddEdge("A", "H", 3)

	graph.AddEdge("F", "B", 1)
	graph.AddEdge("F", "C", 2)

	graph.AddEdge("C", "E", 1)
	graph.AddEdge("C", "G", 1)
	graph.AddEdge("C", "H", 2)

	graph.AddEdge("E", "D", 1)

	fmt.Println(graph.GetShortestPath("A", "D"))
}


