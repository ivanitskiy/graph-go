package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ivanitskiy/graph-go/pkg/graph"
)

var (
	in       io.Reader = os.Stdin
	out      io.Writer = os.Stdout
	filePath string
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func newDirectedGraph() *graph.Graph {
	defer timeTrack(time.Now(), "Reading DirectedGraph")
	g := graph.NewGraph(true)

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
		g.InsertEdge(graph.ID(n0), graph.ID(n1))
	}
	return g
}

func computeScc() {
	defer timeTrack(time.Now(), "computeSccGrapMethod()")
	ng := newDirectedGraph()
	// Compute SCC
	ngScc := make(map[graph.ID]int)
	css := ng.SCC()
	for _, lead := range *css {
		ngScc[lead]++
	}

	sccList := make([]int, 0)
	for _, v := range ngScc {
		sccList = append(sccList, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sccList)))
	fmt.Println(sccList[:5])
}

func main() {
	flag.StringVar(&filePath, "f", "", "Path to a file. If not provided then from stdin")
	flag.Parse()
	computeScc()
}
