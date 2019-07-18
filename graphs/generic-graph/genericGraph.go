package genericGraph

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

/*
generic type https://github.com/cheekybits/genny

genny -in genericGraph.go -out gen-graph.go gen "Item=int,string"

in the command line now, if the file is called genericGraph.go it will generate a gen-graph.go

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Item=int"

run int the shell > go generate
*/
type Item generic.Type

// Node a single node that composes the tree
type Node struct {
	value  Item
	weight Item
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

// ItemGraph the Items graph
type ItemGraph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
	g.lock.Unlock()
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) String() {
	g.lock.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.RUnlock()
}
