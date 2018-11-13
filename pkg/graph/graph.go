package graph

import (
	"sort"
)

// ID is a unique ID of the Vertex/Node in the graph
type ID int

// byID implements sort.Interface for []ID
type byID []ID

func (a byID) Len() int           { return len(a) }
func (a byID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byID) Less(i, j int) bool { return a[i] < a[j] }

// AdjacencyList Graph representation
type AdjacencyList [][]Edge

// Graph representation using an adjancency map
// no parallel edges supported
type Graph struct {
	vertices map[ID]struct{}
	edges    map[ID]map[ID]struct{}
	directed bool
}

// DirectedGraph is a directed graph
type DirectedGraph struct {
	Graph
}

// UndirectedGraph is a directed graph
type UndirectedGraph struct {
	Graph
}

// VertexCount returns the number of vertices in the graph
func (g *Graph) VertexCount() int {
	return len(g.vertices)
}

// Vertices returns a slice of vertixes in the graph
func (g *Graph) Vertices() []ID {
	nodes := make([]ID, 0)
	for k := range g.vertices {
		nodes = append(nodes, k)
	}
	sort.Sort(byID(nodes))
	return nodes
}

// EdgeCount returns the number of edges in the graph
func (g *Graph) EdgeCount() int {
	return len(g.Edges())
}

// Edges returns a set of all the edges in the graph
func (g *Graph) Edges() []Edge {
	edges := make([]Edge, 0)
	for from, v := range g.edges {
		for to := range v {
			e := &Edge{
				From: from, To: to}
			edges = append(edges, *e)
		}
	}
	return edges
}

// GetEdge returns the edge from u to v, or nil if not adjacent
func (g *Graph) GetEdge(u, v ID) {

}

// Degre returns the number of (outgoing) edges incident to vertex v in the graph
func (g *Graph) Degre(v ID) int {
	return len(g.edges[v])
}

// IncidentEdges return all (outgoing) edges incident to vertex v in the graph
func (g *Graph) IncidentEdges(v ID) []Edge {
	incidents := make([]Edge, 0)
	for k := range g.edges[v] {
		e := &Edge{
			From: v,
			To:   k,
		}
		incidents = append(incidents, *e)
	}
	return incidents

}

// InsertVertex inserts and returns a new Vertex with element x
func (g *Graph) InsertVertex(v ID) bool {
	if _, ok := g.vertices[v]; ok {
		return false
	}
	g.vertices[v] = struct{}{}
	return true
}

// RemoveVertex removes the Vertex from the graph.
func (g *Graph) RemoveVertex(v ID) {
	// FIXME: get IncidentEdges() to remove them
	delete(g.vertices, v)

}

// InsertEdge inserts and returns a new Edge from u to v with auxilary element x
func (g *Graph) InsertEdge(u, v ID) {
	if _, ok := g.vertices[u]; !ok {
		g.InsertVertex(u)
	}
	if _, ok := g.vertices[v]; !ok {
		g.InsertVertex(v)
	}
	if _, ok := g.edges[u]; !ok {
		g.edges[u] = make(map[ID]struct{})
	}
	g.edges[u][v] = struct{}{}

	if !g.directed {
		// For undirected graph add edge from v to u.
		if _, ok := g.edges[v]; !ok {
			g.edges[v] = make(map[ID]struct{})
		}
		g.edges[v][u] = struct{}{}
	}

}

// RemoveEdge removes the eedge from the Graph
func (g *Graph) RemoveEdge(u, v ID) {
	delete(g.edges[u], v)
	if !g.directed {
		delete(g.edges[v], u)
	}
}

// Reverse reverts the graph
func (g *Graph) Reverse() {
	reversedEdges := make(map[ID]map[ID]struct{})
	for from := range g.edges {
		for to := range g.edges[from] {

			if _, ok := reversedEdges[to]; !ok {
				reversedEdges[to] = make(map[ID]struct{})
			}
			reversedEdges[to][from] = struct{}{}

			if !g.directed {
				// For undirected graph add edge from v to u.
				if _, ok := reversedEdges[from]; !ok {
					reversedEdges[from] = make(map[ID]struct{})
				}
				reversedEdges[from][to] = struct{}{}
			}
		}
	}
	g.edges = reversedEdges
}

// // TopologicalSort returns a list  of vertices of dicrected acyclic graph g in topological order
// func TopologicalSort(g Graph) {

// }

// NewGraph create a new Graph
func NewGraph(directed bool) *Graph {
	return &Graph{
		vertices: make(map[ID]struct{}),
		edges:    make(map[ID]map[ID]struct{}),
		directed: directed,
	}
}
