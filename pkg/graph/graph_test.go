package graph

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph_VertexCount(t *testing.T) {
	type fields struct {
		vertices map[ID]struct{}
		edges    map[ID]map[ID]struct{}
		directed bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				vertices: tt.fields.vertices,
				edges:    tt.fields.edges,
				directed: tt.fields.directed,
			}
			g.VertexCount()
		})
	}
}

func TestGraph_Vertices(t *testing.T) {
	type fields struct {
		vertices map[ID]struct{}
		edges    map[ID]map[ID]struct{}
		directed bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				vertices: tt.fields.vertices,
				edges:    tt.fields.edges,
				directed: tt.fields.directed,
			}
			g.Vertices()
		})
	}
}

func TestGraph_EdgeCount(t *testing.T) {
	type fields struct {
		vertices map[ID]struct{}
		edges    map[ID]map[ID]struct{}
		directed bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				vertices: tt.fields.vertices,
				edges:    tt.fields.edges,
				directed: tt.fields.directed,
			}
			g.EdgeCount()
		})
	}
}

func TestGraph_Edges(t *testing.T) {
	type fields struct {
		vertices map[ID]struct{}
		edges    map[ID]map[ID]struct{}
		directed bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				vertices: tt.fields.vertices,
				edges:    tt.fields.edges,
				directed: tt.fields.directed,
			}
			g.Edges()
		})
	}
}

func TestGraph_GetEdge(t *testing.T) {
	type fields struct {
		vertices map[ID]struct{}
		edges    map[ID]map[ID]struct{}
		directed bool
	}
	type args struct {
		u ID
		v ID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				vertices: tt.fields.vertices,
				edges:    tt.fields.edges,
				directed: tt.fields.directed,
			}
			g.GetEdge(tt.args.u, tt.args.v)
		})
	}
}

func TestGraph_Degre(t *testing.T) {
	type fields struct {
		vertices map[ID]struct{}
		edges    map[ID]map[ID]struct{}
		directed bool
	}
	type args struct {
		v ID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				vertices: tt.fields.vertices,
				edges:    tt.fields.edges,
				directed: tt.fields.directed,
			}
			g.Degre(tt.args.v)
		})
	}
}

func TestGraph_IncidentEdges(t *testing.T) {
	type fields struct {
		vertices map[ID]struct{}
		edges    map[ID]map[ID]struct{}
		directed bool
	}
	type args struct {
		v ID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				vertices: tt.fields.vertices,
				edges:    tt.fields.edges,
				directed: tt.fields.directed,
			}
			g.IncidentEdges(tt.args.v)
		})
	}
}

func TestGraph_InsertVertex(t *testing.T) {
	type args struct {
		v        ID
		directed bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Assert vertices inserted for directed graph",
			args: args{
				v:        ID(2),
				directed: true,
			},
		},
		{
			name: "Assert vertices inserted for un directed graph",
			args: args{
				v:        ID(2),
				directed: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.args.directed)
			ok := g.InsertVertex(tt.args.v)
			assert.True(t, ok, "v not added")
			ok = g.InsertVertex(tt.args.v)
			assert.False(t, ok, "v not found in the vertices")

		})
	}
}

func TestGraph_RemoveVertex(t *testing.T) {
	type fields struct {
		vertices map[ID]struct{}
		edges    map[ID]map[ID]struct{}
		directed bool
	}
	type args struct {
		v ID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				vertices: tt.fields.vertices,
				edges:    tt.fields.edges,
				directed: tt.fields.directed,
			}
			g.RemoveVertex(tt.args.v)
		})
	}
}

func TestGraph_InsertEdge(t *testing.T) {

	type args struct {
		u        ID
		v        ID
		directed bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Assert vertices created when edge inserted for directed graph",
			args: args{
				u:        ID(1),
				v:        ID(2),
				directed: true,
			},
		},
		{
			name: "Assert vertices created when edge inserted for un directed graph",
			args: args{
				u:        ID(1),
				v:        ID(2),
				directed: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.args.directed)
			g.InsertEdge(tt.args.u, tt.args.v)
			_, ok := g.vertices[tt.args.u]
			assert.True(t, ok, "u not found in the vertices")
			_, ok = g.vertices[tt.args.v]
			assert.True(t, ok, "v not found in the vertices")

		})
	}
}

func TestGraph_RemoveEdge(t *testing.T) {
	type fields struct {
		vertices map[ID]struct{}
		edges    map[ID]map[ID]struct{}
		directed bool
	}
	type args struct {
		u ID
		v ID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				vertices: tt.fields.vertices,
				edges:    tt.fields.edges,
				directed: tt.fields.directed,
			}
			g.RemoveEdge(tt.args.u, tt.args.v)
		})
	}
}

func TestNewGraph(t *testing.T) {
	type args struct {
		directed bool
	}
	tests := []struct {
		name string
		args args
		want *Graph
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGraph(tt.args.directed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_Reverse(t *testing.T) {
	t.Run("reverse the graph", func(t *testing.T) {

		g := NewGraph(true)
		id1 := ID(1)
		id2 := ID(2)
		id3 := ID(3)
		id4 := ID(4)

		g.InsertEdge(id1, id2)
		g.InsertEdge(id1, id4)
		g.InsertEdge(id2, id3)
		g.InsertEdge(id2, id4)
		g.InsertEdge(id3, id1)
		t.Log(g)

		assert.Len(t, g.IncidentEdges(id1), 2, "Incorrect # of edges. Should be 2")
		g.Reverse()
		assert.Len(t, g.IncidentEdges(id1), 1, "u not found in the vertices")
		assert.Equal(t, g.IncidentEdges(id1)[0].From, id1, "u not found in the vertices")
		assert.Equal(t, g.IncidentEdges(id1)[0].To, id3, "u not found in the vertices")
		t.Log(g)
		g.Reverse()
		t.Log(g)
	})
}
