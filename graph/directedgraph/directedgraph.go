package directedgraph

import (
	"fmt"

	"github.com/thekillingspree/gods/queue"
)

// Item represents an item in the graph.
type Item interface{}

// DirectedGraph struct uses adjacency list to represent a Directed Graph
type DirectedGraph struct {
	nodes map[Item][]Item
}

//InsertEdge will create an edge.
func (g *DirectedGraph) InsertEdge(a Item, b ...Item) {
	if _, ok := g.nodes[a]; ok {
		g.nodes[a] = append(g.nodes[a], b...)
	} else if len(b) > 0 {
		g.nodes[a] = b
	} else {
		g.nodes[a] = []Item{}
	}
}

//InitGraph creates a Graph struct and returns the pointer
func InitGraph() *DirectedGraph {
	graph := &DirectedGraph{nodes: make(map[Item][]Item)}
	return graph
}

//KahnsTopologicalSort returns the Topological sorting of nodes with Kahns Algorithm
func (g DirectedGraph) KahnsTopologicalSort() ([]Item, error) {
	nodesWithoutIncomingEdges := queue.Queue{}
	indegrees := make(map[Item]int)
	topologicalOrdering := []Item{}

	// Initialize indegrees
	for node := range g.nodes {
		indegrees[node] = 0
	}

	// update Indegrees
	for _, neighbours := range g.nodes {
		for _, node := range neighbours {
			if pop, ok := indegrees[node]; ok {
				indegrees[node] = pop + 1
			}
		}
	}

	//find nodes with no indegrees
	for node, indegree := range indegrees {
		if indegree == 0 {
			nodesWithoutIncomingEdges.Enqueue(node)
		}
	}
	// topological sort
	visited := 0
	for nodesWithoutIncomingEdges.Len() > 0 {
		currentNode, _ := nodesWithoutIncomingEdges.Dequeue()
		visited++
		topologicalOrdering = append(topologicalOrdering, currentNode)
		for _, neighbour := range g.nodes[currentNode] {
			indegrees[neighbour]--
			if indegrees[neighbour] == 0 {
				nodesWithoutIncomingEdges.Enqueue(neighbour)
			}
		}
	}

	if visited != len(g.nodes) {
		return nil, fmt.Errorf("cycle Exists in graph")
	}

	return topologicalOrdering, nil
}
