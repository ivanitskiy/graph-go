package graph

// import (
// 	"fmt"
// 	"sort"

// 	"github.com/ivanitskiy/graph-go/pkg/stack"
// )

// func SCC(g *Graph) {

// 	visited := make(map[int]bool)
// 	leader := make(map[int]int)
// 	dfsStack := make(stack.Stack, 0)
// 	var s = 0

// 	sCCForDFS := func(adjacencyList map[int][]int, start int) {
// 		visited[start] = true
// 		leader[start] = s
// 		for _, adj := range adjacencyList[start] {
// 			if _, ok := visited[adj]; !ok {
// 				sCCForDFS(adjacencyList, adj)
// 			}
// 		}
// 		dfsStack.Push(start)
// 	}
// 	order := make([]ID, len(g.vertices))

// 	for i := 1; i <= len(g.vertices)

// 	for i := len(g.vertices); i > 0; i-- {
// 		order[] = append(order, i)
// 	}
// 	sort.Sort(sort.Reverse(sort.IntSlice(order)))
// 	dfsLoop(ajList.reverseEdges, order)

// 	fv := make([]int, 0)

// 	for !dfsStack.IsEmpty() {
// 		value, _ := dfsStack.Pop()
// 		fv = append(fv, value)
// 	}
// 	dfsLoop(ajList.edges, fv)

// 	// Compute SCC
// 	scc := make(map[int]int)

// 	for _, lead := range leader {
// 		scc[lead]++
// 	}
// 	sccList := make([]int, 0)
// 	for _, v := range scc {
// 		sccList = append(sccList, v)
// 	}
// 	sort.Sort(sort.Reverse(sort.IntSlice(sccList)))
// 	//  Done. print result
// 	// fmt.Println(scc)
// 	fmt.Println(sccList[:5])

// }

// type graphAdjacencyList struct {
// 	nodes        map[int]struct{}
// 	edges        map[int][]int
// 	reverseEdges map[int][]int
// }

// func (g *graphAdjacencyList) InsertEdge(u, v int) {
// 	if _, ok := g.nodes[u]; !ok {
// 		g.InsertNode(u)
// 	}
// 	if _, ok := g.nodes[v]; !ok {
// 		g.InsertNode(v)
// 	}

// 	if _, ok := g.edges[u]; !ok {
// 		g.edges[u] = make([]int, 0)
// 	}
// 	if _, ok := g.reverseEdges[v]; !ok {
// 		g.reverseEdges[v] = make([]int, 0)
// 	}
// 	g.edges[u] = append(g.edges[u], v)
// 	g.reverseEdges[v] = append(g.reverseEdges[v], u)
// }

// func (g *graphAdjacencyList) InsertNode(v int) bool {
// 	if _, ok := g.nodes[v]; ok {
// 		return false
// 	}
// 	g.nodes[v] = struct{}{}
// 	return true
// }

// // DFS is dfs
// func sCCForDFS(adjacencyList map[int][]int, start int) {
// 	visited[start] = true
// 	leader[start] = s
// 	for _, adj := range adjacencyList[start] {
// 		if _, ok := visited[adj]; !ok {
// 			sCCForDFS(adjacencyList, adj)
// 		}
// 	}
// 	dfsStack.Push(start)
// }

// func dfsLoop(aj map[int][]int, order []int) {
// 	s = 0
// 	visited = make(map[int]bool)
// 	for _, i := range order {
// 		if _, ok := visited[i]; !ok {
// 			s = i
// 			sCCForDFS(aj, i)
// 		}
// 	}
// }
