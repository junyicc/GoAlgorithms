package datastructure

import (
	"testing"
)

var gl GraphAdjList

func initGraphAdjList() {
	nA := Vertex{Value: "A"}
	nB := Vertex{Value: "B"}
	nC := Vertex{Value: "C"}
	nD := Vertex{Value: "D"}
	nE := Vertex{Value: "E"}
	nF := Vertex{Value: "F"}
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

func TestGraphAdjList(t *testing.T) {
	initGraphAdjList()
	if i := gl.GetVertex(4).In; i != 2 {
		t.Errorf("expected 2 and got %d", i)
	}
}

func TestGraphAdjListBFS(t *testing.T) {
	initGraphAdjList()
	var result []*Vertex
	visit := func(v *Vertex) {
		result = append(result, v)
	}
	gl.BFS(gl.V[0], visit)
	exp := []*Vertex{
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
	var result []*Vertex
	visit := func(v *Vertex) {
		result = append(result, v)
	}
	gl.DFS(gl.V[0], visit)
	exp := []*Vertex{
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

var gm GraphAdjMatrix

func initGraphAdjMatrix() {
	nA := Vertex{Value: "A"}
	nB := Vertex{Value: "B"}
	nC := Vertex{Value: "C"}
	nD := Vertex{Value: "D"}
	nE := Vertex{Value: "E"}
	nF := Vertex{Value: "F"}
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
	var result []*Vertex
	visit := func(v *Vertex) {
		result = append(result, v)
	}
	gm.BFS(0, visit)
	exp := []*Vertex{
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
	result := []*Vertex{}
	visit := func(v *Vertex) {
		result = append(result, v)
	}
	gm.DFS(0, visit)
	exp := []*Vertex{
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
