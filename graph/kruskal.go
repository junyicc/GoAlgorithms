package graph

import (
	"fmt"
	"sort"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// Edge struct for Kruskal MST algorithm
type Edge struct {
	Begin  int
	End    int
	Weight float64
}

func (e Edge) String() string {
	return fmt.Sprintf("(%d, %d): %f", e.Begin, e.End, e.Weight)
}

// Edges is a series of Edge
type Edges struct {
	Data []Edge
	Cmp  func(Edge, Edge) bool
}

func (e Edges) Len() int           { return len(e.Data) }
func (e Edges) Swap(i, j int)      { e.Data[i], e.Data[j] = e.Data[j], e.Data[i] }
func (e Edges) Less(i, j int) bool { return e.Cmp(e.Data[i], e.Data[j]) }

func toEdgeArray(g *datastructure.GraphAdjMatrix) []Edge {
	var edges []Edge
	edgeMatrix := g.E
	for i, edgeArray := range edgeMatrix {
		for j, e := range edgeArray {
			if e != nil && e.Weight != 0 {
				edge := Edge{
					Begin:  i,
					End:    j,
					Weight: e.Weight,
				}
				edges = append(edges, edge)
			}
		}
	}
	edgesWrapper := Edges{
		Data: edges,
		Cmp: func(p, q Edge) bool {
			if p.Weight < q.Weight {
				return true
			} else if p.Weight == q.Weight {
				return p.Begin < q.Begin
			}
			return false
		},
	}
	sort.Sort(edgesWrapper)
	edges = edgesWrapper.Data
	return edges
}

// Kruskal MST algorithm
func Kruskal(g *datastructure.GraphAdjMatrix, f func(*Edge)) {
	vNum := g.VertexNum
	eNum := g.EdgeNum

	edges := toEdgeArray(g)
	parent := make([]int, vNum)

	for i := 0; i < eNum; i++ {
		n := find(parent, edges[i].Begin)
		m := find(parent, edges[i].End)
		if n != m {
			// there is no loop
			parent[n] = m
			f(&edges[i])
		}
	}
}

func find(parent []int, f int) int {
	for parent[f] > 0 {
		f = parent[f]
	}
	return f
}
