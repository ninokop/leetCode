package graph

import (
	"log"
	"testing"
)

type node struct{ id string }

func (n node) ID() string { return n.id }
func (n node) Run()       { log.Println(n.id) }

func TestLinkerGraphDFS(t *testing.T) {
	log.Printf("DFS begin")
	cases := []string{"A", "B", "C", "D", "E", "F"}
	edges := [][]string{
		{"A", "B"},
		{"A", "C"},
		{"A", "F"},
		{"C", "B"},
		{"C", "D"},
		{"D", "B"},
		{"D", "E"},
	}

	g := newLinkerGraph(len(cases))
	for i := range cases {
		g.AddNode(node{cases[i]})
	}
	for i := range edges {
		g.AddEdge(edges[i][0], edges[i][1])
	}
	g.DFS()
}

func TestLinkerGraphBFS(t *testing.T) {
	log.Printf("BFS begin")
	cases := []string{"A", "B", "C", "D", "E", "F"}
	edges := [][]string{
		{"A", "B"},
		{"A", "C"},
		{"A", "F"},
		{"C", "B"},
		{"C", "D"},
		{"D", "B"},
		{"D", "E"},
	}

	g := newLinkerGraph(len(cases))
	for i := range cases {
		g.AddNode(node{cases[i]})
	}
	for i := range edges {
		g.AddEdge(edges[i][0], edges[i][1])
	}
	g.BFS()
}

/*go test --bench=. --run=none
goos: linux
goarch: amd64
BenchmarkNewListGraph-4           500000              4015 ns/op            1214 B/op         24 allocs/op
BenchmarkListGraphBFS-4           200000              8066 ns/op            1744 B/op         26 allocs/op
BenchmarkListGraphDFS-4           200000              7117 ns/op            1792 B/op         29 allocs/op
*/

type bnode struct{ id string }

func (n bnode) ID() string { return n.id }
func (n bnode) Run()       {}

func BenchmarkNewListGraph(b *testing.B) {
	cases := []string{"A", "B", "C", "D", "E", "F"}
	edges := [][]string{
		{"A", "B"},
		{"A", "C"},
		{"A", "F"},
		{"C", "B"},
		{"C", "D"},
		{"D", "B"},
		{"D", "E"},
	}
	g := newLinkerGraph(len(cases))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := range cases {
			g.AddNode(bnode{cases[i]})
		}
		for i := range edges {
			g.AddEdge(edges[i][0], edges[i][1])
		}
	}
	b.ReportAllocs()
}

func BenchmarkListGraphBFS(b *testing.B) {
	cases := []string{"A", "B", "C", "D", "E", "F"}
	edges := [][]string{
		{"A", "B"},
		{"A", "C"},
		{"A", "F"},
		{"C", "B"},
		{"C", "D"},
		{"D", "B"},
		{"D", "E"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := newLinkerGraph(len(cases))
		for i := range cases {
			g.AddNode(bnode{cases[i]})
		}
		for i := range edges {
			g.AddEdge(edges[i][0], edges[i][1])
		}
		g.BFS()
	}
	b.ReportAllocs()
}

func BenchmarkListGraphDFS(b *testing.B) {
	cases := []string{"A", "B", "C", "D", "E", "F"}
	edges := [][]string{
		{"A", "B"},
		{"A", "C"},
		{"A", "F"},
		{"C", "B"},
		{"C", "D"},
		{"D", "B"},
		{"D", "E"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := newLinkerGraph(len(cases))
		for i := range cases {
			g.AddNode(bnode{cases[i]})
		}
		for i := range edges {
			g.AddEdge(edges[i][0], edges[i][1])
		}
		g.DFS()
	}
	b.ReportAllocs()
}
