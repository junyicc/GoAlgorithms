package graph

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// TopoSort returns true if graph is a DAG
func TopoSort(g *datastructure.GraphAdjList, f func(*datastructure.Vertex)) bool {
	if g == nil {
		return false
	}
	stack := datastructure.NewDynamicStack()
	var cnt int
	for _, v := range g.V {
		if v.In == 0 {
			stack.Push(v)
		}
	}

	for !stack.IsEmpty() {
		v := (*stack.Pop()).(*datastructure.Vertex)
		f(v)
		cnt++

		edges := g.E[v]
		if edges != nil {
			for _, e := range edges {
				e.To.In--
				if e.To.In == 0 {
					stack.Push(e.To)
				}
			}
		}
	}

	if cnt == len(g.V) {
		return true
	}
	return false
}
