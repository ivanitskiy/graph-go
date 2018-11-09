package graph

// ID is a unique ID of the Vertex/Node in the graph
type ID int

// Graph representation using an adjancency map
type Graph struct {
	vertices map[ID]struct{}
	edges    map[ID]map[ID]struct{}
	directed bool
}

// VertexCount returns the number of vertices in the graph
func (g *Graph) VertexCount() {

}

// Vertices returns an iteration of vertixes in the graph
func (g *Graph) Vertices() {

}

// EdgeCount returns the number of edges in the graph
func (g *Graph) EdgeCount() {

}

// Edges returns a set of all the edges in the graph
func (g *Graph) Edges() {

}

// GetEdge returns the edge from u to v, or nil if not adjacent
func (g *Graph) GetEdge(u, v ID) {

}

// Degre returns the number of (outgoing) edges incident to vertex v in the graph
func (g *Graph) Degre(v ID) {

}

// IncidentEdges return all (outgoing) edges incident to vertex v in the graph
func (g *Graph) IncidentEdges(v ID) {

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
