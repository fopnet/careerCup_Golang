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
L - Line
C - Column

L->C = Horizontal
C->C = Horizontal
L->L = Vertical
C->L = Vertical
*/

func TestMazeMinRoute(t *testing.T) {
	// Example
	graph := graph.NewDirectedGraph(8)
	/*
		graph.AddEdge("A", "F", 1)
		graph.AddEdge("A", "G", 2)
		graph.AddEdge("A", "H", 3)

		graph.AddEdge("F", "B", 1)
		graph.AddEdge("F", "C", 2)

		graph.AddEdge("C", "E", 1)
		graph.AddEdge("C", "G", 1)
		graph.AddEdge("C", "H", 2)

		graph.AddEdge("E", "D", 1)
	*/
	graph.AddEdge("A", "F", 1)
	graph.AddEdge("A", "G", 1)
	graph.AddEdge("A", "H", 1)

	graph.AddEdge("F", "B", 1)
	graph.AddEdge("F", "C", 1)

	graph.AddEdge("C", "E", 1)
	graph.AddEdge("C", "G", 1)
	graph.AddEdge("C", "H", 1)

	graph.AddEdge("E", "D", 1)

	fmt.Println(graph.GetShortestPath("A", "D"))
}

func TestGetWeight(t *testing.T) {

	l, c := 0, 3
	w := getWeight(l, c)

	expected := 6
	if expected != w {
		t.Errorf("The result %d should be %d ", w, expected)
	}

	l, c = 1, 1
	expected = 3
	w = getWeight(l, c)
	if expected != w {
		t.Errorf("The result %d should be %d ", w, expected)
	}
}

func TestMatrix2Graph(t *testing.T) {
	matrix := [][]rune{
		{'O', 'O', 'O', 'O'},
		{'D', 'O', 'D', 'O'},
		{'O', 'O', 'O', 'O'},
		{'X', 'D', 'D', 'O'},
	}

	destiny, graph := convertMatrixToGraph(matrix)

	fmt.Printf("destiny %s \n graph=>%v\n", destiny, graph)
}

func convertMatrixToGraph(m [][]rune) (string, graph.Graph) {

	var destiny string

	n := len(m)
	graph := graph.NewDirectedGraph(n * n)

	const A = int('A')
	const E = int('E')

	for l, _ := range m {
		for c, _ := range m {
			v := (m[l][c])
			// v := string(matrix[l][c])
			var node string
			var edge string

			switch v {
			case 'X':
				destiny = getNode(l, c)
			case 'O':
				// Ci->Cj check horizontal right moving
				if c < n-1 && canMove(l, c+1, m) {
					node = getNode(l, c)
					edge = getNode(l, c+1)
					// graph.AddEdge(node, edge, c+1+l)
					graph.AddEdge(node, edge, 1)
				}
				// Cj->Ci check horizontal left moving
				if c > 0 && canMove(l, c-1, m) {
					node = getNode(l, c)
					edge = getNode(l, c-1)
					// graph.AddEdge(node, edge, c-1+l)
					graph.AddEdge(node, edge, 1)
				}
				// Ci->Lj+1 check down vertical moving
				if l < n-1 && canMove(l+1, c, m) {
					node = getNode(l, c)
					edge = getNode(l+1, c)
					// graph.AddEdge(node, edge, l+1+c)
					graph.AddEdge(node, edge, 1)
				}
				break
			case 'D':
			default:
			}
		}
	}

	return destiny, graph
}

func getWeight(l, c int) int {
	const lTarget, cTarget = 3, 0

	return (c - cTarget) + (lTarget - l)
}

func canMove(l, c int, m [][]rune) bool {
	return m[l][c] == 'O' || m[l][c] == 'X'
}

func getNode(l, c int) string {
	return fmt.Sprintf("%d,%d", l, c)
}

func TestMatrix2GraphPlusShortestPath(t *testing.T) {
	matrix := [][]rune{
		{'O', 'O', 'O', 'O'},
		{'D', 'O', 'D', 'O'},
		{'O', 'O', 'O', 'O'},
		{'X', 'D', 'D', 'O'},
	}

	destiny, graph := convertMatrixToGraph(matrix)

	w, path, mapPath := graph.GetShortestPath("0,0", destiny)
	fmt.Printf("%s -> %s has weight %d walking %v ::: %v \n", "0,0", destiny, w, path, mapPath)
}
