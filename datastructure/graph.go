package datastructure

import (
	"fmt"
	"math"
	"sync"
)

// Inf represents infinite weight
var Inf = math.Inf(1)

func isInf(f float64) bool {
	return math.IsInf(f, 1)
}

// Vertex of Graph
type Vertex struct {
	Value Elem
	In    int // in degree
}

func (v Vertex) String() string {
	return fmt.Sprintf("%v", v.Value)
}

// Edge struct
type Edge struct {
	From   *Vertex
	To     *Vertex
	Weight float64
}

// New an edge
func (e Edge) New() *Edge {
	e = Edge{
		From:   &Vertex{},
		To:     &Vertex{},
		Weight: Inf,
	}
	return &e
}

// GraphAdjList is an adjacency list strcut of graph
type GraphAdjList struct {
	V    []*Vertex
	E    map[*Vertex][]*Edge
	lock sync.RWMutex
}

// GetVertex retuerns a vertex
func (g *GraphAdjList) GetVertex(k int) *Vertex {
	if k < 0 || k >= len(g.V) {
		return nil
	}
	g.lock.RLock()
	v := g.V[k]
	g.lock.RUnlock()
	return v
}

// GetEdge returns an edge
func (g *GraphAdjList) GetEdge(from, to *Vertex) *Edge {
	if from == nil || to == nil {
		return nil
	}

	near, ok := g.E[from]
	if !ok {
		return nil
	}
	for _, e := range near {
		if e.To == to {
			return e
		}
	}
	return nil
}

// InsertVertex inserts a vertex
func (g *GraphAdjList) InsertVertex(v *Vertex) {
	g.lock.Lock()
	g.V = append(g.V, v)
	g.lock.Unlock()
}

// RemoveVertex removes a vertex
func (g *GraphAdjList) RemoveVertex(k int) error {
	if k < 0 || k >= len(g.V) {
		return fmt.Errorf("RemoveVertex: %d out of range", k)
	}
	g.lock.Lock()
	vertex := g.V[k]
	edge := g.E[vertex]
	// remove vertex
	g.V = append(g.V[:k], g.V[k+1:]...)
	// remove edge
	for _, e := range edge {
		g.RemoveEdge(vertex, e.To)
		g.RemoveEdge(e.To, vertex)
	}
	delete(g.E, vertex)
	g.lock.Unlock()

	return nil
}

// InsertEdge inserts an edge
func (g *GraphAdjList) InsertEdge(from, to *Vertex, weight float64) {
	g.lock.Lock()
	if g.E == nil {
		g.E = make(map[*Vertex][]*Edge)
	}
	e := Edge{From: from, To: to, Weight: weight}
	g.E[from] = append(g.E[from], &e)
	(*to).In++
	g.lock.Unlock()
}

// RemoveEdge removes an edge
func (g *GraphAdjList) RemoveEdge(from, to *Vertex) error {
	edge, ok := g.E[from]
	if !ok {
		return fmt.Errorf("RemoveEdge: does not exist vertex %v", from)
	}

	g.lock.Lock()
	for i, e := range edge {
		if to == e.To {
			g.E[from] = append(g.E[from][:i], g.E[from][i+1:]...)
			e.To.In--
			break
		}
	}
	g.lock.Unlock()
	return nil
}

// BFS traverse graph adjacency list
func (g *GraphAdjList) BFS(v *Vertex, f func(*Vertex)) {
	if v == nil || f == nil {
		return
	}
	g.lock.Lock()
	defer g.lock.Unlock()

	q := new(DynamicQueue)
	q = q.New()
	discovered := make(map[*Vertex]bool)
	q.Enqueue(v)
	discovered[v] = true
	for !q.IsEmpty() {
		v := (*q.Dequeue()).(*Vertex)
		f(v)
		near := g.E[v]
		for _, e := range near {
			if !discovered[e.To] {
				q.Enqueue(e.To)
				discovered[e.To] = true
			}
		}
	}
}

// DFS traverse graph adjacency list
func (g *GraphAdjList) DFS(v *Vertex, f func(*Vertex)) {
	if v == nil || f == nil {
		return
	}

	g.lock.Lock()
	defer g.lock.Unlock()

	s := &DynamicStack{}
	s = s.New()
	discovered := make(map[*Vertex]bool)
	s.Push(v)
	discovered[v] = true
	for !s.IsEmpty() {
		v := (*s.Pop()).(*Vertex)
		f(v)
		near := g.E[v]
		for _, e := range near {
			if !discovered[e.To] {
				s.Push(e.To)
				discovered[e.To] = true
			}
		}
	}
}

