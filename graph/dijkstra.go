package graph

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// DijkstraPath returns shortest path using Dijkstra algorithm
func DijkstraPath(g *datastructure.GraphAdjMatrix, o, d int, f func(*datastructure.Vertex)) float64 {
	if g == nil || o < 0 || o >= g.VertexNum || d < 0 || d >= g.VertexNum {
		return datastructure.Inf
	}
	if o == d {
		return 0
	}

	path, weight := Dijkstra(g, o)

	ok := getPath(g, path, o, d, f)
	if !ok {
		return datastructure.Inf
	}
	return weight[d]
}

// Dijkstra algorithm uses f to find shortest path and returns its weight
func Dijkstra(g *datastructure.GraphAdjMatrix, source int) ([]int, []float64) {
	final, path, weight := make([]int, g.VertexNum), make([]int, g.VertexNum), make([]float64, g.VertexNum)
	// init weight final and path
	for i := 0; i < g.VertexNum; i++ {
		if g.E[source][i] != nil {
			weight[i] = g.E[source][i].Weight
			path[i] = source
		} else {
			weight[i] = datastructure.Inf
			path[i] = -1
		}
	}

	weight[source] = 0
	final[source] = 1

	for i := 1; i < g.VertexNum; i++ {
		min := datastructure.Inf
		k := 0
		for j := 0; j < g.VertexNum; j++ {
			if final[j] == 0 && weight[j] < min {
				k = j
				min = weight[j]
			}
		}

		final[k] = 1 // represents having finded shortest path to vertex k

		for j := 0; j < g.VertexNum; j++ {
			if g.E[k][j] != nil {
				if final[j] == 0 && min+g.E[k][j].Weight < weight[j] {
					weight[j] = min + g.E[k][j].Weight
					path[j] = k
				}
			}
		}
	}
	return path, weight
}

func getPath(g *datastructure.GraphAdjMatrix, path []int, o, d int, f func(*datastructure.Vertex)) bool {
	if path[d] == -1 {
		return false
	}

	if o == d {
		f(g.GetVertex(d))
	} else {
		ok := getPath(g, path, o, path[d], f)
		if !ok {
			return false
		}
		f(g.GetVertex(d))
	}
	return true
}
