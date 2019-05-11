package datastructure

import (
	"sync"
)

const (
	leftHeigher  = 1
	rightHeigher = -1
	equalHeight  = 0
)

// AVLNode is a self-balance binary search tree node
type AVLNode struct {
	Key    int
	Data   Elem
	LChild *AVLNode
	RChild *AVLNode
	bf     int
}

// AVL tree is a self-balance binary search tree
type AVL struct {
	Root *AVLNode
	lock sync.RWMutex
}

// Insert a node into AVL tree
func (avl *AVL) Insert(key int, data Elem) bool {
	avl.lock.Lock()
	defer avl.lock.Unlock()
	node := AVLNode{key, data, nil, nil, equalHeight}
	if avl.Root == nil {
		avl.Root = &node
		return true
	}
	taller := false
	return insertAVLNode(&avl.Root, &node, &taller)
}

func insertAVLNode(root **AVLNode, node *AVLNode, taller *bool) bool {
	if *root == nil {
		*root = node
		*taller = true
	} else {

		if (*root).Key == node.Key && (*root).Data == node.Data {
			*taller = false
			return false
		}
		if node.Key < (*root).Key {
			// search in left subtree
			if !insertAVLNode(&(*root).LChild, node, taller) {
				// have not inserted
				return false
			}
			if *taller {
				// have inserted into left subtree, and left subtree become taller
				switch (*root).bf {
				case leftHeigher:
					// left subtree was higher than right subtree
					leftBalance(root)
					*taller = false
				case equalHeight:
					// left subtree was equally high as right subtree
					(*root).bf = leftHeigher
					*taller = true
				case rightHeigher:
					// right subtree was higher than left subtree
					(*root).bf = equalHeight
					*taller = false
				}
			}
		} else {
			// search in right subtree
			if !insertAVLNode(&(*root).RChild, node, taller) {
				// have not inserted
				return false
			}
			if *taller {
				switch (*root).bf {
				case leftHeigher:
					// left subtree was higher than right subtree
					(*root).bf = equalHeight
					*taller = false
				case equalHeight:
					// left subtree was equally high as right subtree
					(*root).bf = rightHeigher
					*taller = true
				case rightHeigher:
					// right subtree was higher than left subtree
					rightBalance(root)
					*taller = false
				}
			}
		}
	}
	return true
}

// Search returns true if node exists
func (avl *AVL) Search(key int) (*Elem, bool) {
	avl.lock.RLock()
	defer avl.lock.RUnlock()
	return searchAVL(avl.Root, key)
}

func searchAVL(node *AVLNode, key int) (*Elem, bool) {
	if node == nil {
		return nil, false
	}
	if key < node.Key {
		return searchAVL(node.LChild, key)
	} else if key > node.Key {
		return searchAVL(node.RChild, key)
	}

	return &node.Data, true
}

// RightRotate when bf > 1
func rightRotate(node **AVLNode) {
	p := (*node).LChild
	(*node).LChild = p.RChild
	p.RChild = *node

	*node = p
}

// LeftRotate when bf < -1
func leftRotate(node **AVLNode) {
	p := (*node).RChild
	(*node).RChild = p.LChild
	p.LChild = *node

	*node = p
}

// leftBalance at node when the new node was inserted at left subtree,
// and the sign of leftright childs's bf is differenet from left child's
func leftBalance(node **AVLNode) {
	left := (*node).LChild
	switch left.bf {
	case leftHeigher:
		// new node was inserted at left side of left
		(*node).bf, left.bf = equalHeight, equalHeight
		rightRotate(node)
	case rightHeigher:
		// new node was inserted at right side of left
		leftRight := left.RChild
		switch leftRight.bf {
		case leftHeigher:
			(*node).bf = rightHeigher
			left.bf = equalHeight
		case equalHeight:
			(*node).bf, left.bf = equalHeight, equalHeight
		case rightHeigher:
			(*node).bf = equalHeight
			left.bf = leftHeigher
		}
		leftRight.bf = equalHeight
		// important: cannot pass left as parameter
		leftRotate(&(*node).LChild)
		rightRotate(node)
	}
}

