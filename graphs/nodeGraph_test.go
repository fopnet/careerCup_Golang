package graph

import "testing"

var g IntGraph

func fillGraph() {
	nA := Node{1, 1}
	nB := Node{2, 1}
	nC := Node{3, 2}
	nD := Node{4, 2}
	nE := Node{5, 3}
	nF := Node{6, 4}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)
}

func TestAdd(t *testing.T) {
	fillGraph()
	g.String()
}
