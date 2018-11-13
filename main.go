package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ivanitskiy/graph-go/pkg/stack"
)

var (
	in       io.Reader = os.Stdin
	out      io.Writer = os.Stdout
	filePath string
	leader   map[int]int
	s        int
	visited  map[int]bool
	dfsStack stack.Stack
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func readGraph() *graphAdjacencyList {
	g := &graphAdjacencyList{
		nodes:        make(map[int]struct{}),
		edges:        make(map[int][]int),
		reverseEdges: make(map[int][]int),
	}
	var scanner *bufio.Scanner
	if filePath != "" {
		f, err := os.Open(filePath)

		checkError(err)
		defer f.Close()
		in = bufio.NewReader(f)
	}
	scanner = bufio.NewScanner(in)

	for scanner.Scan() {
		line := scanner.Text()
		sArray := strings.Split(line, " ")
		n0, err := strconv.Atoi(sArray[0])
		checkError(err)
		n1, err := strconv.Atoi(sArray[1])
		checkError(err)
		g.InsertEdge(n0, n1)
	}

	return g
}

type graphAdjacencyList struct {
	nodes        map[int]struct{}
	edges        map[int][]int
	reverseEdges map[int][]int
}

func (g *graphAdjacencyList) InsertEdge(u, v int) {
	if _, ok := g.nodes[u]; !ok {
		g.InsertNode(u)
	}
	if _, ok := g.nodes[v]; !ok {
		g.InsertNode(v)
	}

	if _, ok := g.edges[u]; !ok {
		g.edges[u] = make([]int, 0)
	}
	if _, ok := g.reverseEdges[v]; !ok {
		g.reverseEdges[v] = make([]int, 0)
	}
	g.edges[u] = append(g.edges[u], v)
	g.reverseEdges[v] = append(g.reverseEdges[v], u)
}

func (g *graphAdjacencyList) InsertNode(v int) bool {
	if _, ok := g.nodes[v]; ok {
		return false
	}
	g.nodes[v] = struct{}{}
	return true
}

func main() {
	flag.StringVar(&filePath, "f", "", "Path to a file. If not provided then from stdin")
	flag.Parse()
	g := readGraph()

	visited = make(map[int]bool)
	leader = make(map[int]int)

	order := make([]int, 0)

	for i := len(g.nodes); i > 0; i-- {
		order = append(order, i)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(order)))
	dfsLoop(g.reverseEdges, order)

	fv := make([]int, 0)

	for !dfsStack.IsEmpty() {
		value, _ := dfsStack.Pop()
		fv = append(fv, value)
	}
	dfsLoop(g.edges, fv)

	// Compute SCC
	scc := make(map[int]int)

	for _, lead := range leader {
		scc[lead]++
	}
	sccList := make([]int, 0)
	for _, v := range scc {
		sccList = append(sccList, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sccList)))
	//  Done. print result
	// fmt.Println(scc)
	fmt.Println(sccList[:5])
}

// DFS is dfs
func DFS(adjacencyList map[int][]int, start int) {
	visited[start] = true
	leader[start] = s
	for _, adj := range adjacencyList[start] {
		if _, ok := visited[adj]; !ok {
			DFS(adjacencyList, adj)
		}
	}
	dfsStack.Push(start)
}

func dfsLoop(aj map[int][]int, order []int) {
	s = 0
	visited = make(map[int]bool)
	for _, i := range order {
		if _, ok := visited[i]; !ok {
			s = i
			DFS(aj, i)
		}
	}
}
