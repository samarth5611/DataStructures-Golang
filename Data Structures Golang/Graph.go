package main

import "fmt"

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Key  int
	List []*Vertex
}

func BuildGraph() Graph {
	newGraph := Graph{Vertices: make([]*Vertex, 0)}
	return newGraph
}

// Add Vertex
func (graph *Graph) AddVertex(v int) {
	if graph.Contains(v) {
		fmt.Println("Cannot Insert Vertex, It's already present!")
	}
	graph.Vertices = append(graph.Vertices, &Vertex{Key: v})
}

// Add Edge
func (graph *Graph) AddEdge(from int, to int) {
	fromVertex := graph.getVertex(from)
	toVertex := graph.getVertex(to)

	if fromVertex == nil || toVertex == nil{
		fmt.Println("Invalid Edge ", from, "-->", to)
		return
	}

	for _, i := range(fromVertex.List){
		if i.Key == toVertex.Key{
			fmt.Println("Edge Already exists ", from, "-->", to)
		}
	}

	fromVertex.List = append(fromVertex.List, toVertex)
	toVertex.List = append(toVertex.List, fromVertex)
}

// Contains
func (graph *Graph) Contains(vertex int) bool {
	for _, v := range graph.Vertices {
		if v.Key == vertex {
			return true
		}

	}
	return false
}

// getVertexIndex
func (graph *Graph) getVertex(vertex int) *Vertex {
	for _, v := range graph.Vertices {
		if v.Key == vertex {
			return v
		}
	}
	return nil
}

// Print Graph
func (graph *Graph) Print() {
	for _, v := range graph.Vertices {
		fmt.Printf("Vertex %v :", v.Key)
		for _, num := range v.List {
			fmt.Print(num.Key, " ")
		}
		fmt.Println("")
	}
}

func main() {
	myGraph := BuildGraph()

	for i := 1; i <= 5; i++ {
		myGraph.AddVertex(i)
	}
	myGraph.AddEdge(1, 4)
	myGraph.AddEdge(2, 5)
	myGraph.AddEdge(2, 5)
	myGraph.AddEdge(3, 4)
	myGraph.AddEdge(5, 3)
	myGraph.AddEdge(5, 9)
	myGraph.Print()
}
