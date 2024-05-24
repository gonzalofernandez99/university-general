package main

import "fmt"

// Definición de un grafo no dirigido usando un map
type Graph struct {
	nodes map[string][]string
}

// Crear un nuevo grafo
func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]string)}
}

// Agregar una arista al grafo no dirigido
func (g *Graph) AddEdge(node1, node2 string) {
	g.nodes[node1] = append(g.nodes[node1], node2)
	g.nodes[node2] = append(g.nodes[node2], node1)
}

// Imprimir el grafo
func (g *Graph) Print() {
	for node, adjList := range g.nodes {
		fmt.Printf("%s -> %v\n", node, adjList)
	}
}

func main() {
	g := NewGraph()
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "D")

	g.Print()
}
