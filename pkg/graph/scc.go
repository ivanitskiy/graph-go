package graph

import (
	"sort"

	"github.com/ivanitskiy/graph-go/pkg/stack"
)

type visitedMap map[ID]bool

func makeDecrisingOrder(g *Graph) []ID {
	order := make([]ID, g.VertexCount())
	i := 0
	for node := range g.vertices {
		order[i] = node
		i++
	}
	sort.Slice(order, func(i, j int) bool {
		return order[i] > order[j]
	})
	return order
}

func drainStack(dfsStack *stack.Stack) []ID {
	fv := make([]ID, len(*dfsStack))
	i := 0
	for !dfsStack.IsEmpty() {
		value, _ := dfsStack.Pop()
		fv[i] = ID(value)
		i++
	}
	return fv
}

// SCC computes strongly connected components
func (g *Graph) SCC() *map[ID]ID {
	// init values

	// dfsStack := make(stack.Stack, 0)
	var dfsStack stack.Stack

	// invoke DFS against Reversed graph on nodes orderd by ID n to 1
	order := makeDecrisingOrder(g)
	dfsSCCLoop(g.Reverse(), &dfsStack, order)

	order = drainStack(&dfsStack)
	leader := dfsSCCLoop(g.edges, &dfsStack, order)
	return leader
}

func dfsSCCLoop(aj AdjacencyList, dfsStack *stack.Stack, order []ID) *map[ID]ID {
	leader := make(map[ID]ID)
	var leaderNode ID
	visited := make(visitedMap)
	for _, source := range order {
		if _, ok := visited[source]; !ok {
			leaderNode = source
			dfsSCC(
				aj,
				source,
				visited,
				leader,
				leaderNode,
				dfsStack)
		}
	}
	return &leader
}

// dfs is depth first search on graph.
// sets
func dfsSCC(aj AdjacencyList, start ID, visited visitedMap, leader map[ID]ID, leaderNode ID, dfsStack *stack.Stack) {
	visited[start] = true
	leader[start] = leaderNode
	for to := range aj[start] {
		if _, ok := visited[to]; !ok {
			dfsSCC(
				aj,
				to,
				visited,
				leader,
				leaderNode,
				dfsStack)
		}
	}
	dfsStack.Push(int(start))
}
