package leetcode

// Given a reference of a node in a connected undirected graph.

// Return a deep copy (clone) of the graph.

// Each node in the graph contains a val (int) and a list (List[Vertex]) of its neighbors.

func cloneGraph(node *Vertex) *Vertex {
	if node == nil {
		return nil
	}
	visited := make(map[*Vertex]*Vertex)
	return cloneGraphVertex(node, visited)
}

func cloneGraphVertex(vertex *Vertex, visited map[*Vertex]*Vertex) *Vertex {
	if vertex == nil {
		return nil
	}
	if v, ok := visited[vertex]; ok {
		return v
	}

	cloneVertex := &Vertex{
		Val:       vertex.Val,
		Neighbors: make([]*Vertex, len(vertex.Neighbors)),
	}
	visited[vertex] = cloneVertex

	for i, neighborVertex := range vertex.Neighbors {
		cloneVertex.Neighbors[i] = cloneGraphVertex(neighborVertex, visited)
	}
	return cloneVertex
}
