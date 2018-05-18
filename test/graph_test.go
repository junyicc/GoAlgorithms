package test

import (
	"testing"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
	"github.com/CAIJUNYI/GoAlgorithms/graph"
)

var gl datastructure.GraphAdjList

func initGraphAdjList() {
	nA := datastructure.Vertex{"A"}
	nB := datastructure.Vertex{"B"}
	nC := datastructure.Vertex{"C"}
	nD := datastructure.Vertex{"D"}
	nE := datastructure.Vertex{"E"}
	nF := datastructure.Vertex{"F"}
	gl.InsertVertex(&nA)
	gl.InsertVertex(&nB)
	gl.InsertVertex(&nC)
	gl.InsertVertex(&nD)
	gl.InsertVertex(&nE)
	gl.InsertVertex(&nF)

	gl.InsertEdge(&nA, &nB, 1)
	gl.InsertEdge(&nA, &nC, 1)
	gl.InsertEdge(&nB, &nE, 1)
	gl.InsertEdge(&nC, &nE, 1)
	gl.InsertEdge(&nE, &nF, 1)
	gl.InsertEdge(&nD, &nA, 1)
}

func TestGraphAdjListBFS(t *testing.T) {
	initGraphAdjList()
	result := []*datastructure.Vertex{}
	visit := func(v *datastructure.Vertex) {
		result = append(result, v)
	}
	gl.BFS(gl.V[0], visit)
	exp := []*datastructure.Vertex{
		gl.GetVertex(0),
		gl.GetVertex(1),
		gl.GetVertex(2),
		gl.GetVertex(4),
		gl.GetVertex(5),
	}
	for i, e := range result {
		if e != exp[i] {
			t.Errorf("expected %v and got %v", exp[i], e)
		}
	}
}

func TestGraphAdjListDFS(t *testing.T) {
	initGraphAdjList()
	result := []*datastructure.Vertex{}
	visit := func(v *datastructure.Vertex) {
		result = append(result, v)
	}
	gl.DFS(gl.V[0], visit)
	exp := []*datastructure.Vertex{
		gl.GetVertex(0),
		gl.GetVertex(2),
		gl.GetVertex(4),
		gl.GetVertex(5),
		gl.GetVertex(1),
	}
	for i, e := range result {
		if e != exp[i] {
			t.Errorf("expected %v and got %v", exp[i], e)
		}
	}
}

var gm datastructure.GraphAdjMatrix

func initGraphAdjMatrix() {
	nA := datastructure.Vertex{"A"}
	nB := datastructure.Vertex{"B"}
	nC := datastructure.Vertex{"C"}
	nD := datastructure.Vertex{"D"}
	nE := datastructure.Vertex{"E"}
	nF := datastructure.Vertex{"F"}
	gm.InsertVertex(&nA)
	gm.InsertVertex(&nB)
	gm.InsertVertex(&nC)
	gm.InsertVertex(&nD)
	gm.InsertVertex(&nE)
	gm.InsertVertex(&nF)

	gm.InsertEdge(0, 1, 1)
	gm.InsertEdge(0, 2, 1)
	gm.InsertEdge(1, 4, 1)
	gm.InsertEdge(2, 4, 1)
	gm.InsertEdge(4, 5, 1)
	gm.InsertEdge(3, 0, 1)
}

func TestGraphAdjMatrix(t *testing.T) {
	initGraphAdjMatrix()
	e, err := gm.RemoveEdge(3, 0)
	if err != nil || e.From != gm.GetVertex(3) || e.To != gm.GetVertex(0) || e.Weight != 1 {
		t.Errorf("failed removeing edge")
	}

	v, err := gm.RemoveVertex(0)
	if err != nil || v.Value != "A" {
		t.Errorf("failed removing vertex")
	}

}

