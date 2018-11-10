package graph

// Edge of in the graph
// That is, a {"ID1": {ID2: {} }} is an edge from ID1 to ID2
//
// type Edge map[ID]map[ID]struct{}
type Edge struct {
	from ID
	to   ID
}
