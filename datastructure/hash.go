package datastructure

import (
	"fmt"
	"math"
	"sync"
)

var hashsize int

// HashTable struct
type HashTable struct {
	data  []Elem
	count int
	lock  sync.RWMutex
}

// NewHashTable new hash table
func NewHashTable(n int) (*HashTable, error) {
	if n < 1 {
		return nil, fmt.Errorf("invalid initial number %d", n)
	}
	h := &HashTable{
		data: make([]Elem, 0, n),
	}
	hashsize = n
	return h, nil
}

// hash function
func (h *HashTable) hash(key Elem) int {
	return key.(int) % hashsize
}

// reHash second order detect
func (h *HashTable) reHash(addr int) int {
	for i := 1; i <= hashsize>>1; i++ {
		addr = (addr + int(math.Pow(float64(i), 2))) % hashsize
		if h.data[addr] != nil {
			return addr
		}

		addr = (addr - int(math.Pow(float64(i), 2))) % hashsize
		if h.data[addr] != nil {
			return addr
		}
	}
	return addr
}

// Get the element index of hash table
func (h *HashTable) Get(key Elem) (int, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()

	addr := h.hash(key)
	for h.data[addr] != key {
		addr = h.reHash(addr)
		if h.data[addr] == nil || addr == h.hash(key) {
			// nil address or cycle to start-up
			return -1, false
		}
	}
	return addr, true
}

// Put hash table element
func (h *HashTable) Put(key Elem) {
	h.lock.Lock()
	defer h.lock.Unlock()

	addr := h.hash(key)
	if h.data[addr] == nil {
		h.data[addr] = key
	} else {
		addr = h.reHash(addr)
		h.data[addr] = key
	}
	h.count++
}

// Remove hash table element
func (h *HashTable) Remove(key Elem) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	addr, ok := h.Get(key)
	if !ok {
		return fmt.Errorf("does not exist key %v", key)
	}
	h.data[addr] = nil
	return nil
}

// Length of hash table
func (h *HashTable) Length() int {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return h.count
}