// ----------------------------------------------------------------------------

// GraphAdjMatrix is an adjacency matrix struct of graph
type GraphAdjMatrix struct {
	V         []*Vertex
	E         [][]*Edge
	VertexNum int
	EdgeNum   int
	lock      sync.RWMutex
}

// GetVertex returns vertex
func (g *GraphAdjMatrix) GetVertex(i int) *Vertex {
	if i < 0 || i >= len(g.V) {
		return nil
	}
	g.lock.RLock()
	v := g.V[i]
	g.lock.RUnlock()
	return v
}

// GetEdge returns edge
func (g *GraphAdjMatrix) GetEdge(i, j int) *Edge {
	if i < 0 || j < 0 || i >= len(g.V) || j >= len(g.V) {
		return nil
	}
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.E[i][j]
}

// InsertVertex inserts a vertex
func (g *GraphAdjMatrix) InsertVertex(v *Vertex) {
	g.lock.Lock()
	// insert edge
	e := make([]*Edge, len(g.V)+1)
	g.E = append(g.E, e)
	for i := 0; i < len(g.V); i++ {
		g.E[i] = append(g.E[i], nil)
	}

	// insert vertex
	g.V = append(g.V, v)
	g.E[len(g.V)-1][len(g.V)-1] = &Edge{From: v, To: v, Weight: 0}
	g.VertexNum++
	g.lock.Unlock()
}

// RemoveVertex removes a vertex
func (g *GraphAdjMatrix) RemoveVertex(k int) (*Vertex, error) {
	if k < 0 || k >= len(g.V) {
		return nil, fmt.Errorf("RemoveVertex: %d out of range", k)
	}
	g.lock.Lock()
	v := g.V[k]
	// remove edge
	for i := 0; i < len(g.V); i++ {
		g.E[i] = append(g.E[i][:k], g.E[i][k+1:]...)
	}
	g.E = append(g.E[:k], g.E[k+1:]...)
	// remove vertex
	g.V = append(g.V[:k], g.V[k+1:]...)
	g.VertexNum--
	g.lock.Unlock()

	return v, nil
}

// InsertEdge inserts an edge
func (g *GraphAdjMatrix) InsertEdge(i, j int, weight float64) error {
	from := g.GetVertex(i)
	to := g.GetVertex(j)
	if from == nil || to == nil {
		return fmt.Errorf("InsertEdge: cannot insert the edge from %d to %d", i, j)
	}
	if i == j {
		weight = 0
	}

	g.E[i][j] = &Edge{From: from, To: to, Weight: weight}
	to.In++
	g.EdgeNum++
	return nil
}

// RemoveEdge removes an edge
func (g *GraphAdjMatrix) RemoveEdge(i, j int) (*Edge, error) {
	if i < 0 || i >= len(g.V) || j < 0 || j >= len(g.V) {
		return nil, fmt.Errorf("RemoveEdge: index out of range")
	}

	e := *g.E[i][j]
	g.E[i][j].To.In--
	g.E[i][j] = nil
	g.EdgeNum--
	return &e, nil
}

// BFS algorithm
func (g *GraphAdjMatrix) BFS(k int, f func(*Vertex)) {
	q := new(DynamicQueue)
	q = q.New()
	discovered := make(map[int]bool)
	q.Enqueue(k)
	discovered[k] = true
	for !q.IsEmpty() {
		i := (*q.Dequeue()).(int)
		v := g.GetVertex(i)
		f(v)
		edge := g.E[i]
		for j, e := range edge {
			if e != nil && !discovered[j] {
				q.Enqueue(j)
				discovered[j] = true
			}
		}
	}
}

// DFS algorithm
func (g *GraphAdjMatrix) DFS(k int, f func(*Vertex)) {
	s := new(DynamicStack)
	s = s.New()
	discovered := make(map[int]bool)
	s.Push(k)
	discovered[k] = true

	for !s.IsEmpty() {
		i := (*s.Pop()).(int)
		v := g.GetVertex(i)
		f(v)
		edge := g.E[i]
		for j, e := range edge {
			if e != nil && !discovered[j] {
				s.Push(j)
				discovered[j] = true
			}
		}
	}
}
