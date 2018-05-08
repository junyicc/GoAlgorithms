package datastructure

import (
	"sync"
)

// TreeNode struct
type TreeNode struct {
	Key    int
	Data   Elem
	LChild *TreeNode
	RChild *TreeNode
}

// BST is a binary search tree strcut
type BST struct {
	Root *TreeNode
	lock sync.RWMutex
}

// Insert node
func (bst *BST) Insert(key int, data Elem) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	node := &TreeNode{key, data, nil, nil}
	if bst.Root == nil {
		// change root
		bst.Root = node
	} else {
		insertTreeNode(bst.Root, node)
	}
}

func insertTreeNode(root, node *TreeNode) {
	if node.Key <= root.Key {
		if root.LChild == nil {
			root.LChild = node
		} else {
			insertTreeNode(root.LChild, node)
		}
	} else {
		if root.RChild == nil {
			root.RChild = node
		} else {
			insertTreeNode(root.RChild, node)
		}
	}
}

// Search returns true if node exists
func (bst *BST) Search(key int) (*Elem, bool) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return search(bst.Root, key)
}

func search(node *TreeNode, key int) (*Elem, bool) {
	if node == nil {
		return nil, false
	}
	if key < node.Key {
		return search(node.LChild, key)
	} else if key > node.Key {
		return search(node.RChild, key)
	}

	return &node.Data, true
}

// Remove node
func (bst *BST) Remove(key int) (*Elem, bool) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	return remove(bst.Root, key)
}

func remove(node *TreeNode, key int) (*Elem, bool) {
	if node == nil {
		return nil, false
	}
	if key < node.Key {
		return remove(node.LChild, key)
	} else if key > node.Key {
		return remove(node.RChild, key)
	}

	e := node.Data // key == node.Key
	// remove node
	if node.LChild == nil && node.RChild == nil {
		node = nil
		return &e, true
	}

	if node.LChild == nil {
		node = node.RChild
	} else if node.RChild == nil {
		node = node.LChild
	} else {
		r := node.LChild
		p := node
		for r.RChild != nil {
			p = r
			r = r.RChild
		}
		node.Key = r.Key
		node.Data = r.Data
		if p != node {
			p.RChild = r.LChild
		} else {
			p.LChild = r.LChild
		}
	}
	return &e, true
}

// PreOrderTraverseRecur visits BST nodes in pre-order recursively
func (bst *BST) PreOrderTraverseRecur(f func(Elem)) {
	bst.lock.RLock()
	bst.lock.RUnlock()

	preOrderTraverseRecur(bst.Root, f)
}

func preOrderTraverseRecur(node *TreeNode, f func(Elem)) {
	if node != nil {
		f(node.Data)
		preOrderTraverseRecur(node.LChild, f)
		preOrderTraverseRecur(node.RChild, f)
	}
}

// InOrderTraverseRecur visits BST nodes in in-order recursively
func (bst *BST) InOrderTraverseRecur(f func(Elem)) {
	bst.lock.RLock()
	bst.lock.RUnlock()

	inOrderTraverseRecur(bst.Root, f)
}

func inOrderTraverseRecur(node *TreeNode, f func(Elem)) {
	if node != nil {
		inOrderTraverseRecur(node.LChild, f)
		f(node.Data)
		inOrderTraverseRecur(node.RChild, f)
	}
}

// PostOrderTraverseRecur visits BST nodes in post-order recursively
func (bst *BST) PostOrderTraverseRecur(f func(Elem)) {
	bst.lock.RLock()
	bst.lock.RUnlock()

	postOrderTraverseRecur(bst.Root, f)
}

func postOrderTraverseRecur(node *TreeNode, f func(Elem)) {
	if node != nil {
		postOrderTraverseRecur(node.LChild, f)
		postOrderTraverseRecur(node.RChild, f)
		f(node.Data)
	}
}

// PreOrderTraverseIter visits BST nodes in pre-order iteratively
func (bst *BST) PreOrderTraverseIter(f func(Elem)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	preOrderTraverseIter(bst.Root, f)
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
		f(node.Data)
		stack.Push(node.RChild)
		node = node.LChild
	}
}

// InOrderTraverseIter visits BST nodes in in-order iteratively
func (bst *BST) InOrderTraverseIter(f func(Elem)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverseIter(bst.Root, f)
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
		f(node.Data)

		node = node.RChild
		goAlongLeft(stack, node)
	}
}

func goAlongLeft(stack *DynamicStack, node *TreeNode) {
	for node != nil {
		stack.Push(node)
		node = node.LChild
	}
}

// PostOrderTraverseIter visits BST nodes in post-order iteratively
func (bst *BST) PostOrderTraverseIter(f func(Elem)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	postOrderTraverseIter(bst.Root, f)
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
		if pre == nil || pre.LChild == cur || pre.RChild == cur {
			if cur.LChild != nil {
				stack.Push(cur.LChild)
			} else if cur.RChild != nil {
				stack.Push(cur.RChild)
			} else {
				stack.Pop()
				f(cur.Data)
			}
		} else if cur.LChild == pre {
			// traversing up the tree from the left
			if cur.RChild != nil {
				stack.Push(cur.RChild)
			} else {
				stack.Pop()
				f(cur.Data)
			}
		} else if cur.RChild == pre {
			// traversing up the tree from the right
			stack.Pop()
			f(cur.Data)
		}
		pre = cur
	}
}

// LevelTraverse traverse tree level by level
func (bst *BST) LevelTraverse(f func(Elem)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	levelTraverse(bst.Root, f)
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
		f(node.Data)

		if node.LChild != nil {
			queue.Enqueue(node.LChild)
		}
		if node.RChild != nil {
			queue.Enqueue(node.RChild)
		}
	}
}
