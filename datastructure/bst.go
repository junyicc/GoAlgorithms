package datastructure

import (
	"sync"
)

// TreeNode struct
type TreeNode struct {
	key    int
	data   Elem
	lchild *TreeNode
	rchild *TreeNode
}

// BST is a binary search tree strcut
type BST struct {
	root *TreeNode
	lock sync.RWMutex
}

// Insert node
func (bst *BST) Insert(key int, data Elem) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	node := &TreeNode{key, data, nil, nil}
	if bst.root == nil {
		// change root
		bst.root = node
	} else {
		insertTreeNode(bst.root, node)
	}
}

func insertTreeNode(root, node *TreeNode) {
	if node.key <= root.key {
		if root.lchild == nil {
			root.lchild = node
		} else {
			insertTreeNode(root.lchild, node)
		}
	} else {
		if root.rchild == nil {
			root.rchild = node
		} else {
			insertTreeNode(root.rchild, node)
		}
	}
}

// Search returns true if node exists
func (bst *BST) Search(key int) (*Elem, bool) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return search(bst.root, key)
}

func search(node *TreeNode, key int) (*Elem, bool) {
	if node == nil {
		return nil, false
	}
	if key < node.key {
		return search(node.lchild, key)
	} else if key > node.key {
		return search(node.rchild, key)
	}

	return &node.data, true
}

// Remove node
func (bst *BST) Remove(key int) (*Elem, bool) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return remove(bst.root, key)
}

func remove(node *TreeNode, key int) (*Elem, bool) {
	if node == nil {
		return nil, false
	}
	if key < node.key {
		return remove(node.lchild, key)
	} else if key > node.key {
		return remove(node.rchild, key)
	}

	e := node.data // key == node.key
	// remove node
	if node.lchild == nil && node.rchild == nil {
		node = nil
		return &e, true
	}

	if node.lchild == nil {
		node = node.rchild
	} else if node.rchild == nil {
		node = node.lchild
	} else {
		r := node.lchild
		p := node
		for r.rchild != nil {
			p = r
			r = r.rchild
		}
		node.key = r.key
		node.data = r.data
		if p != node {
			p.rchild = r.lchild
		} else {
			p.lchild = r.lchild
		}
	}
	return &e, true
}

// PreOrderTraverseRecur visits BST nodes in pre-order recursively
func (bst *BST) PreOrderTraverseRecur(f func(Elem)) {
	bst.lock.RLock()
	bst.lock.RUnlock()

	preOrderTraverseRecur(bst.root, f)
}

func preOrderTraverseRecur(node *TreeNode, f func(Elem)) {
	if node != nil {
		f(node.data)
		preOrderTraverseRecur(node.lchild, f)
		preOrderTraverseRecur(node.rchild, f)
	}
}

// InOrderTraverseRecur visits BST nodes in in-order recursively
func (bst *BST) InOrderTraverseRecur(f func(Elem)) {
	bst.lock.RLock()
	bst.lock.RUnlock()

	inOrderTraverseRecur(bst.root, f)
}

func inOrderTraverseRecur(node *TreeNode, f func(Elem)) {
	if node != nil {
		inOrderTraverseRecur(node.lchild, f)
		f(node.data)
		inOrderTraverseRecur(node.rchild, f)
	}
}

// PostOrderTraverseRecur visits BST nodes in post-order recursively
func (bst *BST) PostOrderTraverseRecur(f func(Elem)) {
	bst.lock.RLock()
	bst.lock.RUnlock()

	postOrderTraverseRecur(bst.root, f)
}

func postOrderTraverseRecur(node *TreeNode, f func(Elem)) {
	if node != nil {
		postOrderTraverseRecur(node.lchild, f)
		postOrderTraverseRecur(node.rchild, f)
		f(node.data)
	}
}

// PreOrderTraverseIter visits BST nodes in pre-order iteratively
func (bst *BST) PreOrderTraverseIter(f func(Elem)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	preOrderTraverseIter(bst.root, f)
}

func preOrderTraverseIter(root *TreeNode, f func(Elem)) {
	if root == nil {
		return
	}
	stack := &DynamicStack{}
	stack = stack.New()
	stack.Push(root)

	for !stack.IsEmpty() {
		node := (*stack.Pop()).(*TreeNode)
		visitAlongLeft(stack, node, f)
	}
}

func visitAlongLeft(stack *DynamicStack, node *TreeNode, f func(Elem)) {
	for node != nil {
		f(node.data)
		stack.Push(node.rchild)
		node = node.lchild
	}
}

// InOrderTraverseIter visits BST nodes in in-order iteratively
func (bst *BST) InOrderTraverseIter(f func(Elem)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverseIter(bst.root, f)
}

func inOrderTraverseIter(root *TreeNode, f func(Elem)) {
	if root == nil {
		return
	}
	stack := &DynamicStack{}
	stack = stack.New()
	goAlongLeft(stack, root)

	for !stack.IsEmpty() {
		node := (*stack.Pop()).(*TreeNode)
		f(node.data)

		node = node.rchild
		goAlongLeft(stack, node)
	}
}

func goAlongLeft(stack *DynamicStack, node *TreeNode) {
	for node != nil {
		stack.Push(node)
		node = node.lchild
	}
}

// PostOrderTraverseIter visits BST nodes in post-order iteratively
func (bst *BST) PostOrderTraverseIter(f func(Elem)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	postOrderTraverseIter(bst.root, f)
}

func postOrderTraverseIter(root *TreeNode, f func(Elem)) {
	if root == nil {
		return
	}

	stack := &DynamicStack{}
	stack = stack.New()
	stack.Push(root)
	var pre *TreeNode
	for !stack.IsEmpty() {
		cur := (*stack.GetTop()).(*TreeNode)
		// traversing down the tree
		if pre == nil || pre.lchild == cur || pre.rchild == cur {
			if cur.lchild != nil {
				stack.Push(cur.lchild)
			} else if cur.rchild != nil {
				stack.Push(cur.rchild)
			} else {
				stack.Pop()
				f(cur.data)
			}
		} else if cur.lchild == pre {
			// traversing up the tree from the left
			if cur.rchild != nil {
				stack.Push(cur.rchild)
			} else {
				stack.Pop()
				f(cur.data)
			}
		} else if cur.rchild == pre {
			// traversing up the tree from the right
			stack.Pop()
			f(cur.data)
		}
		pre = cur
	}
}

// LevelTraverse traverse tree level by level
func (bst *BST) LevelTraverse(f func(Elem)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	levelTraverse(bst.root, f)
}

func levelTraverse(root *TreeNode, f func(Elem)) {
	if root == nil {
		return
	}

	var queue = &DynamicQueue{}
	queue = queue.New()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		node := (*queue.Dequeue()).(*TreeNode)
		f(node.data)

		if node.lchild != nil {
			queue.Enqueue(node.lchild)
		}
		if node.rchild != nil {
			queue.Enqueue(node.rchild)
		}
	}
}
