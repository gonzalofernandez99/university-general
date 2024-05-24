package main

import "fmt"

// Definición de un grafo dirigido usando un map
type Graph struct {
	nodes map[string][]string
}

// Crear un nuevo grafo
func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]string)}
}

// Agregar una arista al grafo
func (g *Graph) AddEdge(from, to string) {
	g.nodes[from] = append(g.nodes[from], to)
}

// Imprimir el grafo
func (g *Graph) Print() {
	for from, toList := range g.nodes {
		fmt.Printf("%s -> %v\n", from, toList)
	}
}

func main() {
	g := NewGraph()
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "D")

	g.Print()
}
