package graph

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
	"testing"
)

func TestPrim(t *testing.T) {
	var gm datastructure.GraphAdjMatrix

	v0 := datastructure.Vertex{Value: "0"}
	v1 := datastructure.Vertex{Value: "1"}
	v2 := datastructure.Vertex{Value: "2"}
	v3 := datastructure.Vertex{Value: "3"}
	v4 := datastructure.Vertex{Value: "4"}
	v5 := datastructure.Vertex{Value: "5"}
	v6 := datastructure.Vertex{Value: "6"}
	v7 := datastructure.Vertex{Value: "7"}
	v8 := datastructure.Vertex{Value: "8"}

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

	Prim(&gm, 0, visit)

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

func TestKruskal(t *testing.T) {
	var gm datastructure.GraphAdjMatrix

	v0 := datastructure.Vertex{Value: "0"}
	v1 := datastructure.Vertex{Value: "1"}
	v2 := datastructure.Vertex{Value: "2"}
	v3 := datastructure.Vertex{Value: "3"}
	v4 := datastructure.Vertex{Value: "4"}
	v5 := datastructure.Vertex{Value: "5"}
	v6 := datastructure.Vertex{Value: "6"}
	v7 := datastructure.Vertex{Value: "7"}
	v8 := datastructure.Vertex{Value: "8"}

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

	var result []*Edge
	visit := func(e *Edge) {
		result = append(result, e)
	}
	exp := []*Edge{
		{Begin: 4, End: 7, Weight: 7},
		{Begin: 2, End: 8, Weight: 8},
		{Begin: 0, End: 1, Weight: 10},
		{Begin: 0, End: 5, Weight: 11},
		{Begin: 1, End: 8, Weight: 12},
		{Begin: 1, End: 6, Weight: 16},
		{Begin: 3, End: 7, Weight: 16},
		{Begin: 6, End: 7, Weight: 19},
	}

	Kruskal(&gm, visit)

	for i, r := range result {
		if *r != *exp[i] {
			t.Errorf("expected %v\ngot %v\n", *exp[i], *r)
		}
	}

}

func TestDijkstra(t *testing.T) {
	var gm datastructure.GraphAdjMatrix

	v0 := datastructure.Vertex{Value: "0"}
	v1 := datastructure.Vertex{Value: "1"}
	v2 := datastructure.Vertex{Value: "2"}
	v3 := datastructure.Vertex{Value: "3"}
	v4 := datastructure.Vertex{Value: "4"}
	v5 := datastructure.Vertex{Value: "5"}

	gm.InsertVertex(&v0)
	gm.InsertVertex(&v1)
	gm.InsertVertex(&v2)
	gm.InsertVertex(&v3)
	gm.InsertVertex(&v4)
	gm.InsertVertex(&v5)

	gm.InsertEdge(0, 1, 1)
	gm.InsertEdge(0, 2, 7)
	gm.InsertEdge(1, 3, 9)
	gm.InsertEdge(1, 5, 15)
	gm.InsertEdge(2, 4, 4)
	gm.InsertEdge(3, 4, 10)
	gm.InsertEdge(3, 5, 5)
	gm.InsertEdge(4, 5, 3)

	var result []*datastructure.Vertex
	visit := func(e *datastructure.Vertex) {
		result = append(result, e)
	}

	w1 := DijkstraPath(&gm, 0, 5, visit)
	expPath1 := []*datastructure.Vertex{
		gm.GetVertex(0),
		gm.GetVertex(2),
		gm.GetVertex(4),
		gm.GetVertex(5),
	}

	if w1 != 14 {
		t.Errorf("expected 14 and got %f", w1)
	}

	for i, e := range result {
		if e != expPath1[i] {
			t.Errorf("expected %v\ngot %v\n", expPath1[i], e)
		}
	}
	result = []*datastructure.Vertex{}
	w2 := DijkstraPath(&gm, 1, 2, visit)

	if w2 != datastructure.Inf {
		t.Errorf("expected inf and got %f", w2)
	}

	if len(result) != 0 {
		t.Errorf("expected nil path and got %v\n", result)
	}

}

func TestFloyd(t *testing.T) {
	var gm datastructure.GraphAdjMatrix

	v0 := datastructure.Vertex{Value: "0"}
	v1 := datastructure.Vertex{Value: "1"}
	v2 := datastructure.Vertex{Value: "2"}
	v3 := datastructure.Vertex{Value: "3"}
	v4 := datastructure.Vertex{Value: "4"}
	v5 := datastructure.Vertex{Value: "5"}

	gm.InsertVertex(&v0)
	gm.InsertVertex(&v1)
	gm.InsertVertex(&v2)
	gm.InsertVertex(&v3)
	gm.InsertVertex(&v4)
	gm.InsertVertex(&v5)

	gm.InsertEdge(0, 1, 1)
	gm.InsertEdge(0, 2, 7)
	gm.InsertEdge(1, 3, 9)
	gm.InsertEdge(1, 5, 15)
	gm.InsertEdge(2, 4, 4)
	gm.InsertEdge(3, 4, 10)
	gm.InsertEdge(3, 5, 5)
	gm.InsertEdge(4, 5, 3)

	var result []*datastructure.Vertex
	visit := func(e *datastructure.Vertex) {
		result = append(result, e)
	}

	w1 := FloydPath(&gm, 0, 5, visit)
	expPath1 := []*datastructure.Vertex{
		gm.GetVertex(0),
		gm.GetVertex(2),
		gm.GetVertex(4),
		gm.GetVertex(5),
	}

	if w1 != 14 {
		t.Errorf("expected 14 and got %f", w1)
	}

	for i, e := range result {
		if e != expPath1[i] {
			t.Errorf("expected %v\ngot %v\n", expPath1[i], e)
		}
	}
	result = []*datastructure.Vertex{}
	w2 := FloydPath(&gm, 2, 3, visit)

	if w2 != datastructure.Inf {
		t.Errorf("expected inf and got %f", w2)
	}

	if len(result) != 0 {
		t.Errorf("expected nil path and got %v\n", result)
	}
}

func TestTopoSort(t *testing.T) {
	var gl datastructure.GraphAdjList
	v0 := datastructure.Vertex{Value: "0"}
	v1 := datastructure.Vertex{Value: "1"}
	v2 := datastructure.Vertex{Value: "2"}
	v3 := datastructure.Vertex{Value: "3"}
	v4 := datastructure.Vertex{Value: "4"}
	v5 := datastructure.Vertex{Value: "5"}
	v6 := datastructure.Vertex{Value: "6"}
	v7 := datastructure.Vertex{Value: "7"}
	v8 := datastructure.Vertex{Value: "8"}
	v9 := datastructure.Vertex{Value: "9"}
	v10 := datastructure.Vertex{Value: "10"}
	v11 := datastructure.Vertex{Value: "11"}
	v12 := datastructure.Vertex{Value: "12"}
	v13 := datastructure.Vertex{Value: "13"}

	gl.InsertVertex(&v0)
	gl.InsertVertex(&v1)
	gl.InsertVertex(&v2)
	gl.InsertVertex(&v3)
	gl.InsertVertex(&v4)
	gl.InsertVertex(&v5)
	gl.InsertVertex(&v6)
	gl.InsertVertex(&v7)
	gl.InsertVertex(&v8)
	gl.InsertVertex(&v9)
	gl.InsertVertex(&v10)
	gl.InsertVertex(&v11)
	gl.InsertVertex(&v12)
	gl.InsertVertex(&v13)

	gl.InsertEdge(&v0, &v4, 1)
	gl.InsertEdge(&v0, &v5, 1)
	gl.InsertEdge(&v0, &v11, 1)
	gl.InsertEdge(&v1, &v2, 1)
	gl.InsertEdge(&v1, &v4, 1)
	gl.InsertEdge(&v1, &v8, 1)
	gl.InsertEdge(&v2, &v5, 1)
	gl.InsertEdge(&v2, &v6, 1)
	gl.InsertEdge(&v2, &v9, 1)
	gl.InsertEdge(&v3, &v2, 1)
	gl.InsertEdge(&v3, &v13, 1)
	gl.InsertEdge(&v4, &v7, 1)
	gl.InsertEdge(&v5, &v8, 1)
	gl.InsertEdge(&v5, &v12, 1)
	gl.InsertEdge(&v6, &v5, 1)
	gl.InsertEdge(&v8, &v7, 1)
	gl.InsertEdge(&v9, &v10, 1)
	gl.InsertEdge(&v9, &v11, 1)
	gl.InsertEdge(&v10, &v13, 1)
	gl.InsertEdge(&v12, &v9, 1)

	var result []*datastructure.Vertex
	visit := func(e *datastructure.Vertex) {
		result = append(result, e)
	}

	ok := TopoSort(&gl, visit)
	if !ok {
		t.Errorf("expected true and got %v", ok)
	}
}