// rightBalance at node when the new node was inserted at left subtree,
// and the sign of rightleft childs's bf is differenet from right child's
func rightBalance(node **AVLNode) {
	right := (*node).RChild
	switch right.bf {
	case rightHeigher:
		// new node was inserted at right side of right
		(*node).bf, right.bf = equalHeight, equalHeight
		leftRotate(node)
	case leftHeigher:
		// new node was inserted at left side of right
		rightLeft := right.LChild
		switch rightLeft.bf {
		case rightHeigher:
			(*node).bf = leftHeigher
			right.bf = equalHeight
		case equalHeight:
			(*node).bf, right.bf = equalHeight, equalHeight
		case leftHeigher:
			(*node).bf = equalHeight
			right.bf = rightHeigher
		}

		rightLeft.bf = equalHeight
		// important: cannot pass right as parameter
		rightRotate(&(*node).RChild)
		leftRotate(node)
	}
}

// PreOrderTraverseRecur visits AVL nodes in pre-order recursively
func (avl *AVL) PreOrderTraverseRecur(node interface{}, f func(Elem)) {
	avl.lock.RLock()
	avl.lock.RUnlock()

	if node, ok := node.(*AVLNode); ok && node != nil {
		f(node.Data)
		avl.PreOrderTraverseRecur(node.LChild, f)
		avl.PreOrderTraverseRecur(node.RChild, f)
	}
}

// InOrderTraverseRecur visits AVL nodes in in-order recursively
func (avl *AVL) InOrderTraverseRecur(node interface{}, f func(Elem)) {
	avl.lock.RLock()
	avl.lock.RUnlock()

	if node, ok := node.(*AVLNode); ok && node != nil {
		avl.InOrderTraverseRecur(node.LChild, f)
		f(node.Data)
		avl.InOrderTraverseRecur(node.RChild, f)
	}
}

// PostOrderTraverseRecur visits AVL nodes in post-order recursively
func (avl *AVL) PostOrderTraverseRecur(node interface{}, f func(Elem)) {
	avl.lock.RLock()
	avl.lock.RUnlock()

	if node, ok := node.(*AVLNode); ok && node != nil {
		avl.PostOrderTraverseRecur(node.LChild, f)
		avl.PostOrderTraverseRecur(node.RChild, f)
		f(node.Data)
	}
}

// PreOrderTraverseIter visits AVL nodes in pre-order iteratively
func (avl *AVL) PreOrderTraverseIter(root interface{}, f func(Elem)) {
	avl.lock.RLock()
	defer avl.lock.RUnlock()

	if root, ok := root.(*AVLNode); ok && root != nil {
		stack := NewDynamicStack()
		stack.Push(root)

		for !stack.IsEmpty() {
			node := (*stack.Pop()).(*AVLNode)
			// visit along left subtree
			for node != nil {
				f(node.Data)
				stack.Push(node.RChild)
				node = node.LChild
			}
		}
	}
}

// InOrderTraverseIter visits AVL nodes in in-order iteratively
func (avl *AVL) InOrderTraverseIter(root interface{}, f func(Elem)) {
	avl.lock.RLock()
	defer avl.lock.RUnlock()

	if root, ok := root.(*AVLNode); ok && root != nil {
		stack := NewDynamicStack()

		var node = root
		for {
			// go along left subtree
			for node != nil {
				stack.Push(node)
				node = node.LChild
			}
			if stack.IsEmpty() {
				break
			}

			root := (*stack.Pop()).(*AVLNode)
			f(root.Data)
			node = root.RChild
		}
	}
}

// PostOrderTraverseIter visits AVL nodes in post-order iteratively
func (avl *AVL) PostOrderTraverseIter(root interface{}, f func(Elem)) {
	avl.lock.RLock()
	defer avl.lock.RUnlock()

	if root, ok := root.(*AVLNode); ok && root != nil {
		stack := NewDynamicStack()
		stack.Push(root)
		var pre *AVLNode
		for !stack.IsEmpty() {
			cur := (*stack.GetTop()).(*AVLNode)
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
}

// LevelTraverse traverse tree level by level
func (avl *AVL) LevelTraverse(root interface{}, f func(Elem)) {
	avl.lock.RLock()
	defer avl.lock.RUnlock()

	if root, ok := root.(*AVLNode); ok && root != nil {
		queue := NewDynamicQueue()
		queue.Enqueue(root)

		for !queue.IsEmpty() {
			node := (*queue.Dequeue()).(*AVLNode)
			f(node.Data)

			if node.LChild != nil {
				queue.Enqueue(node.LChild)
			}
			if node.RChild != nil {
				queue.Enqueue(node.RChild)
			}
		}
	}
}
