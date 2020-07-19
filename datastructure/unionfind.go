package datastructure

// UnionFind 用来解决图的动态连通性问题
//
// 动态连通性其实就是一种等价关系，具有「自反性」「传递性」和「对称性」
// 1、自反性：节点p和p是连通的。
// 2、对称性：如果节点p和q连通，那么q和p也连通。
// 3、传递性：如果节点p和q连通，q和r连通，那么p和r也连通。

// 使用：主要思路是适时增加虚拟节点，想办法让元素「分门别类」，建立动态连通关系。

type UnionFind interface {
	Connected(p, q interface{}) bool
	Union(p, q interface{})
	Count() int
}

type UnionFindInt struct {
	parent []int
	size   []int
	count  int
}

func NewUnionFindInt(n int) *UnionFindInt {
	uf := UnionFindInt{
		parent: make([]int, n),
		size:   make([]int, n),
		count:  n,
	}
	for i := 0; i < len(uf.parent); i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return &uf
}

func (u *UnionFindInt) Connected(p, q interface{}) bool {
	rootP := u.find(p.(int))
	rootQ := u.find(q.(int))
	return rootP == rootQ
}

func (u *UnionFindInt) Union(p, q interface{}) {
	rootP := u.find(p.(int))
	rootQ := u.find(q.(int))

	if rootP == rootQ {
		return
	}

	// attach to the node with more children
	// lower the tree height
	if u.size[rootP] > u.size[rootQ] {
		u.parent[rootQ] = rootP
		u.size[rootP] += u.size[rootQ]
	} else {
		u.parent[rootP] = rootQ
		u.size[rootQ] += u.size[rootP]
	}
	u.count--
}

func (u *UnionFindInt) Count() int {
	return u.count
}

func (u *UnionFindInt) find(x int) int {
	for u.parent[x] != x {
		// compact path
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}
