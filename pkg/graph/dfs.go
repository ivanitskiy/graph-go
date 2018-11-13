package graph

// DFS performs Depth first search of the undiscovered portion of Graph g starting at Vertex u
func DFS(g *Graph, start ID) map[ID]bool {
	visited := make(map[ID]bool)
	dfsRecursive(g, start, visited)
	return visited
}

func dfsRecursive(g *Graph, start ID, visited map[ID]bool) {
	visited[start] = true
	for _, id := range g.IncidentEdges(start) {
		if _, ok := visited[id.To]; !ok {
			dfsRecursive(g, id.To, visited)
		}
	}
}
