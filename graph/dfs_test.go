package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDFS(t *testing.T) {

	// 1 5
	// 5 6
	// 6 9
	// 9 7
	// 7 8
	// 8 9
	// 4 5
	// 4 2
	// 2 3
	// 3 4

	g := NewGraph(true)
	g.InsertEdge(ID(1), ID(5))
	g.InsertEdge(ID(5), ID(6))
	g.InsertEdge(ID(6), ID(9))
	g.InsertEdge(ID(9), ID(7))
	g.InsertEdge(ID(7), ID(8))
	g.InsertEdge(ID(8), ID(9))
	g.InsertEdge(ID(4), ID(5))
	g.InsertEdge(ID(4), ID(2))
	g.InsertEdge(ID(2), ID(3))
	g.InsertEdge(ID(3), ID(4))

	t.Run("DFS", func(t *testing.T) {

		visited := DFS(g, 1)
		assert.Len(t, visited, 6)
	})

}
