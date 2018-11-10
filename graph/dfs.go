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
		if _, ok := visited[id]; !ok {
			dfsRecursive(g, id, visited)
		}
	}
}

// DFSComplete performs Depth first search of the entire Graph and returns forest as map
func DFSComplete(g *Graph) {

}
