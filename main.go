package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ivanitskiy/graph-go/graph"
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

func newGraph() *graph.Graph {

	digraph := graph.NewGraph(true)
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
		fmt.Println("Node is", graph.ID(n0), graph.ID(n1))
		digraph.InsertEdge(graph.ID(n0), graph.ID(n1))
		// back to string
		// s := strconv.Itoa(-42)
	}
	return digraph
}

func main() {
	flag.StringVar(&filePath, "f", "", "Path to a file. If not provided then from stdin")
	flag.Parse()

	g := newGraph()
	fmt.Println(g)
}
