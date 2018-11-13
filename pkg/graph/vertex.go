package graph

// Vertex is a Node/Vettex in the graph
type Vertex map[ID]struct {
	visited bool
	scc     ID
}
