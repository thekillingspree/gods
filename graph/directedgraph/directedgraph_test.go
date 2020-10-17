package directedgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertEdge(t *testing.T) {
	graph := InitGraph()
	normalCase := []Item{2, 3, 4, 5}
	graph.InsertEdge(1, 2, 3, 4, 5)
	assert.Equal(t, normalCase, graph.nodes[1], "they should be equal")

	noOutgoing := []Item{}
	graph.InsertEdge(2)
	assert.Equal(t, noOutgoing, graph.nodes[2], "they should be equal")
}

func TestKahnsTopologicalSort(t *testing.T) {

	// Test 1 - Empty Graph
	graph := InitGraph()
	topologicalOrder, err := graph.KahnsTopologicalSort()
	assert.Nil(t, err, "there should be no error")
	assert.Empty(t, topologicalOrder, "topological order should be empty")

	//Test 2 - Single Node.
	graph.InsertEdge(1)
	topologicalOrder, err = graph.KahnsTopologicalSort()
	assert.Nil(t, err, "there should be no error")
	assert.NotEmpty(t, topologicalOrder, "topological order should be empty")

	// Test 3 - DAG
	graph.InsertEdge(5, 2, 0)
	graph.InsertEdge(2, 3)
	graph.InsertEdge(3, 1)
	graph.InsertEdge(4, 1, 0)
	graph.InsertEdge(0)
	topologicalOrder, err = graph.KahnsTopologicalSort()
	assert.Nil(t, err, "there should be no error")
	assert.NotEmpty(t, topologicalOrder, "topological order should not be empty")

	// Test 4 - Graph with Cycle
	graph.InsertEdge(1, 2)
	topologicalOrder, err = graph.KahnsTopologicalSort()
	assert.NotNil(t, err, "there should be an error")
	assert.Empty(t, topologicalOrder, "topological order should be empty")
}
