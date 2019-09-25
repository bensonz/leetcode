package main

import (
	"fmt"
)

// Vertex defines a vertex of our graph
type Vertex struct {
	id        int   // used to identify node
	relatives []int // contains the ids of its neighbors
	distances []int
}

type path struct {
	distance int
	to       Vertex
}

// each row contains a vertex and its path to nearest neighbors
type row map[int]path

// Table is a list of rows
type Table []row

// Graph a graph contains a list of all nodes
type Graph []Vertex

func (n *Vertex) isEqual(m Vertex) bool {
	return n.id == m.id
}

func (n *Vertex) getDistance(m Vertex) int {
	for i, relative := range n.relatives {
		if relative == m.id {
			return n.distances[i]
		}
	}
	// not in its nearby relatives
	return -1
}

func (g *Graph) contains(v Vertex) bool {
	for _, checkV := range *g {
		if checkV.isEqual(v) {
			return true
		}
	}
	return false
}

func (g *Graph) remove(v Vertex) bool {
	for i, checkV := range *g {
		if checkV.isEqual(v) {
			s := *g
			// we do this because order doesnt matter in our case
			s[len(s)-1], s[i] = s[i], s[len(s)-1]
			*g = s[:len(s)-1]
			return true
		}
	}
	return false // remove failed
}

func (g *Graph) add(v Vertex) bool {
	*g = append(*g, v)
	return true
}

func pick(seen, unseen *Graph, distanceTable Table) Vertex {
	var min, minIdx int
	for i, p := range distanceTable {
		if !seen.contains(p.from) && p.distance < min {
			min = p.distance
			minIdx = i
		}
	}
	v := distanceTable[minIdx].from
	seen.add(v)
	unseen.remove(v)
	return v
}

func update(seen, unseen *Graph, distanceTable Table, v Vertex) {
	// get unvisited neighbors,
	// calculated shortest distance from

}

// For simplicity just going to take the first element as our start
func dj(graph Graph) Table {
	start := graph[0]
	// init
	var visited Graph
	unvisited := graph

	maxInt := 1 << 31
	var empty Vertex
	// initialize our result with empty/max values
	result := Table{
		row{
			start.id: path{
				start,
				0,
				empty,
			},
		},
	}
	for _, n := range graph {
		result = append(result, row{
			n.id: path{
				n,
				maxInt,
				empty,
			},
		})
	}
	// pick the unvisited vertex with the smallest known distance
	for len(unvisited) > 0 {
		v := pick(&visited, &unvisited, result)
		// fmt.Println(v)
		// fmt.Println(visited)
		// fmt.Println(unvisited)
		update(&visited, &unvisited, &result, v) // update result with this vertex
	}

	return result
}

func main() {
	/*
	   my graph looks like this
	   A ---5--- B ---2-- E
	   |         |        |
	   1         2        3
	   |         |        |
	   C ---1--- D -------
	*/
	A := Vertex{
		0,
		[]int{1, 2},
		[]int{5, 1},
	}

	B := Vertex{
		1,
		[]int{0, 3},
		[]int{5, 2},
	}

	C := Vertex{
		2,
		[]int{0, 3},
		[]int{1, 1},
	}

	D := Vertex{
		3,
		[]int{1, 2, 4},
		[]int{2, 1, 3},
	}

	E := Vertex{
		4,
		[]int{1, 3},
		[]int{2, 3},
	}

	simpleGraph := []Vertex{A, B, C, D, E}
	fmt.Println("Doing work..")
	// testing my definition
	// fmt.Println(A.isEqual(B))
	// fmt.Println(A.isEqual(A))
	// fmt.Println(A.getDistance(B))
	// fmt.Println(A.getDistance(E))
	dj(simpleGraph)
}
