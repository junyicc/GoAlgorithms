package datastructure

import (
	"fmt"
	"hash/fnv"
	"sync"
)

const (
	LRU_LENGTH = 100
	HOST_BIT   = uint64(^uint(0)) == ^uint64(0)
)

type LRU interface {
	Get(string) interface{}
	Put(string, interface{}) error
}

type LRUNode struct {
	k string
	v interface{}

	prev  *LRUNode
	next  *LRUNode
	hnext *LRUNode
}

func NewLRUNode(key string, value interface{}) *LRUNode {
	return &LRUNode{
		k:     key,
		v:     value,
		prev:  nil,
		next:  nil,
		hnext: nil,
	}
}

type LRUCache struct {
	// doubly linked list
	head *LRUNode
	tail *LRUNode

	// linked hash table with head node
	nodes []LRUNode

	capacity int
	used     int

	mu sync.RWMutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		head:     nil,
		tail:     nil,
		nodes:    make([]LRUNode, LRU_LENGTH),
		capacity: capacity,
		used:     0,
	}
}

func (lru *LRUCache) Get(key string) interface{} {
	if lru.tail == nil {
		return nil
	}
	if node := lru.searchNode(key); node != nil {
		lru.moveToTail(node)
		return node.v
	}
	return nil
}

func (lru *LRUCache) Put(key string, value interface{}) error {
	// check in cache
	if node := lru.searchNode(key); node != nil {
		// in cache: update value and move to tail
		node.v = value
		lru.moveToTail(node)
		return nil
	}
	// full cache ? delete head
	if lru.used >= lru.capacity {
		lru.deleteNode()
	}
	// add node
	return lru.addNode(key, value)
}

func (lru *LRUCache) searchNode(key string) *LRUNode {
	lru.mu.RLock()
	defer lru.mu.RUnlock()
	if lru.tail == nil {
		return nil
	}
	idx, _ := hashFunc(key)
	node := (&lru.nodes[idx]).hnext
	for node != nil {
		if node.k == key {
			return node
		}
		node = node.hnext
	}
	return nil
}

func (lru *LRUCache) addNode(key string, value interface{}) error {
	lru.mu.Lock()
	defer lru.mu.Unlock()
	newNode := NewLRUNode(key, value)
	// put into hash table
	idx, err := hashFunc(key)
	if err != nil {
		return err
	}
	head := &lru.nodes[idx]
	newNode.hnext = head.hnext
	head.hnext = newNode
	lru.used++

	// put into doubly linked list
	// check empty cache
	if lru.tail == nil {
		lru.head = newNode
		lru.tail = newNode
		return nil
	}
	lru.tail.next = newNode
	newNode.prev = lru.tail
	lru.tail = newNode

	return nil
}

func (lru *LRUCache) deleteNode() {
	lru.mu.Lock()
	defer lru.mu.Unlock()
	if lru.head == nil {
		return
	}
	deletingNode := lru.head
	// delete head node of doubly linked list
	lru.head = lru.head.next
	lru.head.prev = nil
	lru.used--

	// delete node in hash table
	idx, _ := hashFunc(deletingNode.k)
	node := &lru.nodes[idx]
	for node != nil {
		if node.hnext == deletingNode {
			node.hnext = deletingNode.hnext
			break
		}
		node = node.hnext
	}

}

func (lru *LRUCache) moveToTail(node *LRUNode) {
	lru.mu.Lock()
	defer lru.mu.Unlock()
	if lru.tail == node {
		return
	}

	if lru.head == node {
		lru.head = node.next
		lru.head.prev = nil
	} else {
		// delete from origin position
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	lru.tail.next = node
	node.prev = lru.tail
	node.next = nil
	lru.tail = node
}

// hashFunc firstly calculate hash, and then use A % B = A & (B - 1) to decide index
func hashFunc(input string) (int, error) {
	if HOST_BIT {
		// 64-bit
		h := fnv.New64()
		if n, err := h.Write([]byte(input)); n == 0 || err != nil {
			return 0, fmt.Errorf("cannot hash %s", input)
		}

		k := h.Sum64()
		return int((k ^ (k >> 32)) & (LRU_LENGTH - 1)), nil
	}
	h := fnv.New32()
	if n, err := h.Write([]byte(input)); n == 0 || err != nil {
		return 0, fmt.Errorf("cannot hash %s", input)
	}
	k := h.Sum32()
	return int((k ^ (k >> 16)) & (LRU_LENGTH - 1)), nil
}