func TestGraphAdjMatrixBFS(t *testing.T) {
	initGraphAdjMatrix()
	result := []*datastructure.Vertex{}
	visit := func(v *datastructure.Vertex) {
		result = append(result, v)
	}
	gm.BFS(0, visit)
	exp := []*datastructure.Vertex{
		gm.GetVertex(0),
		gm.GetVertex(1),
		gm.GetVertex(2),
		gm.GetVertex(4),
		gm.GetVertex(5),
	}
	for i, e := range result {
		if e != exp[i] {
			t.Errorf("expected %v and got %v", exp[i], e)
		}
	}
}

func TestGraphAdjMatrixDFS(t *testing.T) {
	initGraphAdjMatrix()
	result := []*datastructure.Vertex{}
	visit := func(v *datastructure.Vertex) {
		result = append(result, v)
	}
	gm.DFS(0, visit)
	exp := []*datastructure.Vertex{
		gm.GetVertex(0),
		gm.GetVertex(2),
		gm.GetVertex(4),
		gm.GetVertex(5),
		gm.GetVertex(1),
	}
	for i, e := range result {
		if e != exp[i] {
			t.Errorf("expected %v and got %v", exp[i], e)
		}
	}
}

func TestPrim(t *testing.T) {
	var gm datastructure.GraphAdjMatrix

	v0 := datastructure.Vertex{"0"}
	v1 := datastructure.Vertex{"1"}
	v2 := datastructure.Vertex{"2"}
	v3 := datastructure.Vertex{"3"}
	v4 := datastructure.Vertex{"4"}
	v5 := datastructure.Vertex{"5"}
	v6 := datastructure.Vertex{"6"}
	v7 := datastructure.Vertex{"7"}
	v8 := datastructure.Vertex{"8"}

	gm.InsertVertex(&v0)
	gm.InsertVertex(&v1)
	gm.InsertVertex(&v2)
	gm.InsertVertex(&v3)
	gm.InsertVertex(&v4)
	gm.InsertVertex(&v5)
	gm.InsertVertex(&v6)
	gm.InsertVertex(&v7)
	gm.InsertVertex(&v8)

	gm.InsertEdge(0, 1, 10)
	gm.InsertEdge(1, 0, 10)
	gm.InsertEdge(0, 5, 11)
	gm.InsertEdge(5, 0, 11)
	gm.InsertEdge(1, 2, 18)
	gm.InsertEdge(2, 1, 18)
	gm.InsertEdge(1, 6, 16)
	gm.InsertEdge(6, 1, 16)
	gm.InsertEdge(1, 8, 12)
	gm.InsertEdge(8, 1, 12)
	gm.InsertEdge(5, 6, 17)
	gm.InsertEdge(6, 5, 17)
	gm.InsertEdge(5, 4, 26)
	gm.InsertEdge(4, 5, 26)
	gm.InsertEdge(6, 7, 19)
	gm.InsertEdge(7, 6, 19)
	gm.InsertEdge(7, 4, 7)
	gm.InsertEdge(4, 7, 7)
	gm.InsertEdge(7, 3, 16)
	gm.InsertEdge(3, 7, 16)
	gm.InsertEdge(3, 4, 20)
	gm.InsertEdge(4, 3, 20)
	gm.InsertEdge(6, 3, 24)
	gm.InsertEdge(3, 6, 24)
	gm.InsertEdge(8, 3, 21)
	gm.InsertEdge(3, 8, 21)
	gm.InsertEdge(2, 3, 22)
	gm.InsertEdge(3, 2, 22)
	gm.InsertEdge(8, 2, 8)
	gm.InsertEdge(2, 8, 8)
	var result []*datastructure.Edge
	visit := func(e *datastructure.Edge) {
		result = append(result, e)
	}

	graph.Prim(&gm, 0, visit)

	exp := []*datastructure.Edge{
		gm.GetEdge(0, 1),
		gm.GetEdge(0, 5),
		gm.GetEdge(1, 8),
		gm.GetEdge(8, 2),
		gm.GetEdge(1, 6),
		gm.GetEdge(6, 7),
		gm.GetEdge(7, 4),
		gm.GetEdge(7, 3),
	}

	for i, e := range exp {
		if e != result[i] {
			t.Errorf("expected %v\n got %v\n", e, result[i])
		}
	}

}
