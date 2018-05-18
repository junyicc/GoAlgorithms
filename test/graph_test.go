package test

import (
	"testing"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
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
