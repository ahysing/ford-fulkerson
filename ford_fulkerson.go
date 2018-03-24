package main

import (
	"errors"
	"log"
	"os"
)

// Edge from vertex to vertex with a cost.
type Edge struct {
	u        string
	v        string
	capacity float32
	redge    *Edge
}

// Graph of vertecies with edges between them.
type Graph struct {
	adj  map[Edge]Edge
	flow map[Edge]float32
}

// Tuple is path and its residual flow
type Tuple struct {
	edge     Edge
	residual float32
}

// New creates an empty graph
func New() Graph {
	var g Graph
	g.adj = make(map[Edge]Edge)
	g.flow = make(map[Edge]float32)
	return g
}

func (g Graph) addEdge(u string, v string, capacity float32) error {
	if u == v {
		return errors.New("U == V")
	}

	edge := Edge{u, v, capacity, nil}
	redge := Edge{v, u, 0.0, &edge}
	edge.redge = &edge

	g.adj[edge] = redge
	g.adj[redge] = edge

	g.flow[edge] = 0.0
	g.flow[redge] = 0.0

	return nil
}

func (g Graph) findPath(source string, sink string, path []Tuple, pathSet map[Tuple]bool) []Tuple {
	if source == sink {
		return path
	}

	for edge := range g.getEdges(source) {
		residual := edge.capacity - g.flow[edge]
		tuple := Tuple{
			source,
			residual
		}

		_, hasTuple := pathSet
		if residual > 0 && !hasTuple {
			pathSet[Tuple] = true
			result := g.findPath(edge.sink, sink, pathSet)
			if result != nil {
				return result
			}
		}
	}

	return nil
}

func (g Graph) maxFlow(source string, sink string) {
	slice := make([]Tuple)
	start := make(map[Edge]float32)
	path := g.findPath(source, sink, slice, start)
	for path != nil {
		flow := 2147483647.0
		for _, res := range path {
			if res < flow {
				flow = res
			}
		}

		for edge, res := range path {
			g.flow[edge] += flow
			g.flow[edge.redge] -= flow
		}

		slice = make([]Tuple)
		start = make(map[Edge]float32)
		path = g.findPath(source, sink, slice, start)
	}

	var sum float32
	for edge := range g.getEdges(source) {
		sum += g.flow[edge]
	}

	return sum
}

func fordFulkerson() {
	var graph = New()
	verts := []string{"s", "o", "p", "q", "r", "t"}
	for vertex := range verts {
		g.addVertex(vertex)
	}

	g.addEdge("s", "o", 3)
	g.addEdge("s", "p", 3)
	g.addEdge("o", "p", 2)
	g.addEdge("o", "q", 3)
	g.addEdge("p", "r", 2)
	g.addEdge("r", "t", 3)
	g.addEdge("q", "r", 4)
	g.addEdge("q", "t", 2)

	result := g.maxFlow("s", "t")
	log.Println(result)
}

func main() {
	if len(os.Args) >= 2 {
		log.Printf("Error")
	} else {
		fordFulkerson()
	}
}
