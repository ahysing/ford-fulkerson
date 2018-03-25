package main

import (
	"errors"
	"log"
)

// Converted to GO from from http://www.cse.unt.edu/~tarau/teaching/AnAlgo/Ford–Fulkerson%20algorithm.pdf

// Edge from vertex to vertex with a cost.
type Edge struct {
	source   string
	sink     string
	capacity float32
	redge    *Edge
}

// Graph of vertecies with edges between them.
type Graph struct {
	adjacent map[string][]Edge
	flow     map[Edge]float32
}

// ConnectedEdge is path and its residual flow
type ConnectedEdge struct {
	edge     Edge
	residual float32
}

// New creates an empty Graph
// Graph is the top level data structure in Ford-Fulkerson
// A graph is a set of vertecies with edges between them.
// Edges are stored in `adjacent` while vertecies are defined as properties in every edge
//
// The graph also has a flow which states now much recidual capacity is allocated through that
// peticular edge.
// After Ford-Fulkerson has run over the all available capacity is allocated to edges.
func New() Graph {
	var g Graph
	g.adjacent = make(map[string][]Edge)
	g.flow = make(map[Edge]float32)
	return g
}

// AddEdge adds an edge to the graph from vertex to vertex.
// An edge is defined with to vertecies and a capacity between them
func (g Graph) AddEdge(u string, v string, capacity float32) error {
	if u == v {
		return errors.New("U == V")
	}

	edge := Edge{u, v, capacity, nil}
	redge := Edge{v, u, 0.0, nil}

	redge.redge = &edge
	edge.redge = &redge

	g.adjacent[u] = append(g.adjacent[u], edge)
	g.adjacent[v] = append(g.adjacent[v], redge)

	g.flow[edge] = 0.0
	g.flow[redge] = 0.0

	return nil
}

// AddVertex adds a named node in the directed graph of vertecies
// A vertex is described by its unique name
func (g Graph) AddVertex(v string) {
	g.adjacent[v] = make([]Edge, 1)
}
func (g Graph) getEdges(v string) []Edge {
	return g.adjacent[v]
}

func (g Graph) findPath(source string, sink string, path []ConnectedEdge, pathSet map[ConnectedEdge]bool) []ConnectedEdge {
	if source == sink {
		return path
	}

	edges := g.getEdges(source)
	for _, edge := range edges {
		residual := edge.capacity - g.flow[edge]
		connectedEdge := ConnectedEdge{edge, residual}

		_, hasConnectedEdge := pathSet[connectedEdge]
		if residual > 0.0 && !hasConnectedEdge {
			pathSet[connectedEdge] = true
			nextPath := append(path, connectedEdge)
			finalPath := g.findPath(edge.sink, sink, nextPath, pathSet)
			if finalPath != nil && len(finalPath) > 0 { // TODO: light checking
				return finalPath
			}
		}
	}

	return nil
}

// FordFulkerson traverses the graph according to the algorithm and assigns the maximum flow in the graph.
// The maximum flow lives inside property `flow`
func FordFulkerson(g Graph, source string, sink string) Graph {
	newPath := make([]ConnectedEdge, 0)
	newPathSet := make(map[ConnectedEdge]bool)

	path := g.findPath(source, sink, newPath, newPathSet)
	for path != nil {
		flow := float32(2147483647.0)
		for _, p := range path {
			if p.residual > 0.0 && p.residual < flow {
				flow = p.residual
			}
		}

		for _, p := range path {
			edge := p.edge

			g.flow[edge] = g.flow[edge] + flow
			g.flow[*edge.redge] -= flow
			if edge == *edge.redge {
				panic("edge == reverse edge")
			}
		}

		newPath = make([]ConnectedEdge, 0)
		newPathSet = make(map[ConnectedEdge]bool)
		path = g.findPath(source, sink, newPath, newPathSet)
	}

	return g
}

// MaxFlow finds the maximum flow through a graph where Ford-Fulkerson has assigned the maximum flow
func (g Graph) MaxFlow(source string, sink string) float32 {
	var maxFlow float32

	for _, edge := range g.getEdges(source) {
		maxFlow += g.flow[edge]
	}

	return maxFlow
}

func buildExampleGraph() Graph {
	var g = New()
	verts := []string{"s", "o", "p", "q", "r", "t"}
	for _, vertex := range verts {
		g.AddVertex(vertex)
	}

	g.AddEdge("s", "o", 3)
	g.AddEdge("s", "p", 3)
	g.AddEdge("o", "p", 2)
	g.AddEdge("o", "q", 3)
	g.AddEdge("p", "r", 2)
	g.AddEdge("r", "t", 3)
	g.AddEdge("q", "r", 4)
	g.AddEdge("q", "t", 2)
	return g
}

func main() {
	g := buildExampleGraph()
	graphTraversed := FordFulkerson(g, "s", "t")
	result := graphTraversed.MaxFlow("s", "t")

	log.Println("== Printing example graph ==")
	for vertex, edges := range g.adjacent {
		log.Println("Vertex ", vertex)
		for _, edge := range edges {
			if edge.capacity > 0 {
				log.Printf("Connected to %v with capcaity %v", edge.sink, edge.capacity)
			}
		}
	}
	log.Println(" == Graph finished ==")

	log.Printf("Flow from s to t is %v", result)
}
