package graph

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// FloydPath returns shortest path using Floyd algorithm
func FloydPath(g *datastructure.GraphAdjMatrix, o, d int, f func(*datastructure.Vertex)) float64 {
	if g == nil || o < 0 || o >= g.VertexNum || d < 0 || d >= g.VertexNum {
		return datastructure.Inf
	}
	if o == d {
		return 0
	}

	path, weight := Floyd(g)
	getFloydPath(g, path, o, d, f)
	return weight[o][d]
}

// Floyd algorithm: find shortest path from n to n
func Floyd(g *datastructure.GraphAdjMatrix) ([][]int, [][]float64) {
	path, weight := make([][]int, g.VertexNum), make([][]float64, g.VertexNum)
	for i := 0; i < g.VertexNum; i++ {
		path[i] = make([]int, g.VertexNum)
		weight[i] = make([]float64, g.VertexNum)
	}
	// init path and weight matrix
	for i := 0; i < g.VertexNum; i++ {
		for j := 0; j < g.VertexNum; j++ {
			path[i][j] = j
			if g.E[i][j] != nil {
				weight[i][j] = g.E[i][j].Weight
			} else {
				weight[i][j] = datastructure.Inf
			}
		}
	}

	for i := 0; i < g.VertexNum; i++ {
		for j := 0; j < g.VertexNum; j++ {
			for k := 0; k < g.VertexNum; k++ {
				if weight[j][i]+weight[i][k] < weight[j][k] {
					weight[j][k] = weight[j][i] + weight[i][k]
					path[j][k] = path[j][i]
				}
			}
		}
	}
	return path, weight
}

func getFloydPath(g *datastructure.GraphAdjMatrix, path [][]int, o, d int, f func(*datastructure.Vertex)) {
	v := path[o][d]
	if v == d {
		return
	}
	f(g.GetVertex(o))
	for v != d {
		f(g.GetVertex(v))
		v = path[v][d]
	}
	f(g.GetVertex(v))
}
