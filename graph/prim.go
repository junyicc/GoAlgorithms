package graph

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// Prim minimum spanning tree algorithm
func Prim(g *datastructure.GraphAdjMatrix, source int, f func(*datastructure.Edge)) {
	vNum := len(g.V)
	if g == nil || source < 0 || source >= vNum {
		return
	}

	adjVex := make([]int, vNum)      // adjacent vertex
	lowcost := make([]float64, vNum) // adjacent vertex weight

	// init adjacent vertex and lowcost weight array
	for i := 0; i < vNum; i++ {
		if g.E[source][i] != nil {
			lowcost[i] = g.E[source][i].Weight
		} else {
			lowcost[i] = datastructure.Inf
		}
		adjVex[i] = source
	}

	for i := 0; i < vNum; i++ {
		min := datastructure.Inf
		// find nearest vertex k
		k := 0
		for j := 0; j < vNum; j++ {
			if lowcost[j] != 0 && lowcost[j] < min {
				min = lowcost[j]
				k = j
			}
		}
		if adjVex[k] != k {
			f(g.GetEdge(adjVex[k], k))
		}
		lowcost[k] = 0 // k vertex has been visited
		// update weights from k to other vertexes
		for j := 0; j < vNum; j++ {
			if g.E[k][j] != nil {
				if lowcost[j] != 0 && g.E[k][j].Weight < lowcost[j] {
					lowcost[j] = g.E[k][j].Weight
					adjVex[j] = k
				}
			}
		}
	}
}
